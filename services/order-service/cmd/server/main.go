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
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/auth"
)

func main() {
	logger.InitLogger("order-service", "info")
	rbac := auth.NewRBACManager()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "UP"})
	})

	// TMF622 Routes
	e.POST("/productOrderManagement/v4/productOrder", func(c echo.Context) error {
		return c.JSON(http.StatusAccepted, map[string]string{"id": "ORD-123", "status": "Acknowledged"})
	}, auth.RBACMiddleware(rbac, auth.PermOrderCreate))

	e.GET("/productOrderManagement/v4/productOrder/:id", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"id": c.Param("id"), "status": "InProgress"})
	}, auth.RBACMiddleware(rbac, auth.PermOrderRead))

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
