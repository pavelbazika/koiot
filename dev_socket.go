package koiot

import (
	"context"
	"errors"

	"github.com/pavelbazika/ewelink"
)

type tSocket struct {
	ew      *ewelink.Ewelink
	session *ewelink.Session
	id      string
}

var _ Switch = &tSocket{}

func NewSocket(ew *ewelink.Ewelink, session *ewelink.Session, id string) *tSocket {
	return &tSocket{
		ew:      ew,
		session: session,
		id:      id,
	}
}

func (sock *tSocket) Switch(ctx context.Context, on bool) error {
	status, err := sock.ew.GetDevice(ctx, sock.session, sock.id)
	if err != nil {
		return err
	}

	if !status.Online {
		return errors.New("device is offline")
	}

	_, err = sock.ew.SetDevicePowerState(ctx, sock.session, status, on)
	return err
}
