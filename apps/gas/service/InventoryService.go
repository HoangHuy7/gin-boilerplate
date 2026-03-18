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
			return tx.Create(log).Error
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
