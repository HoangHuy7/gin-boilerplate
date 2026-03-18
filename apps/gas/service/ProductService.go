package service

import (
	"context"
	"monorepo/apps/gas/app/config"
	"monorepo/apps/gas/app/database"
	"monorepo/shares/entities/mekyra_db"
	"monorepo/shares/utils"

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
	org, err := utils.GetOrg(ctx)
	if err {
		p.logger.Error("Failed to get organization from context")
		return nil
	}
	tenancy, err := config.GetTenancy(org)
	if err {
		p.logger.Error("Failed to get tenancy")
		return nil
	}
	_ = database.WithTenant(
		p.db.Mekyra_db,
		ctx,
		tenancy,
		func(tx *gorm.DB) error {
			return tx.Find(&list).Error
		},
	)

	return &list
}
