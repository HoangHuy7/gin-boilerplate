package controller

import (
	"monorepo/apps/gas/controller/graphql"
	"monorepo/internal/base"
	"monorepo/internal/server"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(
			server.NewRouter,
			fx.ParamTags(`group:"controllers"`),
		),
		fx.Annotate(
			NewHelloController,
			fx.As(new(base.Controller)),
			fx.ResultTags(`group:"controllers"`),
		),
		fx.Annotate(
			graphql.NewGraphQLController,
			fx.As(new(base.Controller)),
			fx.ResultTags(`group:"controllers"`),
		),
	),
)
