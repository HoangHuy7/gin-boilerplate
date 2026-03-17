// hoanghuy7 from Vietnamese with love!

package graphql

import (
	"monorepo/internal/base/routerx"
	"monorepo/internal/dto"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

// GraphQLController handles GraphQL endpoints
type GraphQLController struct {
	handler       *handler.Server
	enablePlayground bool
}

// NewGraphQLController creates a new GraphQL controller
func NewGraphQLController(
	handler *handler.Server,
) *GraphQLController {
	return &GraphQLController{
		handler:       handler,
		enablePlayground: true,
	}
}

// Register implements base.Controller
func (c *GraphQLController) Register(r *routerx.Routerx) {
	// GraphQL endpoint - use empty path since the base path is already /api/graphql
	r.GET(dto.OpenEndpoint{
		Path:    "",
		Handler: c.graphqlHandler(),
	})
	r.POST(dto.OpenEndpoint{
		Path:    "",
		Handler: c.graphqlHandler(),
	})
	
	// GraphQL Playground endpoint
	if c.enablePlayground {
		r.GET(dto.OpenEndpoint{
			Path:    "/playground",
			Handler: c.playgroundHandler(),
		})
	}
}

// GetMetadata implements base.Controller
func (c *GraphQLController) GetMetadata() *dto.Metadata {
	return &dto.Metadata{
		Path:          "/graphql",
		Version:       "",
		Tag:           "GraphQL",
		Endpoints:     []dto.OpenEndpoint{},
		EnableOpenAPI: false,
		IsNotAuth:     true, // Set to false if you want authentication
	}
}

// graphqlHandler returns a Gin handler for GraphQL
func (c *GraphQLController) graphqlHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c.handler.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// playgroundHandler returns a Gin handler for GraphQL Playground
func (c *GraphQLController) playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL Playground", "/api/graphql")
	return func(ctx *gin.Context) {
		h.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
