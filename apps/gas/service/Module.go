package service

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewProductService,
		NewRedisService,
		NewMenuService,
	),
)
