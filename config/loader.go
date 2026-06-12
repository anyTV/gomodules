package config

import (
	"fmt"
	"os"
	"strings"

	logger "github.com/anyTV/gomodules/v2/logging"
	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/viper"
)

// New automatically loads configuration based on an environment variable.
// It assumes the following standard naming convention:
//  1. .env.yaml             (Base - Must exist)
//  2. .env.{ENV}.yaml       (Environment Override - Optional)
//  3. .env.{ENV}.local.yaml (Local Environment Override - Optional)
//  4. .env.local.yaml       (Local Override - Optional)
//
// Example: config.LoadEnv[AppConfig]()
func New[T any]() (*T, error) {
	env := os.Getenv("ENV")

	names := []string{
		".env", // Base
	}

	if env != "" {
		names = append(names,
			fmt.Sprintf(".env.%s", env),       // Env specific
			fmt.Sprintf(".env.%s.local", env), // Env + Local specific
		)
	}

	names = append(names, ".env.local") // Local overrides

	return loadConfigNames[T](names...)
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
	v.SetConfigType("yaml")

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
	if err := v.Unmarshal(&cfg, viper.DecodeHook(
		mapstructure.ComposeDecodeHookFunc(
			// Default values:
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			// New values:
			mapstructure.TextUnmarshallerHookFunc(), // Also enables adding custom UnmarshalText hooks
		),
	)); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}

func loadConfigNames[T any](names ...string) (*T, error) {
	if len(names) == 0 {
		return nil, fmt.Errorf("at least one config file must be provided")
	}

	v := viper.New()

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	v.SetConfigType("yaml")

	if err := addConfigPaths(v); err != nil {
		return nil, err
	}

	for i, name := range names {
		v.SetConfigName(name)

		if i == 0 {
			if err := v.ReadInConfig(); err != nil {
				return nil, fmt.Errorf("failed to read base config %s.yaml: %w", name, err)
			}
		} else {
			if err := v.MergeInConfig(); err == nil {
				logger.Debugf("Merged config: %s", v.ConfigFileUsed())
			}
		}
	}

	var cfg T
	if err := v.Unmarshal(&cfg, viper.DecodeHook(
		mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			mapstructure.TextUnmarshallerHookFunc(),
		),
	)); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}

func addConfigPaths(v *viper.Viper) error {
	configPath := "./"

	for range maxDepth {
		v.AddConfigPath(configPath)
		configPath += "../"
	}

	return nil
}
