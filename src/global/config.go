package global

import (
	"gopkg.in/ini.v1"
)

var Config *ini.File

func LoadConfig() (err error) {
	mode, err := ini.Load("./config.ini")
	if err != nil {
		return err
	}

	appMode := mode.Section("").Key("app_mode").String()
	configName := "./config-" + appMode + ".ini"
	Config, err = ini.Load(configName)
	if err != nil {
		return err
	}
	return
}
