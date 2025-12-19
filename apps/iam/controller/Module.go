// hoanghuy7 from Vietnamese with love!

package controller

import (
	"monorepo/apps/iam/controller/v1"
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
			v1.NewHelloController,
			fx.As(new(base.Controller)),
			fx.ResultTags(`group:"controllers"`),
		),
	),
)
