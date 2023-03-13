package main

import (
	"flag"
	"fmt"
	"os"

	"beerchart.cz/koiot"
	"github.com/kardianos/service"
	"gitlab.icewarp.com/go/shared/logger"
)

var (
	svcFlag   = flag.String("service", "", "Control the system service (start, stop, restart, install, uninstall)")
	svcUser   = flag.String("user", "", "User name the service will run under, root by default")
	configDir = flag.String("config", "", "Config directory (default try in app/config directory)")
)

func main() {
	svcOptions := make(service.KeyValue)
	svcOptions["Restart"] = "on-failure"
	svcOptions["Enable"] = false

	svcConfig := &service.Config{
		Name:        "koiot",
		DisplayName: "koiot",
		Description: "KOkes Internet Of Things",
		UserName:    *svcUser,
		Option:      svcOptions,
		Arguments:   []string{fmt.Sprintf("-config=%s", *configDir)},
	}
	prg := &koiot.KoiotService{
		ConfigPath: *configDir,
	}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	// Manage service
	if len(*svcFlag) != 0 {
		err := service.Control(s, *svcFlag)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(2)
		}

		os.Exit(0)
	}

	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
