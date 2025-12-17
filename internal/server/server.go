package server

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func RunServer(lc fx.Lifecycle,
	sr *Router,
	router *gin.Engine) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			sr.RegisterAll(router)
			go router.Run()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})

}
