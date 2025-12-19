// hoanghuy7 from Vietnamese with love!

package dto

import (
	"time"

	"github.com/gin-gonic/gin"
)

type OpenAPIInfo struct {
	Title        string
	Description  string
	Version      string
	ContactName  string
	ContactURL   string
	ContactEmail string
	LicenseName  string
	LicenseURL   string
}
type AppMetadata struct {
	AppName     string
	Instance    string
	Port        int
	ContextPath string
	OpenAPIInfo OpenAPIInfo
}
type OpenEndpoint struct {
	Method      string
	Path        string
	Handler     gin.HandlerFunc
	Request     any
	Query       any
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

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

type CreatePostRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}
