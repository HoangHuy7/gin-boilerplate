// hoanghuy7 from Vietnamese with love!

package server

import (
	"context"
	"monorepo/internal/base/security"
	"monorepo/internal/dto"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func RunServer(lc fx.Lifecycle,
	sr *Router,
	am *dto.AppMetadata,
	se *security.Security,
	router *gin.Engine) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			sr.RegisterAll(router, am, se)
			go router.Run(":8082")

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})

}
