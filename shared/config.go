package shared

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort string
	DBHost  string
	DBPort  string
	DBUser  string
	DBPass  string
	DBName  string
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	return &Config{
		AppPort: viper.GetString("APP_PORT"),
        DBHost:  viper.GetString("DB_HOST"),
        DBPort:  viper.GetString("DB_PORT"),
        DBUser:  viper.GetString("DB_USER"),
        DBPass:  viper.GetString("DB_PASS"),
        DBName:  viper.GetString("DB_NAME"),
	}
}