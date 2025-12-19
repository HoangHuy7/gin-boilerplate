// hoanghuy7 from Vietnamese with love!

package routerx

import (
	"monorepo/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Routerx struct {
	Gin  *gin.RouterGroup
	Meta *dto.Metadata
}

func NewRouterx(g *gin.RouterGroup, meta *dto.Metadata) *Routerx {
	return &Routerx{
		Gin:  g,
		Meta: meta,
	}
}

func (this *Routerx) GET(oe dto.OpenEndpoint) {
	oe.Method = http.MethodGet
	this.Gin.GET(oe.Path, oe.Handler)
	this.addEndpoint(oe)
}

func (this *Routerx) POST(oe dto.OpenEndpoint) {
	oe.Method = http.MethodPost
	this.Gin.POST(oe.Path, oe.Handler)
	this.addEndpoint(oe)
}

func (this *Routerx) DELETE(oe dto.OpenEndpoint) {
	oe.Method = http.MethodDelete
	this.Gin.GET(oe.Path, oe.Handler)
	this.addEndpoint(oe)

}

func (this *Routerx) PUT(oe dto.OpenEndpoint) {
	oe.Method = http.MethodPut
	this.Gin.POST(oe.Path, oe.Handler)
	this.addEndpoint(oe)
}

func (this *Routerx) addEndpoint(oe dto.OpenEndpoint) {
	this.Meta.Endpoints = append(this.Meta.Endpoints, oe)
}
