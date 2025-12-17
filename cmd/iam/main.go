package main

import (
	"monorepo/iam/controller"
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
	//config.LoadConfig()
	fx.New(
		fx.WithLogger(func(gl *logger.GoLogger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: gl.Zap}
		}),
		fx.Provide(
			NewGinEngine,
		),
		controller.Module,
		//middleware.Module,
		//database.Module,
		//http.Module,
		logger.Module,
		//casbinConfig.Module,
		fx.Invoke(server.RunServer),
	).Run()
}
