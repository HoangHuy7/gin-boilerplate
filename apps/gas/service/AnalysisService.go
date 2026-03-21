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

type AnalysisService struct {
	db     *database.DataSources
	logger *zap.Logger
}

func NewAnalysisService(db *database.DataSources, lg *zap.Logger) *AnalysisService {
	return &AnalysisService{
		db:     db,
		logger: lg,
	}
}

func (s *AnalysisService) GetSalesSummary(
	ctx context.Context,
	filter *model.SalesFilter,
	offset int,
	limit int,
) ([]*mekyra_db.Vw_Sales_Summary, int64, error) {
	var list []*mekyra_db.Vw_Sales_Summary
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
			query := tx.Model(&mekyra_db.Vw_Sales_Summary{})

			if filter != nil {
				if filter.FromDate != nil {
					query = query.Where("sale_date >= ?", *filter.FromDate)
				}
				if filter.ToDate != nil {
					query = query.Where("sale_date <= ?", *filter.ToDate)
				}
			}

			// Count total records
			if err := query.Count(&total).Error; err != nil {
				return err
			}

			return query.
				Order("sale_date desc").
				Offset(offset).
				Limit(limit).
				Find(&list).Error
		},
	)

	if errDB != nil {
		return nil, 0, errDB
	}

	return list, total, nil
}

func (s *AnalysisService) GetSalesByDay(
	ctx context.Context,
	filter *model.SalesFilter,
	offset int,
	limit int,
) ([]*mekyra_db.Vw_Sales_Summary, int64, error) {
	var list []*mekyra_db.Vw_Sales_Summary
	var total int64

	org, err := utils.GetOrg(ctx)
	if err {
		return nil, 0, fmt.Errorf("failed to get organization from context")
	}
	tenancy, err := config.GetTenancy(org)
	if err {
		return nil, 0, fmt.Errorf("failed to get tenancy")
	}

	errDB := database.WithTenant(
		s.db.Mekyra_db,
		ctx,
		tenancy,
		func(tx *gorm.DB) error {
			query := tx.Model(&mekyra_db.Vw_Sales_Summary{}).
				Select("sale_date, sum(quantity) as quantity, sum(total) as total").
				Group("sale_date")

			if filter != nil {
				if filter.FromDate != nil {
					query = query.Where("sale_date >= ?", *filter.FromDate)
				}
				if filter.ToDate != nil {
					query = query.Where("sale_date <= ?", *filter.ToDate)
				}
			}

			// Count total groups
			if err := tx.Table("(?) as subquery", query).Count(&total).Error; err != nil {
				return err
			}

			return query.
				Order("sale_date desc").
				Offset(offset).
				Limit(limit).
				Find(&list).Error
		},
	)

	return list, total, errDB
}

func (s *AnalysisService) GetSalesByMonth(
	ctx context.Context,
	filter *model.SalesFilter,
	offset int,
	limit int,
) ([]*mekyra_db.Vw_Sales_Summary, int64, error) {
	var list []*mekyra_db.Vw_Sales_Summary
	var total int64

	org, err := utils.GetOrg(ctx)
	if err {
		return nil, 0, fmt.Errorf("failed to get organization from context")
	}
	tenancy, err := config.GetTenancy(org)
	if err {
		return nil, 0, fmt.Errorf("failed to get tenancy")
	}

	errDB := database.WithTenant(
		s.db.Mekyra_db,
		ctx,
		tenancy,
		func(tx *gorm.DB) error {
			query := tx.Model(&mekyra_db.Vw_Sales_Summary{}).
				Select("sale_month, sum(quantity) as quantity, sum(total) as total").
				Group("sale_month")

			if filter != nil {
				if filter.FromDate != nil {
					query = query.Where("sale_date >= ?", *filter.FromDate)
				}
				if filter.ToDate != nil {
					query = query.Where("sale_date <= ?", *filter.ToDate)
				}
			}

			// Count total groups
			if err := tx.Table("(?) as subquery", query).Count(&total).Error; err != nil {
				return err
			}

			return query.
				Order("sale_month desc").
				Offset(offset).
				Limit(limit).
				Find(&list).Error
		},
	)

	return list, total, errDB
}

func (s *AnalysisService) GetSalesByYear(
	ctx context.Context,
	filter *model.SalesFilter,
	offset int,
	limit int,
) ([]*mekyra_db.Vw_Sales_Summary, int64, error) {
	var list []*mekyra_db.Vw_Sales_Summary
	var total int64

	org, err := utils.GetOrg(ctx)
	if err {
		return nil, 0, fmt.Errorf("failed to get organization from context")
	}
	tenancy, err := config.GetTenancy(org)
	if err {
		return nil, 0, fmt.Errorf("failed to get tenancy")
	}

	errDB := database.WithTenant(
		s.db.Mekyra_db,
		ctx,
		tenancy,
		func(tx *gorm.DB) error {
			query := tx.Model(&mekyra_db.Vw_Sales_Summary{}).
				Select("sale_year, sum(quantity) as quantity, sum(total) as total").
				Group("sale_year")

			if filter != nil {
				if filter.FromDate != nil {
					query = query.Where("sale_date >= ?", *filter.FromDate)
				}
				if filter.ToDate != nil {
					query = query.Where("sale_date <= ?", *filter.ToDate)
				}
			}

			// Count total groups
			if err := tx.Table("(?) as subquery", query).Count(&total).Error; err != nil {
				return err
			}

			return query.
				Order("sale_year desc").
				Offset(offset).
				Limit(limit).
				Find(&list).Error
		},
	)

	return list, total, errDB
}
