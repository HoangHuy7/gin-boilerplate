// hoanghuy7 from Vietnamese with love!

package graph

import (
	"context"
	"errors"
	"monorepo/shares/exception"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/vektah/gqlparser/v2/gqlerror"
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

	// 👇 đặt ở đây
	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		var appErr *exception.AppError
		if errors.As(e, &appErr) {
			return &gqlerror.Error{
				Message: appErr.Message,
				Extensions: map[string]interface{}{
					"code": appErr.Code,
				},
				Err: appErr,
			}
		}

		return &gqlerror.Error{
			Message: "Internal server error",
		}
	})
	return srv
}
