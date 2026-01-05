package domain

import (
	"monorepo/apps/gas/service"

	"go.uber.org/zap"
)

type MenuDomainService struct {
	logger      *zap.Logger
	menuService *service.MenuService
}

func NewMenuDomainService(logger *zap.Logger, menuService *service.MenuService) *MenuDomainService {
	return &MenuDomainService{
		logger:      logger,
		menuService: menuService,
	}
}
