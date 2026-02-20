package config

import "github.com/spf13/viper"

type Config struct {
	Port int `mapstructure:"PORT"`
}

func LoadConfig() (*Config, error) {
	v := viper.New()
	v.SetConfigFile(".env")
	v.AutomaticEnv()

	_ = v.ReadInConfig()

	setDefaultConfig(v)

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func setDefaultConfig(v *viper.Viper) {
	v.SetDefault("PORT", 8080)
}
