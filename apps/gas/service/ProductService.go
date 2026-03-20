package service

import (
	"context"
	"fmt"
	"monorepo/apps/gas/app/config"
	"monorepo/apps/gas/app/database"
	"monorepo/apps/gas/graph/model"
	"monorepo/shares/entities/mekyra_db"
	"monorepo/shares/utils"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductService struct {
	db     *database.DataSources
	logger *zap.Logger
	logSvc *LogService
}

func NewProductService(db *database.DataSources, logSvc *LogService, lg *zap.Logger) *ProductService {
	return &ProductService{
		db:     db,
		logger: lg,
		logSvc: logSvc,
	}
}

// SetLogService sets the log service for audit logging
func (p *ProductService) SetLogService(logSvc *LogService) {
	p.logSvc = logSvc
}

func (p *ProductService) FindAll(ctx context.Context) ([]*mekyra_db.Mkrtb_Product, error) {
	var list []*mekyra_db.Mkrtb_Product
	org, err := utils.GetOrg(ctx)
	if err {
		p.logger.Error("Failed to get organization from context")
		return nil, fmt.Errorf("failed to get organization from context")
	}
	tenancy, err := config.GetTenancy(org)
	if err {
		p.logger.Error("Failed to get tenancy")
		return nil, fmt.Errorf("failed to get tenancy")
	}
	_ = database.WithTenant(
		p.db.Mekyra_db,
		ctx,
		tenancy,
		func(tx *gorm.DB) error {
			return tx.Order("updated_at desc").Find(&list).Error
		},
	)

	return list, nil
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
	org, getOrgErr := utils.GetOrg(ctx)
	if getOrgErr {
		p.logger.Error("Failed to get organization from context")
		return fmt.Errorf("failed to get organization from context")
	}
	tenancy, tenancyErr := config.GetTenancy(org)
	if tenancyErr {
		p.logger.Error("Failed to get tenancy")
		return fmt.Errorf("failed to get tenancy")
	}

	err := database.WithTenant(
		p.db.Mekyra_db,
		ctx,
		tenancy,
		func(tx *gorm.DB) error {
			return tx.Create(product).Error
		},
	)
	if err != nil {
		if p.logSvc != nil {
			p.logSvc.LogError(ctx, "CREATE", "PRODUCT", product, err.Error())
		}
		return err
	}

	// Log the create action
	if p.logSvc != nil {
		p.logSvc.LogSuccess(ctx, "CREATE", "PRODUCT", nil, product)
	}

	return nil
}

func (p *ProductService) Update(ctx context.Context, product *mekyra_db.Mkrtb_Product) error {
	org, getOrgErr := utils.GetOrg(ctx)
	if getOrgErr {
		p.logger.Error("Failed to get organization from context")
		return fmt.Errorf("failed to get organization from context")
	}
	tenancy, tenancyErr := config.GetTenancy(org)
	if tenancyErr {
		p.logger.Error("Failed to get tenancy")
		return fmt.Errorf("failed to get tenancy")
	}

	// Get old data before update
	var oldProduct mekyra_db.Mkrtb_Product
	err := database.WithTenant(
		p.db.Mekyra_db,
		ctx,
		tenancy,
		func(tx *gorm.DB) error {
			return tx.First(&oldProduct, "id = ?", product.Id).Error
		},
	)
	if err != nil {
		if p.logSvc != nil {
			p.logSvc.LogError(ctx, "UPDATE", "PRODUCT", nil, err.Error())
		}
		return err
	}

	err = database.WithTenant(
		p.db.Mekyra_db,
		ctx,
		tenancy,
		func(tx *gorm.DB) error {
			return tx.Save(product).Error
		},
	)
	if err != nil {
		if p.logSvc != nil {
			p.logSvc.LogError(ctx, "UPDATE", "PRODUCT", &oldProduct, err.Error())
		}
		return err
	}

	// Log the update action
	if p.logSvc != nil {
		p.logSvc.LogSuccess(ctx, "UPDATE", "PRODUCT", &oldProduct, product)
	}

	return nil
}

func (p *ProductService) Delete(ctx context.Context, id string) error {
	org, getOrgErr := utils.GetOrg(ctx)
	if getOrgErr {
		p.logger.Error("Failed to get organization from context")
		return fmt.Errorf("failed to get organization from context")
	}
	tenancy, tenancyErr := config.GetTenancy(org)
	if tenancyErr {
		p.logger.Error("Failed to get tenancy")
		return fmt.Errorf("failed to get tenancy")
	}

	// Get old data before delete
	var oldProduct mekyra_db.Mkrtb_Product
	err := database.WithTenant(
		p.db.Mekyra_db,
		ctx,
		tenancy,
		func(tx *gorm.DB) error {
			return tx.First(&oldProduct, "id = ?", id).Error
		},
	)
	if err != nil {
		if p.logSvc != nil {
			p.logSvc.LogError(ctx, "DELETE", "PRODUCT", nil, err.Error())
		}
		return err
	}

	err = database.WithTenant(
		p.db.Mekyra_db,
		ctx,
		tenancy,
		func(tx *gorm.DB) error {
			return tx.Delete(&mekyra_db.Mkrtb_Product{}, "id = ?", id).Error
		},
	)
	if err != nil {
		if p.logSvc != nil {
			p.logSvc.LogError(ctx, "DELETE", "PRODUCT", &oldProduct, err.Error())
		}
		return err
	}

	// Log the delete action
	if p.logSvc != nil {
		p.logSvc.LogSuccess(ctx, "DELETE", "PRODUCT", &oldProduct, nil)
	}

	return nil
}
func (p *ProductService) FindWithPagination(
	ctx context.Context,
	filter *model.ProductFilter,
	offset int,
	limit int,
) ([]*mekyra_db.Mkrtb_Product, int64, error) {

	var list []*mekyra_db.Mkrtb_Product
	var total int64

	org, err := utils.GetOrg(ctx)
	if err {
		p.logger.Error("Failed to get organization from context")
		return nil, 0, fmt.Errorf("failed to get organization from context")
	}

	tenancy, err := config.GetTenancy(org)
	if err {
		p.logger.Error("Failed to get tenancy")
		return nil, 0, fmt.Errorf("failed to get tenancy")
	}

	errDB := database.WithTenant(
		p.db.Mekyra_db,
		ctx,
		tenancy,
		func(tx *gorm.DB) error {

			// 👉 base query
			query := tx.Model(&mekyra_db.Mkrtb_Product{})

			// =========================
			// FILTER
			// =========================
			if filter != nil {

				if filter.Search != nil && *filter.Search != "" {
					search := "%" + *filter.Search + "%"
					query = query.Where("name ILIKE ? OR barcode ILIKE ?", search, search)
				}

				if filter.Barcode != nil && *filter.Barcode != "" {
					query = query.Where("barcode = ?", *filter.Barcode)
				}

				if filter.Category != nil && *filter.Category != "" {
					query = query.Where("category = ?", *filter.Category)
				}

				if filter.MinPrice != nil {
					query = query.Where("price >= ?", *filter.MinPrice)
				}

				if filter.MaxPrice != nil {
					query = query.Where("price <= ?", *filter.MaxPrice)
				}
			}

			// 👉 count (clone query để tránh bị dính limit/offset)
			if err := query.Count(&total).Error; err != nil {
				return err
			}

			// 👉 query data
			if err := query.
				Order("updated_at desc").
				Offset(offset).
				Limit(limit).
				Find(&list).Error; err != nil {
				return err
			}

			return nil
		},
	)

	if errDB != nil {
		return nil, 0, errDB
	}

	return list, total, nil
}

// Restock adds quantity to product stock
func (p *ProductService) Restock(ctx context.Context, productID uuid.UUID, quantity int) (*mekyra_db.Mkrtb_Product, error) {
	org, err := utils.GetOrg(ctx)
	if err {
		p.logger.Error("Failed to get organization from context")
		return nil, fmt.Errorf("failed to get organization from context")
	}
	tenancy, tenancyErr := config.GetTenancy(org)
	if tenancyErr {
		p.logger.Error("Failed to get tenancy")
		return nil, fmt.Errorf("failed to get tenancy")
	}

	var product mekyra_db.Mkrtb_Product
	errDB := database.WithTenant(
		p.db.Mekyra_db,
		ctx,
		tenancy,
		func(tx *gorm.DB) error {
			if err := tx.First(&product, "id = ?", productID).Error; err != nil {
				return err
			}

			// Lưu old data trước khi update
			oldStockQuantity := product.StockQuantity

			product.StockQuantity += quantity

			// Log the restock action với old_data
			if p.logSvc != nil {
				oldProduct := &mekyra_db.Mkrtb_Product{
					StockQuantity: oldStockQuantity,
				}
				p.logSvc.LogSuccess(ctx, "RESTOCK", "PRODUCT", oldProduct, map[string]interface{}{
					"product_id":     product.Id.String(),
					"quantity_added": quantity,
					"new_stock":      product.StockQuantity,
				})
			}

			return tx.Save(&product).Error
		},
	)

	if errDB != nil {
		if p.logSvc != nil {
			p.logSvc.LogError(ctx, "RESTOCK", "PRODUCT", nil, errDB.Error())
		}
		return nil, errDB
	}

	return &product, nil
}
