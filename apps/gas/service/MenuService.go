package service

import (
	"monorepo/shares/entities/workerdb/view"

	"gorm.io/gorm"
)

type MenuService struct {
}

func NewMenuService() *MenuService {
	return &MenuService{}
}

func (this MenuService) GetMenuTree(tx *gorm.DB) []view.Vw_UserMenu {
	list := make([]view.Vw_UserMenu, 0)
	tx.Find(&list)
	return list
}
