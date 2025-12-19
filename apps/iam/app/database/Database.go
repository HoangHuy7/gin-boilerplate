// hoanghuy7 from Vietnamese with love!

package database

import (
	"context"
	"fmt"
	"monorepo/apps/iam/app/config"
	"monorepo/internal/logger"
	"monorepo/internal/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DataSources struct {
	Master *gorm.DB
	Worker *gorm.DB
	logger *zap.Logger
}

func NewDataSources(cfg *config.Config, gLogger *logger.GoLogger) *DataSources {
	master, err := utils.Connect(&cfg.Database.Master, gLogger.Zap)
	if err != nil {
		panic(err)
	}

	worker, err := utils.Connect(&cfg.Database.Worker, gLogger.Zap)
	if err != nil {
		panic(err)
	}

	return &DataSources{
		Master: master,
		Worker: worker,
	}
}

type ctxKey string

const SchemaKey ctxKey = "schema"

func WithSchema(ctx context.Context, schema string) context.Context {
	return context.WithValue(ctx, SchemaKey, schema)
}
func WithTenantTx(
	db *gorm.DB,
	ctx context.Context,
	schema string,
	fn func(tx *gorm.DB) error,
) error {

	return db.WithContext(
		WithSchema(ctx, schema),
	).Transaction(fn)
}

func WithTenant(
	db *gorm.DB,
	ctx context.Context,
	schema string,
	fn func(tx *gorm.DB) error,
) error {

	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(
			fmt.Sprintf(`SET LOCAL search_path TO "%s"`, schema),
		).Error; err != nil {
			return err
		}
		return fn(tx)
	})
}
