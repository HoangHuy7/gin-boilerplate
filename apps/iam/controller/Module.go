package controller

import (
	v1 "monorepo/iam/controller/v1"
	"monorepo/internal/base"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(
			NewRouter,
			fx.ParamTags(`group:"controllers"`),
		),
		fx.Annotate(
			v1.NewHelloController,
			fx.As(new(base.Controller)),
			fx.ResultTags(`group:"controllers"`),
		),
	),
)
