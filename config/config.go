package config

import (
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var maxDepth = 5

func NewConfig() {
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

		err := viper.ReadInConfig()

		if err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				configPath = configPath + "../"
				viper.AddConfigPath(configPath)
			}
		} else {
			break
		}
	}

	viper.OnConfigChange(func(e fsnotify.Event) {})

	viper.WatchConfig()
}

func SetMaxDepth(v int) {
	maxDepth = v
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)

}

func GetStringMapString(key string) map[string]string {
	return viper.GetStringMapString(key)
}

func Get(key string) any {
	return viper.Get(key)
}
