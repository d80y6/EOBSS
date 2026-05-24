package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/logger"
	"github.com/telcoflow/telcoflow/services/assurance-service/internal/correlation"
	"github.com/telcoflow/telcoflow/services/assurance-service/internal/handler"
	"github.com/telcoflow/telcoflow/services/assurance-service/internal/service"
)

func main() {
	logger.InitLogger("assurance-service", "info")

	engine := &correlation.AlarmCorrelationEngine{}
	svc := service.NewAssuranceService(engine)
	hdl := handler.NewAlarmHandler(svc)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "UP"})
	})

	// TMF642 Routes
	e.POST("/alarmManagement/v4/alarm", hdl.HandleAlarm)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
