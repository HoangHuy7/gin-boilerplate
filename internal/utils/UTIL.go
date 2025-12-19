// hoanghuy7 from Vietnamese with love!

package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"monorepo/internal/dto"
	"os"
	"regexp"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func LoadConfig[T any](appName string) (*T, error) {
	v := viper.New()
	v.SetConfigName("application")
	v.SetConfigType("yaml")
	v.AddConfigPath(fmt.Sprintf("configs/%s", appName))
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Expand environment variables in the config file
	// This allows using ${VAR} syntax in application.yaml
	configKeys := v.AllKeys()
	for _, key := range configKeys {
		val := v.Get(key)
		if str, ok := val.(string); ok {
			v.Set(key, os.ExpandEnv(str))
		}
	}

	log.Println("Config file used:", v.ConfigFileUsed())

	var cfg T
	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatalf("Error parsing config: %v", err)
	}

	return &cfg, nil
}

func Connect(dbCfg *dto.DatabaseConfig, logger *zap.Logger) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		dbCfg.Host,
		dbCfg.User,
		dbCfg.Password,
		dbCfg.DBName,
		dbCfg.Port,
		dbCfg.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		return nil, err
	}
	sqlDB, err2 := db.DB()
	if err2 != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(30)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	sqlDB.SetConnMaxIdleTime(2 * time.Minute)
	logger.Info("Connect to database",
		zap.String("host", dbCfg.Host),
		zap.String("user", dbCfg.User),
		zap.String("dbname", dbCfg.DBName),
		zap.Int("port", dbCfg.Port),
		zap.String("sslmode", dbCfg.SSLMode),
	)
	return db, nil
}

func GinPathToOpenAPI(path string) string {
	re := regexp.MustCompile(`:([a-zA-Z0-9_]+)`)
	return re.ReplaceAllString(path, `{$1}`)
}

func NVL[T any](nvl *T, vcl T) T {
	if nvl != nil {
		return *nvl
	}
	return vcl
}

func ValueOr[T any](condition bool, trueVal T, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

func ToJSON[T any](v T) ([]byte, error) {
	return json.Marshal(v)
}

func FromJSON[T any](value string, out *T) error {
	return json.Unmarshal([]byte(value), out)
}
