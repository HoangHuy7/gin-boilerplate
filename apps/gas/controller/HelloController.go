// hoanghuy7 from Vietnamese with love!

package controller

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
			Tag:           "Hello Controller",
			Version:       "/v1",
			Path:          "/hello",
			Endpoints:     []dto.OpenEndpoint{},
			EnableOpenAPI: true,
		},
	}
}

func (this *HelloController) Register(rg *routerx.Routerx) {
	rg.GET(dto.OpenEndpoint{
		Path:        "",
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
