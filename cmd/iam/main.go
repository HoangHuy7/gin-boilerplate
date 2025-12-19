// hoanghuy7 from Vietnamese with love!

package main

import (
	"monorepo/apps/iam/app"
	"monorepo/apps/iam/controller"
	"monorepo/internal/logger"
	"monorepo/internal/server"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func NewGinEngine(gl *logger.GoLogger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(logger.ZapLogger(gl.Zap))
	return r
}

func main() {
	fx.New(
		app.Module,
		controller.Module,
		logger.Module,
		fx.WithLogger(func(gl *logger.GoLogger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: gl.Zap}
		}),
		fx.Provide(NewGinEngine),
		fx.Invoke(server.RunServer),
	).Run()
}
