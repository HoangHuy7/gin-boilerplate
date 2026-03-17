// hoanghuy7 from Vietnamese with love!

package graph

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"go.uber.org/fx"
)

// Module provides all dependencies for GraphQL
var Module = fx.Options(
	fx.Provide(
		NewResolver,
		NewGraphQLHandler,
	),
)

// NewGraphQLHandler creates the GraphQL handler with all dependencies injected
func NewGraphQLHandler(resolver *Resolver) *handler.Server {
	// Create the executable schema with the resolver
	srv := handler.NewDefaultServer(NewExecutableSchema(Config{
		Resolvers: resolver,
	}))

	return srv
}
