package service

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewCustomerService,
		NewProductService,
		NewRedisService,
		NewOrderService,
		NewInventoryService,
		NewDebtService,
		NewDeliveryService,
		NewLogService,
		NewAnalysisService,
	),
)
