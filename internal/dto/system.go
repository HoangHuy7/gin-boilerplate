package dto

import (
	"github.com/gin-gonic/gin"
)

type OpenEndpoint struct {
	Method      string
	Path        string
	Handler     gin.HandlerFunc
	Request     any // body
	Query       any // body
	Responses   map[int]any
	Summary     string
	Description string
}
type OpenGroup struct {
	Gin       *gin.RouterGroup
	BasePath  string
	Endpoints []OpenEndpoint
}
type Metadata struct {
	Path          string
	Version       string
	Tag           string
	Endpoints     []OpenEndpoint
	EnableOpenAPI bool
}
