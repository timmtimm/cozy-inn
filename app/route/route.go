package route

import (
	"cozy-inn/controller/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	LoggerMiddleware echo.MiddlewareFunc
	UserController   *users.UserController
}

func (cl *ControllerList) InitRoute(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	// v1 := e.Group("/api/v1/users")

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Hello World!",
		})
	})

	e.POST("/register", cl.UserController.Register)
}
