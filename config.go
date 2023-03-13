package koiot

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type (
	cfgGeneralHeader struct {
		Document string `yaml:"document"`
	}

	cfgAuth struct {
		Region   string `yaml:"region"`
		Email    string `yaml:"email"`
		Password string `yaml:"password"`
		TimeoutS uint   `yaml:"timeout_s"`
	}

	cfgEwDevice struct {
		Name string `yaml:"name"`
		Id   string `yaml:"id"`
	}

	cfgMultiSwitch struct {
		Name     string   `yaml:"name"`
		Switches []string `yaml:"switches"`
	}

	cfgDevices struct {
		EwDevices     []*cfgEwDevice    `yaml:"ew_devices"`
		MultiSwitches []*cfgMultiSwitch `yaml:"multiswitches"`
		LoadTimeoutS  uint              `yaml:"load_timeout_s"`
	}

	cfgSilentHeater struct {
		Name   string `yaml:"name"`
		Device struct {
			Temp   string `yaml:"temp"`
			Switch string `yaml:"switch"`
		} `yaml:"device"`
		Plan     Schedule[SilentHeaterAction] `yaml:"plan"`
		TimeoutS uint                         `yaml:"timeout_s"`
	}

	tConfig struct {
		auth          cfgAuth
		devices       cfgDevices
		silentHeaters []*cfgSilentHeater
	}
)

func LoadConfig(configDir string) (*tConfig, error) {
	var result tConfig

	err := filepath.WalkDir(configDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			if configDir == path {
				return nil
			} else {
				return filepath.SkipDir
			}
		}

		if filepath.Ext(path) != ".yaml" {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		var header cfgGeneralHeader
		if err = yaml.Unmarshal(data, &header); err != nil {
			return err
		}

		switch header.Document {
		case "koiot_auth":
			result.auth, err = cfgLoadAuth(data)
			if err != nil {
				return err
			}

		case "koiot_devices":
			result.devices, err = cfgLoadDevices(data)
			if err != nil {
				return err
			}

		case "koiot_plan_silentheater":
			heater, err := cfgLoadSilentHeater(data)
			if err != nil {
				return err
			}
			result.silentHeaters = append(result.silentHeaters, &heater)

		default:
			return fmt.Errorf("unknown config document type %s", header.Document)
		}
		return nil
	})

	if err == nil {
		return &result, nil
	} else {
		return nil, err
	}
}

func cfgLoadAuth(data []byte) (auth cfgAuth, err error) {
	err = yaml.Unmarshal(data, &auth)
	return
}

func cfgLoadDevices(data []byte) (devices cfgDevices, err error) {
	err = yaml.Unmarshal(data, &devices)
	return
}

func cfgLoadSilentHeater(data []byte) (heater cfgSilentHeater, err error) {
	err = yaml.Unmarshal(data, &heater)
	return
}
