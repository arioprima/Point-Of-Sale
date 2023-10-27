package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	DBUser         string        `mapstructure:"DB_USER"`
	DBPassword     string        `mapstructure:"DB_PASSWORD"`
	DBHost         string        `mapstructure:"DB_HOST"`
	DBPort         string        `mapstructure:"DB_PORT"`
	DBName         string        `mapstructure:"DB_NAME"`
	ServerPort     string        `mapstructure:"PORT"`
	TokenSecret    string        `mapstructure:"JWT_TOKEN_SECRET"`
	TokenExpiresIn time.Duration `mapstructure:"JWT_TOKEN_EXPIRED_IN"`
	TokenMaxAge    int           `mapstructure:"TOKEN_MAXAGE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&config)

	return config, err
}
