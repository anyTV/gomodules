package config

import (
	"os"

	L "github.com/anyTV/gomodules/logging"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const maxDepth = 5

var log = L.New("config.go")

func init() {
	var configFile = ".env.yaml"
	var fileType = "yaml"

	if os.Getenv("ENV") == "production" {
		configFile = ".env.production"
		gin.SetMode(gin.ReleaseMode)
		viper.Set("env", "production")
	}

	viper.SetConfigName(configFile)
	viper.SetConfigType(fileType)
	var configPath = "./"

	for i := 0; i < maxDepth; i++ {
		viper.AddConfigPath(configPath)

		log.Debugf(
			"Reading %s (%s) file for config...",
			configFile,
			fileType,
		)

		err := viper.ReadInConfig()

		if err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				configPath = configPath + "../"
				log.Warnf("configuration not found. Adding path: %s", configPath)
				viper.AddConfigPath(configPath)
			} else {
				log.Warnf("error with configuration: %s", err)
			}
		} else {
			log.Infof("File `%s` found and loaded", configFile)
			break
		}
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Configuration updated: %s", e.Name)
	})

	log.Infof("%s", GetString("greeting"))
	viper.WatchConfig()
}

func GetString(key string) string {

	val := viper.GetString(key)

	if val == "" {
		log.Warnf("Key `%s` is empty", key)
	}

	return val
}

func GetStringSlice(key string) []string {
	val := viper.GetStringSlice(key)

	if len(val) == 0 {
		log.Warnf("Key `%s` is empty", key)
	}

	return val
}

func GetStringMapString(key string) map[string]string {
	val := viper.GetStringMapString(key)

	if len(val) == 0 {
		log.Warnf("Key `%s` is empty", key)
	}

	return val
}

func Get(key string) any {
	val := viper.Get(key)

	if val == nil {
		log.Warnf("Key `%s` is empty", key)
	}

	return val
}
