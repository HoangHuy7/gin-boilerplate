package controller

import (
	"monorepo/apps/gas/controller/customer"
	"monorepo/apps/gas/controller/menu"
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
			customer.NewCustomerController,
			fx.As(new(base.Controller)),
			fx.ResultTags(`group:"controllers"`),
		),
		fx.Annotate(
			menu.NewMenuControllerV1,
			fx.As(new(base.Controller)),
			fx.ResultTags(`group:"controllers"`),
		),
	),
)
