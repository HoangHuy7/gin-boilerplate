package v1

import (
	"monorepo/internal/base/routerx"
	"monorepo/internal/dto"

	"github.com/gin-gonic/gin"
)

type HelloController struct {
	Metadata dto.Metadata
}

func (this *HelloController) GetMetadata() *dto.Metadata {
	return &this.Metadata
}

func NewHelloController() *HelloController {
	return &HelloController{
		Metadata: dto.Metadata{
			Path:          "",
			Tag:           "Hello Controller",
			Endpoints:     []dto.OpenEndpoint{},
			EnableOpenAPI: true,
		},
	}
}

func (this *HelloController) Register(rg *routerx.Routerx) {
	rg.GET(dto.OpenEndpoint{
		Path:        "/hello",
		Handler:     this.Hello,
		Summary:     "Hello World Summary",
		Description: "Hello World Description",
	})

}

func (this *HelloController) Hello(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
