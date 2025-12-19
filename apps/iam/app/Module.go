// hoanghuy7 from Vietnamese with love!

package app

import (
	"monorepo/apps/iam/app/casbin"
	"monorepo/apps/iam/app/config"
	"monorepo/apps/iam/app/database"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Invoke(config.NewConfig),
	fx.Provide(
		config.NewAppMetadata,
		database.NewDataSources,
		casbin.NewCasbin,
	),
)
