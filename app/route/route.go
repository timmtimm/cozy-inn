package route

import (
	_middleware "cozy-inn/app/middleware"
	"cozy-inn/controller/rooms"
	"cozy-inn/controller/transactions"
	"cozy-inn/controller/users"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	LoggerMiddleware      echo.MiddlewareFunc
	JWTMiddleware         middleware.JWTConfig
	UserController        *users.UserController
	RoomController        *rooms.RoomController
	TransactionController *transactions.TransactionController
}

func (cl *ControllerList) Init(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	// single role
	userMiddleware := _middleware.RoleMiddleware{Role: []string{"user"}}
	adminMiddleware := _middleware.RoleMiddleware{Role: []string{"admin"}}
	receptionistMiddleware := _middleware.RoleMiddleware{Role: []string{"receptionist"}}

	// multiple role
	AllMiddleware := _middleware.RoleMiddleware{Role: []string{"user", "receptionist", "admin"}}

	e.GET("/api/v1/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Hello World!",
		})
	})

	user := e.Group("/api/v1/user")
	user.POST("/register", cl.UserController.UserRegister)
	user.POST("/login", cl.UserController.Login)
	user.GET("/profile", cl.UserController.GetUserProfile, AllMiddleware.CheckToken)
	user.PUT("/profile", cl.UserController.UpdateUserProfile, AllMiddleware.CheckToken)

	room := e.Group("/api/v1/room")
	room.GET("/", cl.RoomController.GetAllRoom)
	room.POST("/", cl.RoomController.CreateRoom, adminMiddleware.CheckToken)
	room.PUT("/:room-type", cl.RoomController.UpdateRoom, adminMiddleware.CheckToken)
	room.DELETE("/:room-type", cl.RoomController.DeleteRoom, adminMiddleware.CheckToken)

	transaction := e.Group("/api/v1/transaction")
	transaction.GET("/", cl.TransactionController.GetAllTransaction, userMiddleware.CheckToken)
	transaction.POST("/", cl.TransactionController.CreateTransaction, userMiddleware.CheckToken)
	transaction.PUT("/:transaction-id", cl.TransactionController.UpdatePayment, userMiddleware.CheckToken)
	transaction.GET("/verification", cl.TransactionController.GetPaymentNotVerified, receptionistMiddleware.CheckToken)

	admin := e.Group("/api/v1/admin", adminMiddleware.CheckToken)
	admin.GET("/user-list", cl.UserController.GetUserList)
	admin.POST("/register", cl.UserController.AdminRegister)
	admin.GET("/profile/:user-email", cl.UserController.AdminGetUserProfile)
	admin.PUT("/profile/:user-email", cl.UserController.AdminUpdateUser)
	admin.DELETE("/profile/:user-email", cl.UserController.AdminDeleteUser)
}
