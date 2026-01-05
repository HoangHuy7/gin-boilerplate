// hoanghuy7 from Vietnamese with love!

package main

import (
	"monorepo/apps/gas"
	"monorepo/apps/gas/app"
	"monorepo/internal/base/security"
	"monorepo/internal/logger"
	"monorepo/internal/server"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func NewGinEngine(gl *logger.GoLogger, s *security.Security) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(logger.ZapLogger(gl.Zap))
	return r
}

func main() {
	fx.New(
		app.Module,
		gas.Module,
		logger.Module,
		security.Module,
		fx.WithLogger(func(gl *logger.GoLogger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: gl.Zap}
		}),
		fx.Provide(NewGinEngine),
		fx.Invoke(server.RunServer),
	).Run()
}
