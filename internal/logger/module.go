// hoanghuy7 from Vietnamese with love!

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
		func(gl *GoLogger) *zap.Logger {
			return gl.Zap
		},
	),
)

type GoLogger struct {
	Zap     *zap.Logger
	skipped *zap.Logger // dùng cho wrapper methods

}

// ANSI color codes
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[97m"
)

// levelWithColorEncoder: in level có màu, sau đó reset về trắng
// → fields JSON sau đó sẽ là trắng, không bị nhuộm màu level
func levelWithColorEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	var levelColor string
	switch l {
	case zapcore.DebugLevel:
		levelColor = colorCyan
	case zapcore.InfoLevel:
		levelColor = colorCyan
	case zapcore.WarnLevel:
		levelColor = colorYellow
	case zapcore.ErrorLevel, zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		levelColor = colorRed
	default:
		levelColor = colorWhite
	}
	// Sau level reset về trắng → fields không bị nhuộm màu
	enc.AppendString(levelColor + l.CapitalString() + colorReset + colorWhite)
}

// colorCore: colorize message theo level
type colorCore struct {
	zapcore.Core
}

func (c *colorCore) With(fields []zapcore.Field) zapcore.Core {
	return &colorCore{Core: c.Core.With(fields)}
}

func (c *colorCore) Check(entry zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(entry.Level) {
		return ce.AddCore(entry, c)
	}
	return ce
}

func (c *colorCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	color := colorWhite
	if entry.Level >= zapcore.ErrorLevel {
		color = colorRed
	}
	entry.Message = color + entry.Message + colorReset + colorWhite
	return c.Core.Write(entry, fields)
}

func NewLogger() *GoLogger {
	env := os.Getenv("APP_ENV")

	if env == "prod" {
		cfg := zap.NewProductionConfig()
		cfg.EncoderConfig.EncodeDuration = zapcore.MillisDurationEncoder
		cfg.Encoding = "json"
		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		cfg.OutputPaths = []string{"stdout"}

		logger, _ := cfg.Build()
		return &GoLogger{Zap: logger}
	}

	// ========== DEV MODE ==========
	encCfg := zapcore.EncoderConfig{
		TimeKey:       "T",
		LevelKey:      "L",
		NameKey:       "N",
		CallerKey:     "C",
		MessageKey:    "M",
		StacktraceKey: "S",

		LineEnding: zapcore.DefaultLineEnding,

		// Level có màu, reset về trắng ngay sau → fields không bị nhuộm
		EncodeLevel: levelWithColorEncoder,

		// Time: trắng
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(colorWhite + t.Format("2006-01-02T15:04:05.000Z0700") + colorReset)
		},

		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	consoleEnc := zapcore.NewConsoleEncoder(encCfg)
	sink, _, _ := zap.Open("stdout")

	baseCore := zapcore.NewCore(
		consoleEnc,
		zapcore.AddSync(sink),
		zapcore.DebugLevel,
	)

	core := &colorCore{Core: baseCore}

	logger := zap.New(
		core,
		zap.AddCaller(),
		zap.Development(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	return &GoLogger{Zap: logger, skipped: logger.WithOptions(zap.AddCallerSkip(1))}
}

func ZapLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)

		if query != "" {
			path = path + "?" + query
		}

		logger.Info("GIN_REQUEST",
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("ip", c.ClientIP()),
			zap.Duration("latency", latency),
			zap.String("user-agent", c.Request.UserAgent()),
		)
	}
}

func (l *GoLogger) INFO(msg string, fields ...zap.Field) {
	l.skipped.Info(msg, fields...)
}

func (l *GoLogger) ERROR(msg string, fields ...zap.Field) {
	l.skipped.Error(msg, fields...)
}
