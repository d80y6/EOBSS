package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/logger"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/kafka"
	"github.com/telcoflow/telcoflow/services/mediation-service/internal/service"
)

func main() {
	logger.InitLogger("mediation-service", "info")

	kafkaBrokers := []string{os.Getenv("KAFKA_BROKERS")}
	if len(kafkaBrokers) == 0 || kafkaBrokers[0] == "" {
		kafkaBrokers = []string{"localhost:9092"}
	}
	producer := kafka.NewProducer(kafkaBrokers, "usage-events")
	_ = service.NewUsageTransformer(producer)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "UP"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
