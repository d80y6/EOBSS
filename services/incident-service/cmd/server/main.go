package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/logger"
	"github.com/telcoflow/telcoflow/services/incident-service/internal/handler"
	"github.com/telcoflow/telcoflow/services/incident-service/internal/service"
)

func main() {
	logger.InitLogger("incident-service", "info")

	svc := &service.IncidentService{}
	hdl := handler.NewIncidentHandler(svc)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "UP"})
	})

	// TMF621 Routes
	e.POST("/troubleTicketManagement/v4/troubleTicket", hdl.CreateTicket)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
