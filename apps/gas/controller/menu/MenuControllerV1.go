package menu

import (
	"monorepo/apps/gas/domain"
	"monorepo/internal/base/routerx"
	"monorepo/internal/dto"
	"monorepo/internal/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MenuControllerV1 struct {
	logger     *zap.Logger
	menuDomain *domain.MenuDomainService
}

func (this *MenuControllerV1) GetMetadata() *dto.Metadata {
	return &dto.Metadata{
		Path:          "/menu",
		Version:       "/v1",
		Tag:           "Menu Controller",
		Endpoints:     []dto.OpenEndpoint{},
		EnableOpenAPI: true,
		IsNotAuth:     true,
	}
}

func NewMenuControllerV1(lg *logger.GoLogger, menuDomain *domain.MenuDomainService) *MenuControllerV1 {
	return &MenuControllerV1{
		logger: lg.Zap, menuDomain: menuDomain,
	}
}

func (this *MenuControllerV1) Register(routerx *routerx.Routerx) {
	routerx.GET(dto.OpenEndpoint{
		Path:        "",
		Handler:     this.getMenu,
		Summary:     "",
		Description: "",
	})
}

func (this *MenuControllerV1) getMenu(context *gin.Context) {

	user := map[string]string{}
	tree := this.menuDomain.GetMenuTree(context.Request.Context(), user)

	context.JSON(200, gin.H{
		"data": tree,
	})
}
