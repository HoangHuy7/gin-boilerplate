// hoanghuy7 from Vietnamese with love!

package main

import (
	"monorepo/apps/gas/app"
	"monorepo/apps/gas/controller"
	"monorepo/apps/gas/domain"
	"monorepo/apps/gas/graph"
	"monorepo/apps/gas/service"
	"monorepo/internal/base/security"
	"monorepo/internal/logger"
	"monorepo/internal/server"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func NewGinEngine(gl *logger.GoLogger,
	s *security.Security,
	// s3app *s3app.S3Client,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	//println(s3app)
	// r := gin.Default()
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(logger.ZapLogger(gl.Zap))
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
	}))
	r.Use(logger.ZapLogger(gl.Zap))
	return r
}

func main() {
	fx.New(
		app.Module,
		controller.Module,
		service.Module,
		logger.Module,
		domain.Module,
		graph.Module,
		security.Module,
		// fx.WithLogger(func(gl *logger.GoLogger) fxevent.Logger {
		// return &fxevent.ZapLogger{Logger: gl.Zap}
		// }),
		fx.WithLogger(func(gl *logger.GoLogger) fxevent.Logger {
			// clone logger và nâng level lên WARN
			fxLogger := gl.Zap.WithOptions(
				zap.IncreaseLevel(zap.WarnLevel),
			)
			return &fxevent.ZapLogger{Logger: fxLogger}
		}),

		fx.Provide(NewGinEngine),
		fx.Invoke(server.RunServer),
	).Run()
}
