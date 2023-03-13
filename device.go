package koiot

import "context"

var errOffline error

type (
	Device interface{}

	Thermometer interface {
		Device
		GetTemp(ctx context.Context) (float32, error)
	}

	Switch interface {
		Device
		Switch(ctx context.Context, on bool) error
	}
)
