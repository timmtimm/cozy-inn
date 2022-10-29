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

	userMiddleware := _middleware.RoleMiddleware{Role: []string{"user"}}
	adminMiddleware := _middleware.RoleMiddleware{Role: []string{"admin"}}
	resepsionistMiddleware := _middleware.RoleMiddleware{Role: []string{"resepsionist"}}

	e.GET("/api/v1/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Hello World!",
		})
	})

	user := e.Group("/api/v1/user")
	user.POST("/register", cl.UserController.UserRegister)
	user.POST("/login", cl.UserController.Login)
	user.GET("/profile", cl.UserController.GetUserProfile, userMiddleware.CheckToken)
	user.POST("/profile", cl.UserController.UpdateUserProfile, userMiddleware.CheckToken)

	room := e.Group("/api/v1/room")
	room.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Hello Room!",
		})
	})

	resepsionist := e.Group("/api/v1/resepsionist", resepsionistMiddleware.CheckToken)
	resepsionist.GET("/profile", cl.UserController.GetUserProfile)
	resepsionist.POST("/profile", cl.UserController.UpdateUserProfile)

	admin := e.Group("/api/v1/admin", adminMiddleware.CheckToken)
	admin.POST("/register", cl.UserController.SudoRegister)
	admin.POST("/profile", cl.UserController.SudoGetUserProfile)
}
