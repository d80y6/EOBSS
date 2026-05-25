package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/logger"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/kafka"
	"github.com/telcoflow/telcoflow/services/mediation-service/internal/service"
	"github.com/telcoflow/telcoflow/services/mediation-service/internal/collector"
	"context"
)

func main() {
	logger.InitLogger("mediation-service", "info")

	kafkaBrokers := []string{os.Getenv("KAFKA_BROKERS")}
	if len(kafkaBrokers) == 0 || kafkaBrokers[0] == "" {
		kafkaBrokers = []string{"localhost:9092"}
	}
	producer := kafka.NewProducer(kafkaBrokers, "usage-events")
	transformer := service.NewUsageTransformer(producer)

	// Start IPFIX Collector in background
	dataChan := make(chan map[string]interface{}, 100)
	coll := &collector.IPFIXCollector{Port: 2055}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go coll.Start(ctx, dataChan)

	// Start Transformer consumer loop
	go func() {
		for data := range dataChan {
			_, _ = transformer.Transform(ctx, data)
		}
	}()

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
