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
	"github.com/telcoflow/telcoflow/services/provisioning-service/internal/adapter"
	"github.com/telcoflow/telcoflow/services/provisioning-service/internal/handler"
	"github.com/telcoflow/telcoflow/services/provisioning-service/internal/service"
)

func main() {
	logger.InitLogger("provisioning-service", "info")

	radiusSecret := os.Getenv("RADIUS_SECRET")
	if radiusSecret == "" {
		radiusSecret = "default-secret-for-dev-only" // Still better to allow override
	}

	ra := adapter.NewRadiusAdapter("radius:1812", radiusSecret)
	ka := &adapter.KamailioAdapter{DBConn: os.Getenv("KAMAILIO_DB_URL")}
	oa := &adapter.Open5GSAdapter{BaseURL: "http://open5gs:8080"}

	svc := service.NewProvisioningService(ra, ka, oa)
	hdl := handler.NewProvisioningHandler(svc)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "UP"})
	})

	e.POST("/provision", hdl.Provision)

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
