package domain

import (
	"context"
	"monorepo/apps/gas/app/database"
	"monorepo/apps/gas/service"
	"monorepo/internal/logger"
	"monorepo/shares/entities/workerdb/view"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type MenuDomainService struct {
	logger      *zap.Logger
	menuService *service.MenuService
	db          *database.DataSources
}

func NewMenuDomainService(logger *logger.GoLogger, menuService *service.MenuService, db *database.DataSources) *MenuDomainService {
	return &MenuDomainService{
		logger:      logger.Zap,
		menuService: menuService,
		db:          db,
	}
}

func (this *MenuDomainService) GetMenuTree(context context.Context, user map[string]string) []view.Vw_UserMenu {
	var list = make([]view.Vw_UserMenu, 0)

	this.db.Worker.Transaction(func(tx *gorm.DB) error {
		list = this.menuService.GetMenuTree(tx)

		return nil
	})

	return list
}
