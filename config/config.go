package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	// web application port
	Port int

	// The number of proxies positioned in front of the API. This is used to interpret
	// X-Forwarded-For headers.
	ProxyCount int

	// A secret string used for session cookies, passwords, etc.
	SecretKey []byte

	// JWT secret key
	JWTKey []byte
	// JWT secret key
	RefreshJWTKey []byte

	// AuthToken expire minute
	AuthTokenExpireMinute int

	// Refresh AuthToken expire minute
	RefreshAuthTokenExpireMinute int
}

var AppConfig *Config

func initConfigDefaultValue() {
	viper.SetDefault("authentication.tokenExpireMinute", 15)
	viper.SetDefault("authentication.refreshTokenExpireMinute", 4320)
}

func InitConfig() (*Config, error) {
	initConfigDefaultValue()
	AppConfig = &Config{
		Port:          viper.GetInt("port"),
		ProxyCount:    viper.GetInt("ProxyCount"),
		SecretKey:     []byte(viper.GetString("SecretKey")),
		JWTKey:        []byte(viper.GetString("SecretKey")),
		RefreshJWTKey: []byte(viper.GetString("RefreshSecretKey")),
		AuthTokenExpireMinute: viper.GetInt("authentication.tokenExpireMinute"),
		RefreshAuthTokenExpireMinute: viper.GetInt("authentication.refreshTokenExpireMinute"),
	}

	if len(AppConfig.SecretKey) == 0 {
		return nil, fmt.Errorf("Secret Key must be set!")
	} else if len(AppConfig.SecretKey) < 32 {
		return nil, fmt.Errorf("Secret Key length must at least 32")
	}

	if len(AppConfig.RefreshJWTKey) == 0 {
		return nil, fmt.Errorf("Refresh Secret Key must be set!")
	} else if len(AppConfig.RefreshJWTKey) < 32 {
		return nil, fmt.Errorf("Refresh Secret Key length must at least 32")
	}

	return AppConfig, nil
}
