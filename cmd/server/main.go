package main

import (
	"echo-boilerplate/internal/config"
	"echo-boilerplate/internal/handler"
	"echo-boilerplate/internal/repository"
	"echo-boilerplate/internal/server"
	"echo-boilerplate/internal/service"
	"echo-boilerplate/pkg/http"
	"echo-boilerplate/pkg/log"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	conf := config.NewConfig()
	logger := log.NewLog(conf.Env)
	addr := fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port)
	logger.Info("Server is running on " + addr)
	http.Run(logger, newApp(conf, logger), addr)
}

// newApp creates a new echo app
// and injects the dependencies
func newApp(appConfig *config.AppConfig, logger *log.Logger) *echo.Echo {
	h := handler.NewHandler(logger)
	s := service.NewService(logger)
	r := repository.NewRepository(logger)
	userRepository := repository.NewUserRepository(r)
	userService := service.NewUserService(s, userRepository)
	userHandler := handler.NewUserHandler(h, userService)

	return server.NewServerHTTP(logger, userHandler)
}
