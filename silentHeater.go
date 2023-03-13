package koiot

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type SilentHeaterAction struct {
	Enabled bool    `yaml:"enabled"`
	HardOn  bool    `yaml:"hard_on"`
	HardOff bool    `yaml:"hard_off"`
	MinTemp float32 `yaml:"min_temp"`
	MaxTemp float32 `yaml:"max_temp"`
}

type SilentHeater struct {
	plan      Schedule[SilentHeaterAction]
	temp      Thermometer
	sw        Switch
	opTimeout time.Duration

	enabled bool
	exit    chan chan struct{}
}

func NewSilentHeater(plan Schedule[SilentHeaterAction], temp Thermometer, sw Switch, opTimeout time.Duration) (*SilentHeater, error) {
	result := SilentHeater{
		plan:      plan,
		temp:      temp,
		sw:        sw,
		exit:      make(chan chan struct{}),
		opTimeout: opTimeout,
	}

	ctx, cancel := result.createCtx()
	defer cancel()
	err := result.swtch(ctx, false)
	if err != nil {
		return nil, err
	}

	go result.ticker()
	return &result, nil
}

func (ht *SilentHeater) Stop() {
	ch := make(chan struct{})
	ht.exit <- ch
	<-ch
}

func (ht *SilentHeater) ticker() {
	// first tick
	if err := ht.tick(); err != nil {
		fmt.Printf("SilentHeater error: %s\n", err.Error())
	}

	t := time.NewTicker(time.Minute)

	for {
		select {
		case <-t.C:
			if err := ht.tick(); err != nil {
				fmt.Printf("SilentHeater error: %s\n", err.Error())
			}

		case c := <-ht.exit:
			t.Stop()
			close(c)
			return
		}
	}
}

func (ht *SilentHeater) tick() error {
	ctx, cancel := ht.createCtx()
	defer cancel()

	now := time.Now()
	genAction := ht.plan.GetCurrentItemData(now)
	if action, ok := genAction.(SilentHeaterAction); ok {
		switch {
		case action.HardOn:
			if !action.HardOff {
				return ht.swtch(ctx, true)
			} else {
				return errors.New("SilentHeater has command to be on and off simultaneously")
			}

		case action.HardOff:
			return ht.swtch(ctx, false)

		default:
			curTemp, err := ht.temp.GetTemp(ctx)
			if err != nil {
				return err
			}

			if ht.enabled && (curTemp > action.MaxTemp) {
				return ht.swtch(ctx, false)
			} else if !ht.enabled && (curTemp < action.MinTemp) {
				return ht.swtch(ctx, true)
			}
		}
	}

	return nil
}

func (ht *SilentHeater) createCtx() (context.Context, context.CancelFunc) {
	if ht.opTimeout == 0 {
		return context.Background(), func() {}
	} else {
		return context.WithTimeout(context.Background(), ht.opTimeout)
	}
}

func (ht *SilentHeater) swtch(ctx context.Context, on bool) (err error) {
	if err = ht.sw.Switch(ctx, on); err != nil {
		return
	}

	ht.enabled = on
	return
}
