package config

import (
	"github.com/spf13/viper"
)

type config struct {
	Passphrase string
	CipherKey  string
	Id *string
}

var Config config

func init() {
	viper.AddConfigPath("$HOME/.config")
	viper.SetConfigName("dtp")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	passphrase := viper.GetString("passphrase")
	cipherKey := viper.GetString("cipher-key")

	Config = config{
		Passphrase: passphrase,
		CipherKey:  cipherKey,
	}
}
