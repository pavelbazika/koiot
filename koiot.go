package koiot

import (
	"context"
	"fmt"

	"github.com/pavelbazika/ewelink"
)

const termId = "1001770323"

//const dole = "1000c3a9b5"

func Run() {
	ctx := context.Background()
	instance := ewelink.New()

	// authenticate using email
	session, err := instance.AuthenticateWithEmail(
		context.Background(), ewelink.NewConfiguration("eu"), "kokes@kocovnici.cz", "kamaradi")

	thr := NewTHR316(instance, session, termId)
	term, err := thr.GetTemp(context.Background())
	//devices, err := instance.GetDevices(ctx, session)
	//fmt.Println(devices)

	dole, err := instance.GetDevice(ctx, session, termId)
	fmt.Println(dole)

	// retrieve the list of registered devices
	//term, err := instance.GetDevice(context.Background(), session, termId)

	// turn on the outlet(s) of the first device
	//response, err := instance.SetDevicePowerState(context.Background(), session, &devices.Devicelist[0], true)

	fmt.Println(term)
	fmt.Println(err)
}
