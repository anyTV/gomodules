package config

import (
	"fmt"
	"os"
	"strings"

	logger "github.com/anyTV/gomodules/v2/logging"
	"github.com/spf13/viper"
)

// New automatically loads configuration based on an environment variable.
// It assumes the following standard naming convention:
//  1. .env.yaml             (Base - Must exist)
//  2. .env.{ENV}.yaml       (Environment Override - Optional)
//  3. .env.{ENV}.local.yaml (Local Environment Override - Optional)
//  3. .env.local.yaml       (Local Developer Override - Optional)
//
// Example: config.LoadEnv[AppConfig]("ENV")
func New[T any]() (*T, error) {
	env := os.Getenv("ENV")
	if env == "" {
		return nil, fmt.Errorf("missing env value")
	}

	files := []string{
		".env.yaml",                            // Base
		fmt.Sprintf(".env.%s.yaml", env),       // Env specific
		fmt.Sprintf(".env.%s.local.yaml", env), // Env + Local specific
		".env.local.yaml",                      // Local overrides
	}

	return Load[T](files...)
}

// Load reads configuration from the provided files into a struct of type T.
// It uses a fresh Viper instance, so it won't interfere with the global viper used in NewConfig().
//
// Usage:
//
//	cfg, err := config.Load[AppConfig]("./config.yaml", "./secrets.json")
func Load[T any](paths ...string) (*T, error) {
	if len(paths) == 0 {
		return nil, fmt.Errorf("at least one config file must be provided")
	}

	v := viper.New()

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	for i, path := range paths {
		v.SetConfigFile(path)

		// Read first file
		if i == 0 {
			if err := v.ReadInConfig(); err != nil {
				return nil, fmt.Errorf("failed to read base config %s: %w", path, err)
			}
		} else { // Merge subsequent files (Optional)
			if err := v.MergeInConfig(); err == nil {
				logger.Debugf("Merged config: %s", path)
			}
		}
	}

	var cfg T
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}
