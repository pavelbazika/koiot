package koiot

import (
	"context"
	"errors"
	"strconv"

	"github.com/pavelbazika/ewelink"
)

type Thermometer interface {
	GetTemp(ctx context.Context) (float32, error)
}

type Switch interface {
	Switch(ctx context.Context, on bool) error
}

type tTHR316 struct {
	ew      *ewelink.Ewelink
	session *ewelink.Session
	id      string
}

var _ Thermometer = &tTHR316{}
var _ Switch = &tTHR316{}

func NewTHR316(ew *ewelink.Ewelink, session *ewelink.Session, id string) Thermometer {
	return &tTHR316{
		ew:      ew,
		session: session,
		id:      id,
	}
}

func (thr *tTHR316) GetTemp(ctx context.Context) (float32, error) {
	status, err := thr.getDeviceInfo(ctx)
	if err != nil {
		return 0, err
	}

	if !status.Online {
		return 0, errors.New("device is offline")
	}

	temp, err := strconv.ParseFloat(status.Params.CurrentTemperature, 32)
	if err != nil {
		return 0, err
	}
	return float32(temp), nil
}

func (thr *tTHR316) Switch(ctx context.Context, on bool) error {
	status, err := thr.getDeviceInfo(ctx)
	if err != nil {
		return err
	}

	if !status.Online {
		return errors.New("device is offline")
	}

	_, err = thr.ew.SetDevicePowerState(ctx, thr.session, status, on)
	return err
}

func (thr *tTHR316) getDeviceInfo(ctx context.Context) (*ewelink.Device, error) {
	return thr.ew.GetDevice(ctx, thr.session, thr.id)
}
