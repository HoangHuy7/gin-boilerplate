package security

import (
	"monorepo/apps/gas/app/config"
	"monorepo/internal/dto"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewSecurity,
		// Convert Config.Casdoor to dto.CasdoorConfig
		func(cfg *config.Config) *dto.CasdoorConfig {
			return &dto.CasdoorConfig{
				Organizations: cfg.Casdoor.Organizations,
			}
		},
	),
)
