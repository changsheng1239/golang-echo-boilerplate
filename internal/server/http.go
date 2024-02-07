package server

import (
	"context"
	"echo-boilerplate/internal/handler"
	"echo-boilerplate/pkg/log"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewServerHTTP(logger *log.Logger, userHandler handler.UserHandler) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.Debug = false
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogRemoteIP:  true,
		LogHost:      true,
		LogUserAgent: true,
		LogLatency:   true,
		LogStatus:    true,
		LogRequestID: true,
		LogURI:       true,
		LogError:     true,
		HandleError:  true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			l := HTTPLogger(logger, c, v)
			if v.Error == nil {
				l.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST")
			} else {
				l.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR", slog.String("err", v.Error.Error()))
			}
			return nil
		},
	}))

	e.GET("/user/:id", userHandler.GetUserByID)

	return e
}

func HTTPLogger(logger *log.Logger, c echo.Context, v middleware.RequestLoggerValues) *log.Logger {
	slogger := logger.With(
		"request_id", c.Response().Header().Get(echo.HeaderXRequestID),
		"uri", v.URI,
		"status", v.Status,
	)
	return &log.Logger{Logger: slogger}
}
