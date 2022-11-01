package route

import (
	_middleware "cozy-inn/app/middleware"
	"cozy-inn/controller/rooms"
	"cozy-inn/controller/users"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	LoggerMiddleware echo.MiddlewareFunc
	JWTMiddleware    middleware.JWTConfig
	UserController   *users.UserController
	RoomController   *rooms.RoomController
}

func (cl *ControllerList) Init(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	// single role
	// userMiddleware := _middleware.RoleMiddleware{Role: []string{"user"}}
	adminMiddleware := _middleware.RoleMiddleware{Role: []string{"admin"}}
	// receptionistMiddleware := _middleware.RoleMiddleware{Role: []string{"receptionist"}}

	// multiple role
	userReceptionistMiddleware := _middleware.RoleMiddleware{Role: []string{"user", "receptionist"}}

	e.GET("/api/v1/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Hello World!",
		})
	})

	user := e.Group("/api/v1/user")
	user.POST("/register", cl.UserController.UserRegister)
	user.POST("/login", cl.UserController.Login)
	user.GET("/profile", cl.UserController.GetUserProfile, userReceptionistMiddleware.CheckToken)
	user.PUT("/profile", cl.UserController.UpdateUserProfile, userReceptionistMiddleware.CheckToken)

	room := e.Group("/api/v1/room")
	room.GET("/", cl.RoomController.GetAllRoom)
	room.POST("/create", cl.RoomController.CreateRoom, adminMiddleware.CheckToken)

	// receptionist := e.Group("/api/v1/receptionist", receptionistMiddleware.CheckToken)

	admin := e.Group("/api/v1/admin", adminMiddleware.CheckToken)
	admin.GET("/user-list", cl.UserController.GetUserList)
	admin.POST("/register", cl.UserController.SudoRegister)
	admin.GET("/profile/:user-email", cl.UserController.SudoGetUserProfile)
	admin.PUT("/profile/:user-email", cl.UserController.AdminUpdateUser)
}
