package main

import (
	"monorepo/apps/iam/app"
	"monorepo/apps/iam/controller"
	"monorepo/internal/logger"
	"monorepo/internal/server"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func NewGinEngine(gl *logger.GoLogger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(logger.ZapLogger(gl.Zap))
	return r
}

var (
	APP_NAME = "iam"
	INSTANCE = uuid.New().Version().String()
)

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
