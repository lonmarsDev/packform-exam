package config

import (
	"github.com/lonmarsDev/packform-exam/pkg/log"
	"github.com/lonmarsDev/packform-exam/pkg/utils"
	"github.com/spf13/viper"
)

var AppConfig *viper.Viper

func Init() {

	log.Info("%s", "config initialization...")
	defer log.Info("%s", "config initialization...done!")
	AppConfig = viper.New()
	AppConfig.AddConfigPath(".")
	AppConfig.SetConfigName("config")
	AppConfig.SetConfigType("json")
	err := AppConfig.ReadInConfig()
	if err != nil {
		log.Error("viper error: %v", err)
	}

	log.Info("config: %s", utils.PrettyPrint(AppConfig.AllSettings()))
	return
}
