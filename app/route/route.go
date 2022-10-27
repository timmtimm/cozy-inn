package route

import (
	"context"
	"cozy-inn/businesses/users"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	LoggerMiddleware echo.MiddlewareFunc
	UserRepository   users.Repository
	Firestore        *firestore.Client
}

func (cl *ControllerList) InitRoute(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "Hello World!",
		})
	})

	e.POST("/register", func(c echo.Context) error {
		cl.UserRepository.Register(context.Background(), cl.Firestore)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "Hello World!",
		})
	})
}
