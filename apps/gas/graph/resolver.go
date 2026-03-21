// hoanghuy7 from Vietnamese with love!

package graph

import (
	"monorepo/apps/gas/service"
)

// Resolver is the root resolver that holds all service dependencies
type Resolver struct {
	CustomerService  *service.CustomerService
	MenuService      *service.MenuService
	ProdService      *service.ProductService
	OrderService     *service.OrderService
	InventoryService *service.InventoryService
	DebtService      *service.DebtService
	DeliveryService  *service.DeliveryService
	AnalysisService  *service.AnalysisService
}

// NewResolver creates a new Resolver with all dependencies injected via Uber FX
func NewResolver(
	customerService *service.CustomerService,
	menuService *service.MenuService,
	prodService *service.ProductService,
	orderService *service.OrderService,
	inventoryService *service.InventoryService,
	debtService *service.DebtService,
	deliveryService *service.DeliveryService,
	analysisService *service.AnalysisService,
) *Resolver {
	return &Resolver{
		CustomerService:  customerService,
		MenuService:      menuService,
		ProdService:      prodService,
		OrderService:     orderService,
		InventoryService: inventoryService,
		DebtService:      debtService,
		DeliveryService:  deliveryService,
		AnalysisService:  analysisService,
	}
}
