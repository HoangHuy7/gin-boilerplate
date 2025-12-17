package logger

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Module = fx.Options(
	fx.Provide(
		NewLogger,
	),
)

type GoLogger struct {
	Zap *zap.Logger
}

func NewLogger() *GoLogger {
	env := os.Getenv("APP_ENV") // dev | prod

	if env == "prod" {
		// ========== PROD MODE ==========
		cfg := zap.NewProductionConfig()
		cfg.EncoderConfig.EncodeDuration = zapcore.MillisDurationEncoder
		cfg.Encoding = "json"
		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		cfg.OutputPaths = []string{"stdout"}

		logger, _ := cfg.Build()
		return &GoLogger{Zap: logger}
	}

	// ========== DEV MODE ==========
	cfg := zap.NewDevelopmentConfig()
	cfg.Encoding = "console"
	cfg.EncoderConfig.EncodeDuration = zapcore.MillisDurationEncoder
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.OutputPaths = []string{"stdout"}

	logger, _ := cfg.Build()
	return &GoLogger{Zap: logger}
}

func ZapLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		latency := time.Since(start)

		logger.Info("request completed",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.String("client_ip", c.ClientIP()),
			zap.Duration("latency", latency),
		)
	}
}

func (this GoLogger) INFO(msg string, fields ...zap.Field) {
	this.Zap.Info(msg, fields...)
}
