package koiot

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/kardianos/service"
	"github.com/pavelbazika/ewelink"
)

var log service.Logger

func logFatal(msg string) {
	log.Error(msg)
	os.Exit(2)
}

func logFatalf(msg string, args ...interface{}) {
	log.Errorf(msg, args...)
	os.Exit(2)
}

type tEwork struct {
	*ewelink.Ewelink
}

func (ew *tEwork) authenticate(cfg *cfgAuth) (session *ewelink.Session, err error) {
	if (cfg.Region == "") || (cfg.Email == "") {
		return nil, errors.New("incomplete credentials")
	}

	log.Info("Authenticating")
	var ctx context.Context
	var cancel context.CancelFunc
	if cfg.TimeoutS == 0 {
		ctx = context.Background()
		cancel = func() {}
	} else {
		ctx, cancel = context.WithTimeout(context.Background(), time.Duration(cfg.TimeoutS)*time.Second)
	}
	defer cancel()

	session, err = ew.AuthenticateWithEmail(ctx, ewelink.NewConfiguration(cfg.Region), cfg.Email, cfg.Password)

	log.Info("Done")
	return
}

type tDevicesMap map[string]Device

func (ew *tEwork) loadDevices(session *ewelink.Session, cfg *cfgDevices) (result tDevicesMap, err error) {
	log.Info("Loading devices")
	var ctx context.Context
	var cancel context.CancelFunc
	if cfg.LoadTimeoutS == 0 {
		ctx = context.Background()
		cancel = func() {}
	} else {
		ctx, cancel = context.WithTimeout(context.Background(), time.Duration(cfg.LoadTimeoutS)*time.Second)
	}
	defer cancel()

	ewDevs, err := ew.GetDevices(ctx, session)
	if err != nil {
		return nil, err
	}

	result = make(tDevicesMap)
	for _, device := range cfg.EwDevices {
		if _, found := result[device.Name]; found {
			// already loaded
			return nil, fmt.Errorf("duplicate device in config with name %s", device.Name)
		}

		// find the device
		found := false
		var ewDev ewelink.Device
		for _, ewDev = range ewDevs.Devicelist {
			if ewDev.DeviceID == device.Id {
				found = true
				break
			}
		}

		if !found {
			return nil, fmt.Errorf("cannot find device with id %s", device.Id)
		}

		switch ewDev.Extra.Extra.Model {
		case "SN-ESP32D0-THR3-01":
			result[device.Name] = NewTHR316(ew.Ewelink, session, device.Id)

		case "PSA-B01-GL":
			result[device.Name] = NewSocket(ew.Ewelink, session, device.Id)

		default:
			return nil, fmt.Errorf("device with id %s is of unsupported model %s", device.Id, ewDev.Extra.Extra.Model)
		}
	}

	for _, msw := range cfg.MultiSwitches {
		if _, found := result[msw.Name]; found {
			// already loaded
			return nil, fmt.Errorf("duplicate multiswitch in config with name %s", msw.Name)
		}

		var switches []Switch
		for _, childName := range msw.Switches {
			if sw, found := result[childName]; found {
				if theSw, ok := sw.(Switch); ok {
					switches = append(switches, theSw)
				} else {
					return nil, fmt.Errorf("when creating multiswitch %s, device with name %s is not switch", msw.Name, childName)
				}
			} else {
				return nil, fmt.Errorf("when creating multiswitch %s, device with name %s not found", msw.Name, childName)
			}
		}
		result[msw.Name] = NewMultiSwitch(switches)
	}

	log.Info("Done")
	return result, nil
}

func startSilentHeaters(cfg []*cfgSilentHeater, devices tDevicesMap) (result []*SilentHeater, err error) {
	for _, sh := range cfg {
		temp, ok := devices[sh.Device.Temp]
		if !ok {
			return nil, fmt.Errorf("error configuring silent heater %s, temp sensor %s not found", sh.Name, sh.Device.Temp)
		}
		theTemp, ok := temp.(Thermometer)
		if !ok {
			return nil, fmt.Errorf("error configuring silent heater %s, device %s is not temp sensor", sh.Name, sh.Device.Temp)
		}

		sw, ok := devices[sh.Device.Switch]
		if !ok {
			return nil, fmt.Errorf("error configuring silent heater %s, switch %s not found", sh.Name, sh.Device.Switch)
		}
		theSw, ok := sw.(Switch)
		if !ok {
			return nil, fmt.Errorf("error configuring silent heater %s, device %s is not switch", sh.Name, sh.Device.Switch)
		}

		newSH, err := NewSilentHeater(sh.Plan, theTemp, theSw, time.Duration(sh.TimeoutS)*time.Second)
		if err != nil {
			return nil, fmt.Errorf("failed to configure silent heater %s: %w", sh.Name, err)
		}

		result = append(result, newSH)
	}

	return result, nil
}

type KoiotService struct {
	ConfigPath string
	exitRq     chan chan struct{}
}

func (k *KoiotService) Start(s service.Service) error {
	log, _ = s.Logger(nil)
	k.exitRq = make(chan chan struct{})
	go k.run()
	return nil
}

func (k *KoiotService) Stop(s service.Service) error {
	exitRsp := make(chan struct{})
	k.exitRq <- exitRsp
	<-exitRsp
	return nil
}

func (k *KoiotService) run() {
	if k.ConfigPath == "" {
		k.ConfigPath = "config"
	}
	config, err := LoadConfig(k.ConfigPath)
	if err != nil {
		logFatal(err.Error())
	}

	var ew tEwork
	ew.Ewelink = ewelink.New()

	session, err := ew.authenticate(&config.auth)
	if err != nil {
		logFatalf("eWelink authentication failed: %s", err.Error())
	}

	devices, err := ew.loadDevices(session, &config.devices)
	if err != nil {
		logFatalf("Failed to load all required devices: %s", err.Error())
	}

	_, err = startSilentHeaters(config.silentHeaters, devices)
	if err != nil {
		logFatalf("Failed to load silent heaters: %s", err.Error())
	}

	exitRsp := <-k.exitRq
	defer close(exitRsp)

	// do terminate
}
