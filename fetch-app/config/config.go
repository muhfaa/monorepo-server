package config

import (
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Port       int    `mapstructure:"port"`
	DbDriver   string `mapstructure:"dbDriver"`
	DbName     string `mapstructure:"dbName"`
	DbAddress  string `mapstructure:"dbAddress"`
	DbPort     int    `mapstructure:"dbPort"`
	DbUsername string `mapstructure:"dbUsername"`
	DbPassword string `mapstructure:"dbPassword"`
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

//GetConfig Initiatilize config in singleton way
func GetConfig() *AppConfig {
	if appConfig != nil {
		return appConfig
	}

	lock.Lock()
	defer lock.Unlock()

	//re-check after locking
	if appConfig != nil {
		return appConfig
	}

	appConfig = initConfig()

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig

	defaultConfig.Port = 8000
	defaultConfig.DbDriver = "sqlite"
	defaultConfig.DbName = "efishery-microservice"
	defaultConfig.DbAddress = "localhost"
	defaultConfig.DbPort = 3306
	defaultConfig.DbUsername = ""
	defaultConfig.DbPassword = ""

	viper.AutomaticEnv()
	viper.SetEnvPrefix("efishery")
	viper.BindEnv("port")
	viper.BindEnv("dbDriver")
	viper.BindEnv("dbName")
	viper.BindEnv("dbAddress")
	viper.BindEnv("dbPort")
	viper.BindEnv("dbUsername")
	viper.BindEnv("dbPassword")

	var finalConfig AppConfig
	err := viper.Unmarshal(&finalConfig)
	if err != nil {
		log.Error("failed to extract config, will use default value")
		return &defaultConfig
	}

	return &finalConfig

}
