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
}

func NewAppMetadata() *dto.AppMetadata {
	return &dto.AppMetadata{
		AppName:     "iam",
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
