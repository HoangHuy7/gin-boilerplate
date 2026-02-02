// hoanghuy7 from Vietnamese with love!

package app

import (
	"monorepo/apps/gas/app/config"
	"monorepo/apps/gas/app/database"
	"monorepo/apps/gas/app/redis"
	"monorepo/internal/dto"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(config.NewConfig),
	fx.Provide(func(cf *config.Config) *dto.OIDC {
		return &dto.OIDC{
			Realm:        cf.Oidc.Realm,
			Issuer:       cf.Oidc.Issuer,
			ClientID:     cf.Oidc.ClientID,
			ClientSecret: cf.Oidc.ClientSecret,
		}
	}),
	fx.Provide(
		config.NewAppMetadata,
		database.NewDataSources,
		redis.NewRedisClient,
	),
)
