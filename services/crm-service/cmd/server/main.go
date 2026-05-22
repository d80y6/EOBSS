package main
import (
	"log"
	"net/http"
	"os"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)
func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/health", func(c echo.Context) error { return c.JSON(http.StatusOK, map[string]string{"status": "UP"}) })
	port := os.Getenv("PORT")
	if port == "" { port = "8082" }
	log.Fatal(e.Start(":" + port))
}
