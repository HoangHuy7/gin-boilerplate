// hoanghuy7 from Vietnamese with love!

package app

import (
	"monorepo/apps/gas/app/config"
	"monorepo/apps/gas/app/database"
	"monorepo/apps/gas/app/redis"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(config.NewConfig),
	fx.Provide(
		config.NewAppMetadata,
		database.NewDataSources,
		redis.NewRedisClient,
		//s3app.NewS3Client,
	),
)
