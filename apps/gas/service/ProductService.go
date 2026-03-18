package service

import (
	"context"
	"monorepo/apps/gas/app/database"
	"monorepo/shares/entities/mekyra_db"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductService struct {
	db     *database.DataSources
	logger *zap.Logger
}

func NewProductService(db *database.DataSources, lg *zap.Logger) *ProductService {
	return &ProductService{
		db:     db,
		logger: lg,
	}
}

func (p *ProductService) FindAll(ctx context.Context) *[]*mekyra_db.Mkrtb_Product {

	var list []*mekyra_db.Mkrtb_Product

	_ = database.WithTenant(
		p.db.Mekyra_db,
		ctx,
		"vattumyloc",
		func(tx *gorm.DB) error {
			return tx.Find(&list).Error
		},
	)

	return &list
}
