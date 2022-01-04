package configuration

import (
	"os"

	"github.com/spf13/viper"
	"gopkg.in/inconshreveable/log15.v2"
)

func GetConfiguration() *viper.Viper {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./resource/")
	viper.AutomaticEnv()
	error := viper.ReadInConfig()
	if error != nil {
		log15.Error("Configuration read error = ", error.Error())
		os.Exit(1)
	}
	return viper.GetViper()
}
