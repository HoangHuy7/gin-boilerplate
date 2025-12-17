package app

import (
	"log"
	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}
type GoConfig struct {
	Port int `mapstructure:"port"`
}
type Config struct {
	Database struct {
		Master DatabaseConfig `mapstructure:"master"`
		Worker DatabaseConfig `mapstructure:"worker"`
	} `mapstructure:"database"`

	Go GoConfig `mapstructure:"go"`
}

var C Config // global for convenience

func LoadConfig() {
	viper.SetConfigName("configs")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./configs")
	viper.SetEnvPrefix("")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	if err := viper.Unmarshal(&C); err != nil {
		log.Fatalf("Error parsing config: %v", err)
	}

	if C.Go.Port == 0 {
		C.Go.Port = 8080 // default port
	}
}
