package server

import (
	"context"
	"monorepo/internal/dto"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func RunServer(lc fx.Lifecycle,
	sr *Router,
	am *dto.AppMetadata,
	router *gin.Engine) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			sr.RegisterAll(router, am)
			go router.Run()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})

}
