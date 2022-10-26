package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	LoggerMiddleware echo.MiddlewareFunc
}

func (cl *ControllerList) InitRoute(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "Hello World!",
		})
	})
}
