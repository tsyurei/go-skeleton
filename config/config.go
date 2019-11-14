package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	// web application port
	Port int
	// A secret string used for session cookies, passwords, etc.
	SecretKey []byte
}

func InitConfig() (*Config, error) {
	config := &Config{
		Port:       viper.GetInt("port"),
		SecretKey: []byte(viper.GetString("SecretKey")),
	}

	if len(config.SecretKey) == 0 {
		return nil, fmt.Errorf("Secret Key must be set!")
	}

	return config, nil
}
