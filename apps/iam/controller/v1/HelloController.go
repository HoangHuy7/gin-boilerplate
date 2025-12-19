// hoanghuy7 from Vietnamese with love!

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

	rg.POST(dto.OpenEndpoint{
		Path:        "/json",
		Handler:     this.JSON,
		Summary:     "Hello World Summary",
		Description: "Hello World Description",
		Request:     &dto.CreatePostRequest{},
		Responses: map[int]any{
			200: map[string]any{
				"message": "string",
				"data":    dto.CreatePostRequest{},
			},
			400: gin.H{"error": "string"},
		},
	})

}

func (this *HelloController) Hello(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func (this *HelloController) JSON(c *gin.Context) {
	var req dto.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Hello World",
		"data":    req,
	})
}
