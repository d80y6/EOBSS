package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/logger"
	"github.com/telcoflow/telcoflow/services/inventory-service/internal/handler"
	"github.com/telcoflow/telcoflow/services/inventory-service/internal/integration/netbox"
	"github.com/telcoflow/telcoflow/services/inventory-service/internal/service"
)

func main() {
	logger.InitLogger("inventory-service", "info")

	netboxClient := netbox.NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"))
	svc := service.NewInventoryService(netboxClient)
	hdl := handler.NewInventoryHandler(svc)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "UP"})
	})

	// TMF639 Routes
	e.POST("/resourceInventoryManagement/v4/resource", hdl.CreateResource)
	e.GET("/resourceInventoryManagement/v4/resource/:id", hdl.GetResource)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
