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
	"github.com/telcoflow/telcoflow/services/catalog-service/internal/handler"
	"github.com/telcoflow/telcoflow/services/catalog-service/internal/repository"
	"github.com/telcoflow/telcoflow/services/catalog-service/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	logger.InitLogger("catalog-service", "info")

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=catalog port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("Failed to connect to database")
	}

	repo := repository.NewGormCatalogRepo(db)
	svc := service.NewCatalogService(repo)
	hdl := handler.NewCatalogHandler(svc)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "UP"})
	})

	// TMF620 Routes
	e.POST("/catalogManagement/v4/productOffering", hdl.CreateOffering)
	e.GET("/catalogManagement/v4/productOffering/:id", hdl.GetOffering)
	e.GET("/catalogManagement/v4/productOffering", hdl.ListOfferings)

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
