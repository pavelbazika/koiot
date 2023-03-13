package koiot

import (
	"context"
	"fmt"
)

type tMultiSwitch struct {
	childs []Switch
}

func NewMultiSwitch(childs []Switch) *tMultiSwitch {
	return &tMultiSwitch{
		childs: childs,
	}
}

func (mult *tMultiSwitch) Switch(ctx context.Context, on bool) error {
	var totalErr error
	for _, sw := range mult.childs {
		swErr := sw.Switch(ctx, on)
		if totalErr == nil {
			totalErr = swErr
		} else {
			totalErr = fmt.Errorf("%w, %w", totalErr, swErr)
		}
	}

	return totalErr
}
