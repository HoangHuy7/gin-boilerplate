// hoanghuy7 from Vietnamese with love!

package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewGinEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	return r
}

func runServer(lc fx.Lifecycle, router *gin.Engine) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			router.GET("/ping", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "pong",
				})
			})
			go router.Run()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})

}

func main() {
	//config.LoadConfig()
	fx.New(
		fx.Provide(
			NewGinEngine,
		),
		//middleware.Module,
		//database.Module,
		//http.Module,
		//casbinConfig.Module,
		fx.Invoke(runServer),
	).Run()
}
