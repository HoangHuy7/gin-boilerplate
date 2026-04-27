package service

import (
	"context"
	"fmt"
	"monorepo/apps/gas/app/config"
	"monorepo/apps/gas/app/database"
	"monorepo/shares/entities/mekyra_db"
	"monorepo/shares/exception"
	"monorepo/shares/utils"
	"strings"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type InventoryService struct {
	db     *database.DataSources
	logger *zap.Logger
}

func NewInventoryService(db *database.DataSources, lg *zap.Logger) *InventoryService {
	return &InventoryService{
		db:     db,
		logger: lg,
	}
}

func (s *InventoryService) FindAll(ctx context.Context) *[]*mekyra_db.Mkrtb_InventoryLog {
	var list []*mekyra_db.Mkrtb_InventoryLog
	org, err := utils.GetOrg(ctx)
	if err {
		s.logger.Error("Failed to get organization from context")
		return nil
	}
	tenancy, err := config.GetTenancy(org)
	if err {
		s.logger.Error("Failed to get tenancy")
		return nil
	}
	_ = database.WithTenant(
		s.db.Mekyra_db,
		ctx,
		tenancy,
		func(tx *gorm.DB) error {
			return tx.Find(&list).Error
		},
	)
	return &list
}

func (s *InventoryService) FindByID(ctx context.Context, id string) *mekyra_db.Mkrtb_InventoryLog {
	var log mekyra_db.Mkrtb_InventoryLog
	org, err := utils.GetOrg(ctx)
	if err {
		s.logger.Error("Failed to get organization from context")
		return nil
	}
	tenancy, err := config.GetTenancy(org)
	if err {
		s.logger.Error("Failed to get tenancy")
		return nil
	}
	_ = database.WithTenant(
		s.db.Mekyra_db,
		ctx,
		tenancy,
		func(tx *gorm.DB) error {
			return tx.First(&log, "id = ?", id).Error
		},
	)
	return &log
}

func (s *InventoryService) Create(ctx context.Context, log *mekyra_db.Mkrtb_InventoryLog) error {
	org, err := utils.GetOrg(ctx)
	if err {
		s.logger.Error("Failed to get organization from context")
		return fmt.Errorf("failed to get organization from context")
	}
	tenancy, err := config.GetTenancy(org)
	if err {
		s.logger.Error("Failed to get tenancy")
		return fmt.Errorf("failed to get tenancy")
	}
	return database.WithTenant(
		s.db.Mekyra_db,
		ctx,
		tenancy,
		func(tx *gorm.DB) error {
			logType := strings.ToLower(strings.TrimSpace(log.Type))
			if logType == "" {
				return &exception.AppError{
					Code:    "INVALID_INVENTORY_TYPE",
					Message: "Loại giao dịch kho không hợp lệ",
				}
			}
			if log.Quantity == 0 {
				return &exception.AppError{
					Code:    "INVALID_QUANTITY",
					Message: "Số lượng phải khác 0",
				}
			}

			var delta int
			switch logType {
			case "import":
				if log.Quantity < 0 {
					return &exception.AppError{
						Code:    "INVALID_IMPORT_QUANTITY",
						Message: "Nhập kho phải có số lượng dương",
					}
				}
				delta = log.Quantity
			case "sale":
				if log.Quantity > 0 {
					delta = -log.Quantity
				} else {
					delta = log.Quantity
				}
			case "adjust":
				delta = log.Quantity
			default:
				return &exception.AppError{
					Code:    "INVALID_INVENTORY_TYPE",
					Message: "Loại giao dịch kho chỉ hỗ trợ import, sale, adjust",
				}
			}

			var product mekyra_db.Mkrtb_Product
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
				First(&product, "id = ?", log.ProductId).Error; err != nil {
				return &exception.AppError{
					Code:    "PRODUCT_NOT_FOUND",
					Message: "Sản phẩm không tồn tại",
				}
			}

			newStock := product.StockQuantity + delta
			if newStock < 0 {
				return &exception.AppError{
					Code: "INSUFFICIENT_STOCK",
					Message: fmt.Sprintf(
						"Tồn kho không đủ cho sản phẩm %s: hiện có %d, thay đổi %d",
						product.Name,
						product.StockQuantity,
						delta,
					),
				}
			}

			log.Type = logType
			log.Quantity = delta

			if err := tx.Create(log).Error; err != nil {
				return err
			}

			return tx.Model(&mekyra_db.Mkrtb_Product{}).
				Where("id = ?", product.Id).
				Update("stock_quantity", newStock).Error
		},
	)
}

func (s *InventoryService) Delete(ctx context.Context, id string) error {
	org, err := utils.GetOrg(ctx)
	if err {
		s.logger.Error("Failed to get organization from context")
		return fmt.Errorf("failed to get organization from context")
	}
	tenancy, err := config.GetTenancy(org)
	if err {
		s.logger.Error("Failed to get tenancy")
		return fmt.Errorf("failed to get tenancy")
	}
	return database.WithTenant(
		s.db.Mekyra_db,
		ctx,
		tenancy,
		func(tx *gorm.DB) error {
			return tx.Delete(&mekyra_db.Mkrtb_InventoryLog{}, "id = ?", id).Error
		},
	)
}
