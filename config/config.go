package config

import "github.com/gurkankaymak/hocon"

func blah() {
	_, err := hocon.ParseResource("hoconString")
	panic(err)
}
