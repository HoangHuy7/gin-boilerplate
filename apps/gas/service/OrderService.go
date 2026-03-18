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
	return database.WithTenant(
		s.db.Mekyra_db,
		ctx,
		tenancy,
		func(tx *gorm.DB) error {
			if err := tx.Create(order).Error; err != nil {
				return err
			}
			for _, item := range items {
				item.OrderId = order.Id
				if err := tx.Create(item).Error; err != nil {
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
