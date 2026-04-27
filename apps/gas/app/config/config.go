// hoanghuy7 from Vietnamese with love!

package config

import (
	"monorepo/internal/dto"
	"monorepo/internal/logger"
	"monorepo/internal/utils"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Config struct {
	Database struct {
		Mekyra_db dto.DatabaseConfig `mapstructure:"mekyra_db"`
		//Worker dto.DatabaseConfig `mapstructure:"worker"`
	} `mapstructure:"database"`

	// Casdoor configuration for multi-organization support
	Casdoor struct {
		Organizations map[string]dto.CasdoorOrgConfig `mapstructure:"organizations"`
	} `mapstructure:"casdoor"`

	Tenancies map[string]string `mapstructure:"tenancies"`

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

var appConfig *Config

func GetTenancy(key string) (string, bool) {
	if appConfig == nil {
		panic("config not initialized")
	}

	val, ok := appConfig.Tenancies[key]
	return val, !ok
}

func NewConfig(metadata *dto.AppMetadata, lg *logger.GoLogger) *Config {

	cfg, err := utils.LoadConfig[Config](metadata.AppName)

	if err != nil {
		panic("Load config failed: " + err.Error())
	}
	lg.INFO(
		"Config loaded successfully",
		zap.String("app", metadata.AppName),
		zap.String("instance", metadata.Instance),
	)
	appConfig = cfg
	return cfg
}
