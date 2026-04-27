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

	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
			if debt.Amount.LessThanOrEqual(decimal.Zero) {
				return &exception.AppError{
					Code:    "INVALID_AMOUNT",
					Message: "Số tiền phải lớn hơn 0",
				}
			}
			debtType := strings.ToLower(strings.TrimSpace(debt.Type))
			if debtType != "borrow" && debtType != "pay" {
				return &exception.AppError{
					Code:    "INVALID_DEBT_TYPE",
					Message: "Loại công nợ chỉ hỗ trợ borrow hoặc pay",
				}
			}

			var customer mekyra_db.Mkrtb_Customer
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
				First(&customer, "id = ?", debt.CustomerId).Error; err != nil {
				return &exception.AppError{
					Code:    "CUSTOMER_NOT_FOUND",
					Message: "Khách hàng không tồn tại",
				}
			}

			delta := debt.Amount
			if debtType == "pay" {
				delta = debt.Amount.Neg()
			}
			nextDebt := customer.TotalDebt.Add(delta)
			if nextDebt.LessThan(decimal.Zero) {
				return &exception.AppError{
					Code:    "OVERPAY_NOT_ALLOWED",
					Message: "Số tiền trả nợ vượt quá công nợ hiện tại",
				}
			}

			debt.Type = debtType
			if err := tx.Model(&mekyra_db.Mkrtb_Customer{}).
				Where("id = ?", customer.Id).
				Update("total_debt", nextDebt).Error; err != nil {
				return err
			}

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
			var debt mekyra_db.Mkrtb_DebtTransaction
			if err := tx.First(&debt, "id = ?", id).Error; err != nil {
				return err
			}

			var customer mekyra_db.Mkrtb_Customer
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
				First(&customer, "id = ?", debt.CustomerId).Error; err != nil {
				return &exception.AppError{
					Code:    "CUSTOMER_NOT_FOUND",
					Message: "Khách hàng không tồn tại",
				}
			}

			delta := debt.Amount
			if strings.ToLower(strings.TrimSpace(debt.Type)) == "borrow" {
				delta = debt.Amount.Neg()
			}
			nextDebt := customer.TotalDebt.Add(delta)
			if nextDebt.LessThan(decimal.Zero) {
				return &exception.AppError{
					Code:    "INVALID_DEBT_STATE",
					Message: "Không thể xóa giao dịch vì sẽ làm dữ liệu công nợ âm",
				}
			}

			if err := tx.Model(&mekyra_db.Mkrtb_Customer{}).
				Where("id = ?", customer.Id).
				Update("total_debt", nextDebt).Error; err != nil {
				return err
			}

			return tx.Delete(&mekyra_db.Mkrtb_DebtTransaction{}, "id = ?", id).Error
		},
	)
}
