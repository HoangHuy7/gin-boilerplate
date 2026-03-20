package service

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewCustomerService,
		NewProductService,
		NewRedisService,
		NewMenuService,
		NewOrderService,
		NewInventoryService,
		NewDebtService,
		NewDeliveryService,
		NewLogService,
	),
)
