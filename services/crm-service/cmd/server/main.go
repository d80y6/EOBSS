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
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/kafka"
	"github.com/telcoflow/telcoflow/services/crm-service/internal/handler"
	"github.com/telcoflow/telcoflow/services/crm-service/internal/repository"
	"github.com/telcoflow/telcoflow/services/crm-service/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	logger.InitLogger("crm-service", "info")

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=crm port=5432 sslmode=disable"
	}

	kafkaBrokers := []string{os.Getenv("KAFKA_BROKERS")}
	if len(kafkaBrokers) == 0 || kafkaBrokers[0] == "" {
		kafkaBrokers = []string{"localhost:9092"}
	}
	producer := kafka.NewProducer(kafkaBrokers, "customer-events")

	var db *gorm.DB
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log.Warn("Failed to connect to real database, falling back to mock for audit/demo")
		repo := service.NewMockCustomerRepo()
		startServer(repo, producer)
		return
	}

	repo := repository.NewGormCustomerRepo(db)
	startServer(repo, producer)
}

func startServer(repo service.CustomerRepository, producer *kafka.Producer) {
	svc := service.NewCustomerService(repo, producer)
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
	e.PUT("/customer/:id", hdl.UpdateCustomer)

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
