package service

import (
	"context"
	"fmt"
	"monorepo/apps/gas/app/config"
	"monorepo/apps/gas/app/database"
	"monorepo/apps/gas/graph/model"
	"monorepo/shares/entities/mekyra_db"
	"monorepo/shares/exception"
	"monorepo/shares/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OrderService struct {
	db     *database.DataSources
	logger *zap.Logger
}

func NewOrderService(db *database.DataSources, lg *zap.Logger) *OrderService {
	return &OrderService{
		db:     db,
		logger: lg,
	}
}

func (s *OrderService) FindAll(ctx context.Context) *[]*mekyra_db.Mkrtb_Order {
	var list []*mekyra_db.Mkrtb_Order
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
			return tx.Preload("Items").Find(&list).Error
		},
	)
	return &list
}

func (s *OrderService) FindWithPagination(
	ctx context.Context,
	filter *model.OrderFilter,
	offset int,
	limit int,
) ([]*mekyra_db.Mkrtb_Order, int64, error) {

	var list []*mekyra_db.Mkrtb_Order
	var total int64

	org, err := utils.GetOrg(ctx)
	if err {
		s.logger.Error("Failed to get organization from context")
		return nil, 0, fmt.Errorf("failed to get organization from context")
	}

	tenancy, err := config.GetTenancy(org)
	if err {
		s.logger.Error("Failed to get tenancy")
		return nil, 0, fmt.Errorf("failed to get tenancy")
	}

	errDB := database.WithTenant(
		s.db.Mekyra_db,
		ctx,
		tenancy,
		func(tx *gorm.DB) error {

			query := tx.Model(&mekyra_db.Mkrtb_Order{})

			// =========================
			// FILTER
			// =========================
			if filter != nil {
				if filter.Status != nil && *filter.Status != "" {
					query = query.Where("status = ?", *filter.Status)
				}
				if filter.CustomerID != nil && *filter.CustomerID != "" {
					query = query.Where("customer_id = ?", *filter.CustomerID)
				}
				if filter.FromDate != nil {
					query = query.Where("created_at >= ?", *filter.FromDate)
				}
				if filter.ToDate != nil {
					query = query.Where("created_at <= ?", *filter.ToDate)
				}
			}

			// 👉 count
			if err := query.Count(&total).Error; err != nil {
				return err
			}

			// 👉 data
			if err := query.
				Preload("Items").
				Order("created_at desc").
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

func (s *OrderService) FindByID(ctx context.Context, id string) *mekyra_db.Mkrtb_Order {
	var order mekyra_db.Mkrtb_Order
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
			return tx.Preload("Items").First(&order, "id = ?", id).Error
		},
	)
	return &order
}

func (s *OrderService) Create(ctx context.Context, order *mekyra_db.Mkrtb_Order, items []*mekyra_db.Mkrtb_OrderItem) error {
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
	order.Code = fmt.Sprintf("ORD-%s", utils.GenerateUUID())
	return database.WithTenant(
		s.db.Mekyra_db,
		ctx,
		tenancy,
		func(tx *gorm.DB) error {
			// Validate stock quantity for all items before creating order
			for _, item := range items {
				var product mekyra_db.Mkrtb_Product
				if err := tx.First(&product, "id = ?", item.ProductId).Error; err != nil {
					//return fmt.Errorf("product not found: %w", err)
					return &exception.AppError{
						Code:    "PRODUCT_NOT_FOUND",
						Message: "Sản phẩm không tồn tại",
					}
				}
				if product.StockQuantity < item.Quantity {
					return &exception.AppError{
						Code: "INSUFFICIENT_STOCK",
						Message: fmt.Sprintf("Số lượng hàng tồn kho không đủ cho sản phẩm %s: yêu cầu %d, còn %d",
							product.Name, item.Quantity, product.StockQuantity),
					}
				}
			}

			if err := tx.Create(order).Error; err != nil {
				return err
			}
			for _, item := range items {
				item.OrderId = order.Id
				if err := tx.Create(item).Error; err != nil {
					return err
				}

				// Subtract stock quantity
				if err := tx.Model(&mekyra_db.Mkrtb_Product{}).
					Where("id = ?", item.ProductId).
					Update("stock_quantity", gorm.Expr("stock_quantity - ?", item.Quantity)).Error; err != nil {
					return err
				}

				// Create inventory log
				invLog := &mekyra_db.Mkrtb_InventoryLog{
					ProductId: item.ProductId,
					Type:      "sale",
					Quantity:  -item.Quantity,
					Note:      fmt.Sprintf("Order %s", order.Code),
				}
				if err := tx.Create(invLog).Error; err != nil {
					return err
				}
			}
			return nil
		},
	)
}

func (s *OrderService) Update(ctx context.Context, order *mekyra_db.Mkrtb_Order) error {
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
			return tx.Save(order).Error
		},
	)
}

func (s *OrderService) Delete(ctx context.Context, id string) error {
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
			if err := tx.Delete(&mekyra_db.Mkrtb_OrderItem{}, "order_id = ?", id).Error; err != nil {
				return err
			}
			return tx.Delete(&mekyra_db.Mkrtb_Order{}, "id = ?", id).Error
		},
	)
}

func (s *OrderService) GetOrderItems(ctx context.Context, orderID string) *[]*mekyra_db.Mkrtb_OrderItem {
	var items []*mekyra_db.Mkrtb_OrderItem
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
			return tx.Where("order_id = ?", orderID).Find(&items).Error
		},
	)
	return &items
}
