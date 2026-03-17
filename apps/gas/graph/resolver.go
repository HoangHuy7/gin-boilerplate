// hoanghuy7 from Vietnamese with love!

package graph

import (
	"monorepo/apps/gas/service"
)

// Resolver is the root resolver that holds all service dependencies
// This resolver is injected with Uber FX
type Resolver struct {
	CustomerService *service.CustomerService
	MenuService     *service.MenuService
}

// NewResolver creates a new Resolver with all dependencies injected via Uber FX
func NewResolver(
	customerService *service.CustomerService,
	menuService *service.MenuService,
) *Resolver {
	return &Resolver{
		CustomerService: customerService,
		MenuService:     menuService,
	}
}
