package menu

import (
	"monorepo/internal/base/routerx"
	"monorepo/internal/dto"
	"monorepo/internal/logger"

	"go.uber.org/zap"
)

type MenuControllerV1 struct {
	logger *zap.Logger
}

func (this *MenuControllerV1) GetMetadata() *dto.Metadata {
	return &dto.Metadata{
		Path:          "/menu",
		Version:       "v1",
		Tag:           "Menu Controller",
		Endpoints:     []dto.OpenEndpoint{},
		EnableOpenAPI: true,
		IsNotAuth:     false,
	}
}

func NewMenuControllerV1(lg *logger.GoLogger) *MenuControllerV1 {
	return &MenuControllerV1{
		logger: lg.Zap,
	}
}

func (this *MenuControllerV1) Register(routerx *routerx.Routerx) {
	routerx.GET(dto.OpenEndpoint{
		Path:        "",
		Handler:     nil,
		Summary:     "",
		Description: "",
	})
}
