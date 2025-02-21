package middlewares

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func RegisterMiddlewares(e *echo.Echo) {
	e.Use(logger())
	e.Use(middleware.Recover())
}

func logger() echo.MiddlewareFunc {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)

	logger, _ := config.Build()

	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:      true,
		LogStatus:   true,
		LogLatency:  true,
		LogProtocol: true,
		LogRemoteIP: true,
		LogMethod:   true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("IP", v.RemoteIP),
				zap.String("Protocol", v.Protocol),
				zap.String("URI", v.URI),
				zap.String("method", v.Method),
				zap.Int("status", v.Status),
				zap.Duration("latency", v.Latency),
			)
			return nil
		},
	})
}
