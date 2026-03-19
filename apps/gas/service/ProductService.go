package service

import (
	"context"
	"fmt"
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
			return tx.Order("updated_at desc").Find(&list).Error
		},
	)

	return &list
}

func (p *ProductService) FindByID(ctx context.Context, id string) *mekyra_db.Mkrtb_Product {
	var product mekyra_db.Mkrtb_Product
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
			return tx.First(&product, "id = ?", id).Error
		},
	)
	return &product
}

func (p *ProductService) Create(ctx context.Context, product *mekyra_db.Mkrtb_Product) error {
	org, err := utils.GetOrg(ctx)
	if err {
		p.logger.Error("Failed to get organization from context")
		return fmt.Errorf("failed to get organization from context")
	}
	tenancy, err := config.GetTenancy(org)
	if err {
		p.logger.Error("Failed to get tenancy")
		return fmt.Errorf("failed to get tenancy")
	}
	return database.WithTenant(
		p.db.Mekyra_db,
		ctx,
		tenancy,
		func(tx *gorm.DB) error {
			return tx.Create(product).Error
		},
	)
}

func (p *ProductService) Update(ctx context.Context, product *mekyra_db.Mkrtb_Product) error {
	org, err := utils.GetOrg(ctx)
	if err {
		p.logger.Error("Failed to get organization from context")
		return fmt.Errorf("failed to get organization from context")
	}
	tenancy, err := config.GetTenancy(org)
	if err {
		p.logger.Error("Failed to get tenancy")
		return fmt.Errorf("failed to get tenancy")
	}
	return database.WithTenant(
		p.db.Mekyra_db,
		ctx,
		tenancy,
		func(tx *gorm.DB) error {
			return tx.Save(product).Error
		},
	)
}

func (p *ProductService) Delete(ctx context.Context, id string) error {
	org, err := utils.GetOrg(ctx)
	if err {
		p.logger.Error("Failed to get organization from context")
		return fmt.Errorf("failed to get organization from context")
	}
	tenancy, err := config.GetTenancy(org)
	if err {
		p.logger.Error("Failed to get tenancy")
		return fmt.Errorf("failed to get tenancy")
	}
	return database.WithTenant(
		p.db.Mekyra_db,
		ctx,
		tenancy,
		func(tx *gorm.DB) error {
			return tx.Delete(&mekyra_db.Mkrtb_Product{}, "id = ?", id).Error
		},
	)
}
