package config

import (
	"github.com/gurkankaymak/hocon"
)

var hoconConfig *hocon.Config = nil

func Config() *hocon.Config {
	if hoconConfig == nil {
		conf, err := hocon.ParseResource("config/application.conf")
		if err != nil {
			panic(err)
		}
		hoconConfig = conf
	}

	return hoconConfig
}
