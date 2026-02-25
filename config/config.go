package config

import (
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var maxDepth = 5

// Deprecated: Use config.New[env.AppConfig]() instead
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

// Deprecated: Use config.New[env.AppConfig]() instead
func SetMaxDepth(v int) {
	maxDepth = v
}

// Deprecated: Use config.New[env.AppConfig]() instead
func GetString(key string) string {
	return viper.GetString(key)
}

// Deprecated: Use config.New[env.AppConfig]() instead
func GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)

}

// Deprecated: Use config.New[env.AppConfig]() instead
func GetStringMapString(key string) map[string]string {
	return viper.GetStringMapString(key)
}

// Deprecated: Use config.New[env.AppConfig]() instead
func Get(key string) any {
	return viper.Get(key)
}

// Deprecated: Use config.New[env.AppConfig]() instead
func GetInt(key string) int {
	return viper.GetInt(key)
}

// Deprecated: Use config.New[env.AppConfig]() instead
func GetBool(key string) bool {
	return viper.GetBool(key)
}
