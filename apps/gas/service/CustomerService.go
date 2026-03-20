package service

import (
	"context"
	"fmt"
	"monorepo/apps/gas/app/config"
	"monorepo/apps/gas/app/database"
	"monorepo/apps/gas/graph/model"
	"monorepo/shares/entities/mekyra_db"
	"monorepo/shares/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CustomerService struct {
	db     *database.DataSources
	logger *zap.Logger
}

func NewCustomerService(db *database.DataSources, lg *zap.Logger) *CustomerService {
	return &CustomerService{
		db:     db,
		logger: lg,
	}
}

func (s *CustomerService) FindAll(ctx context.Context) *[]*mekyra_db.Mkrtb_Customer {
	var list []*mekyra_db.Mkrtb_Customer
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
func (s *CustomerService) FindWithPagination(
	ctx context.Context,
	filter *model.CustomerFilter,
	offset int,
	limit int,
) ([]*mekyra_db.Mkrtb_Customer, int64, error) {

	var list []*mekyra_db.Mkrtb_Customer
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

			query := tx.Model(&mekyra_db.Mkrtb_Customer{})

			// =========================
			// FILTER
			// =========================
			if filter != nil {

				if filter.Search != nil && *filter.Search != "" {
					search := "%" + *filter.Search + "%"
					query = query.Where("name ILIKE ? OR phone ILIKE ?", search, search)
				}

				if filter.Phone != nil && *filter.Phone != "" {
					query = query.Where("phone = ?", *filter.Phone)
				}
			}

			// 👉 count
			if err := query.Count(&total).Error; err != nil {
				return err
			}

			// 👉 data
			if err := query.
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
func (s *CustomerService) FindByID(ctx context.Context, id string) *mekyra_db.Mkrtb_Customer {
	var customer mekyra_db.Mkrtb_Customer
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
			return tx.First(&customer, "id = ?", id).Error
		},
	)
	return &customer
}

func (s *CustomerService) Create(ctx context.Context, customer *mekyra_db.Mkrtb_Customer) error {
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
			return tx.Create(customer).Error
		},
	)
}

func (s *CustomerService) Update(ctx context.Context, customer *mekyra_db.Mkrtb_Customer) error {
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
			return tx.Save(customer).Error
		},
	)
}

func (s *CustomerService) Delete(ctx context.Context, id string) error {
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
			return tx.Delete(&mekyra_db.Mkrtb_Customer{}, "id = ?", id).Error
		},
	)
}
