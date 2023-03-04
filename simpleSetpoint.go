package koiot

import "context"

type SimpleSetpoint struct {
	setPoint float32
	max      float32
	min      float32
	temp     Thermometer
	switches []Switch

	action bool
}

func NewSimpleSetpoint(setPoint, max, min float32, temp Thermometer, switches []Switch) *SimpleSetpoint {
	return &SimpleSetpoint{
		setPoint: setPoint,
		max:      max,
		min:      min,
		temp:     temp,
		switches: switches,
	}
}

func (sp *SimpleSetpoint) Tick(ctx context.Context) error {
	temp, err := sp.temp.GetTemp(ctx)
	if err != nil {
		return err
	}

	if !sp.action && (temp < sp.min) {
		err = sp.swtch(ctx, true)
		if err != nil {
			return err
		}
		sp.action = true
	} else if sp.action && (temp > sp.max) {
		err = sp.swtch(ctx, false)
		if err != nil {
			return err
		}
		sp.action = false
	}

	return nil
}

func (sp *SimpleSetpoint) swtch(ctx context.Context, on bool) error {
	var err error

	for _, sw := range sp.switches {
		err2 := sw.Switch(ctx, on)
		if err == nil {
			err = err2
		}
	}

	return err
}
