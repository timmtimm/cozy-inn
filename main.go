package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "Hello World!",
		})
	})
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":3000"))
}
