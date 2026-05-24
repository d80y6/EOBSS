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
	"github.com/telcoflow/telcoflow/services/billing-service/internal/handler"
	"github.com/telcoflow/telcoflow/services/billing-service/internal/rating"
	"github.com/telcoflow/telcoflow/services/billing-service/internal/repository"
	"github.com/telcoflow/telcoflow/services/billing-service/internal/service"
	"database/sql"
)

func main() {
	logger.InitLogger("billing-service", "info")

	// Initialize ClickHouse repository
	chURL := os.Getenv("CLICKHOUSE_URL")
	var repo *repository.ClickHouseCDRRepository
	if chURL != "" {
		db, err := sql.Open("clickhouse", chURL)
		if err == nil {
			repo = repository.NewClickHouseCDRRepository(db)
		}
	}
	_ = repo // Prepared for future persistence depth

	ratingEngine := rating.NewRatingEngine()
	svc := service.NewInvoicingService(ratingEngine)
	hdl := handler.NewBillingHandler(svc)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "UP"})
	})

	// Billing Routes
	e.POST("/billing/invoice/:customerId", hdl.GenerateInvoice)
	e.POST("/billing/usage", hdl.ProcessUsage)
	e.POST("/billing/activate", hdl.ActivateBilling)

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
