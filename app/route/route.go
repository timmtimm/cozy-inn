package route

import (
	_middleware "cozy-inn/app/middleware"
	"cozy-inn/controller/users"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	LoggerMiddleware echo.MiddlewareFunc
	JWTMiddleware    middleware.JWTConfig
	UserController   *users.UserController
}

func (cl *ControllerList) InitRoute(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	e.GET("/api/v1/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Hello World!",
		})
	})

	user := e.Group("/api/v1/user")

	user.POST("/register", cl.UserController.Register)
	user.POST("/login", cl.UserController.Login)

	userMiddleware := _middleware.RoleMiddleware{
		Role: []string{"user"},
	}

	room := e.Group("/api/v1/room", userMiddleware.CheckToken, middleware.JWTWithConfig(cl.JWTMiddleware))

	room.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Hello Room!",
		})
	})
}
