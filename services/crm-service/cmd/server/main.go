package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/logger"
	"github.com/telcoflow/telcoflow/services/crm-service/internal/handler"
	"github.com/telcoflow/telcoflow/services/crm-service/internal/service"
)

func main() {
	logger.InitLogger("crm-service", "info")

	// Initialize Repository, Service and Handler
	repo := service.NewMockCustomerRepo()
	svc := service.NewCustomerService(repo)
	hdl := handler.NewCustomerHandler(svc)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "UP"})
	})

	// Customer Routes
	e.POST("/customer", hdl.CreateCustomer)
	e.GET("/customer/:id", hdl.GetCustomer)

	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Service failed to start")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		logger.Error("Server shutdown error")
	}
}
