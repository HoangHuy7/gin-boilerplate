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

type DebtService struct {
	db     *database.DataSources
	logger *zap.Logger
}

func NewDebtService(db *database.DataSources, lg *zap.Logger) *DebtService {
	return &DebtService{
		db:     db,
		logger: lg,
	}
}

func (s *DebtService) FindAll(ctx context.Context) *[]*mekyra_db.Mkrtb_DebtTransaction {
	var list []*mekyra_db.Mkrtb_DebtTransaction
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

func (s *DebtService) FindByID(ctx context.Context, id string) *mekyra_db.Mkrtb_DebtTransaction {
	var debt mekyra_db.Mkrtb_DebtTransaction
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
			return tx.First(&debt, "id = ?", id).Error
		},
	)
	return &debt
}

func (s *DebtService) Create(ctx context.Context, debt *mekyra_db.Mkrtb_DebtTransaction) error {
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
			return tx.Create(debt).Error
		},
	)
}

func (s *DebtService) Delete(ctx context.Context, id string) error {
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
			return tx.Delete(&mekyra_db.Mkrtb_DebtTransaction{}, "id = ?", id).Error
		},
	)
}
