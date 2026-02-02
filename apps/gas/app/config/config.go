// hoanghuy7 from Vietnamese with love!

package config

import (
	"log"
	"monorepo/internal/dto"
	"monorepo/internal/utils"

	"github.com/google/uuid"
)

type Config struct {
	Database struct {
		Master dto.DatabaseConfig `mapstructure:"master"`
		Worker dto.DatabaseConfig `mapstructure:"worker"`
	} `mapstructure:"database"`

	Oidc struct {
		Realm        string `mapstructure:"realm"`
		ClientID     string `mapstructure:"client_id"`
		ClientSecret string `mapstxructure:"client_secret"`
		Issuer       string `mapstructure:"issuer"`
	}

	Redis struct {
		Host     string `mapstructure:"host"`
		Password string `mapstructure:"password"`
		DB       int    `mapstructure:"db"`
	}
}

func NewAppMetadata() *dto.AppMetadata {
	return &dto.AppMetadata{
		AppName:     "gas",
		Instance:    uuid.New().String(),
		Port:        8080,
		ContextPath: "",
	}
}
func NewConfig(metadata *dto.AppMetadata) *Config {

	cfg, err := utils.LoadConfig[Config](metadata.AppName)

	if err != nil {
		panic("Load config failed: " + err.Error())
	}

	log.Println("Config loaded for app:", metadata.AppName, "instance:", metadata.Instance)

	return cfg
}
