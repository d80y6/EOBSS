package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/logger"
	"github.com/telcoflow/telcoflow/services/iam-service/internal/handler"
	"github.com/telcoflow/telcoflow/services/iam-service/internal/service"
)

func main() {
	logger.InitLogger("iam-service", "info")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Initialize Services
	// iamSvc := service.NewKeycloakService(...)

	// Routes
	e.GET("/health", func(c echo.Context) error { return c.NoContent(200) })

	// Start server
	go func() {
		if err := e.Start(":8080"); err != nil {
			logger.Fatal("IAM Service failed to start")
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		logger.Error("Server shutdown error")
	}
}
