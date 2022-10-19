package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Rest     string
	Memcache string
}

func ParseConfig(path string) (*Config, error){
	//Parsing config
	setDefaults()
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	fmt.Printf(cfg.Rest)
	return &cfg, nil
}
func setDefaults(){
	viper.SetDefault("Rest", "8080")
	viper.SetDefault("Memcache", "11211")
}