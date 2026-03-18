package service

import (
	"context"
	"fmt"
	"monorepo/apps/gas/app/config"
	"monorepo/apps/gas/app/database"
	"monorepo/shares/entities/mekyra_db"
	"monorepo/shares/utils"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DeliveryService struct {
	db     *database.DataSources
	logger *zap.Logger
}

func NewDeliveryService(db *database.DataSources, lg *zap.Logger) *DeliveryService {
	return &DeliveryService{
		db:     db,
		logger: lg,
	}
}

func (s *DeliveryService) FindAll(ctx context.Context) *[]*mekyra_db.Mkrtb_Delivery {
	var list []*mekyra_db.Mkrtb_Delivery
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

func (s *DeliveryService) FindByID(ctx context.Context, id string) *mekyra_db.Mkrtb_Delivery {
	var delivery mekyra_db.Mkrtb_Delivery
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
			return tx.First(&delivery, "id = ?", id).Error
		},
	)
	return &delivery
}

func (s *DeliveryService) FindTodayDeliveries(ctx context.Context) *[]*mekyra_db.Mkrtb_Delivery {
	var list []*mekyra_db.Mkrtb_Delivery
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
	today := time.Now().Format("2006-01-02")
	_ = database.WithTenant(
		s.db.Mekyra_db,
		ctx,
		tenancy,
		func(tx *gorm.DB) error {
			return tx.Where("DATE(delivery_date) = ?", today).Find(&list).Error
		},
	)
	return &list
}

func (s *DeliveryService) Create(ctx context.Context, delivery *mekyra_db.Mkrtb_Delivery) error {
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
			return tx.Create(delivery).Error
		},
	)
}

func (s *DeliveryService) Update(ctx context.Context, delivery *mekyra_db.Mkrtb_Delivery) error {
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
			return tx.Save(delivery).Error
		},
	)
}

func (s *DeliveryService) Delete(ctx context.Context, id string) error {
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
			return tx.Delete(&mekyra_db.Mkrtb_Delivery{}, "id = ?", id).Error
		},
	)
}
