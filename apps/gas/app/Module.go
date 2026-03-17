// hoanghuy7 from Vietnamese with love!

package app

import (
	"go.uber.org/fx"
	"monorepo/apps/gas/app/config"
	"monorepo/apps/gas/app/database"
	"monorepo/apps/gas/app/redis"
)

var Module = fx.Options(
	fx.Provide(config.NewConfig),
	fx.Provide(
		config.NewAppMetadata,
		database.NewDataSources,
		redis.NewRedisClient,
	),
)
