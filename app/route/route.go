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

	// multiple role
	allMiddleware := _middleware.RoleMiddleware{Role: []string{"user", "receptionist", "admin"}}
	adminReceptionistMiddleware := _middleware.RoleMiddleware{Role: []string{"receptionist", "admin"}}

	e.GET("/api/v1/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Hello World!",
		})
	})

	user := e.Group("/api/v1/user")
	user.POST("/register", cl.UserController.UserRegister)
	user.POST("/login", cl.UserController.Login)
	user.GET("/profile", cl.UserController.GetUserProfile, allMiddleware.CheckToken)
	user.PUT("/profile", cl.UserController.UpdateUserProfile, allMiddleware.CheckToken)

	room := e.Group("/api/v1/room")
	room.GET("/", cl.RoomController.GetAllRoom)
	room.GET("/:room-type", cl.RoomController.GetRoom, adminReceptionistMiddleware.CheckToken)
	room.PUT("/:room-type", cl.RoomController.UpdateRoom, adminReceptionistMiddleware.CheckToken)
	room.POST("/check-availability", cl.TransactionController.CheckAvailabilityAllRoom)

	transaction := e.Group("/api/v1/transaction")
	transaction.GET("/", cl.TransactionController.GetAllTransaction, userMiddleware.CheckToken)
	transaction.POST("/", cl.TransactionController.CreateTransaction, userMiddleware.CheckToken)
	transaction.GET("/:transaction-id", cl.TransactionController.GetTransaction, userMiddleware.CheckToken)
	transaction.PUT("/:transaction-id", cl.TransactionController.UpdatePayment, userMiddleware.CheckToken)
	transaction.DELETE("/:transaction-id", cl.TransactionController.AdminDelete, adminMiddleware.CheckToken)
	transaction.GET("/verification", cl.TransactionController.GetAllPaymentNotVerified, adminReceptionistMiddleware.CheckToken)
	transaction.GET("/verification/:transaction-id", cl.TransactionController.GetTransactionOnVerification, adminReceptionistMiddleware.CheckToken)
	transaction.PUT("/verification/:transaction-id", cl.TransactionController.UpdateVerification, adminReceptionistMiddleware.CheckToken)
	transaction.GET("/check-in", cl.TransactionController.GetAllReadyCheckIn, adminReceptionistMiddleware.CheckToken)
	transaction.GET("/check-in/:transaction-id", cl.TransactionController.GetCheckIn, adminReceptionistMiddleware.CheckToken)
	transaction.PUT("/check-in/:transaction-id", cl.TransactionController.UpdateCheckIn, adminReceptionistMiddleware.CheckToken)
	transaction.GET("/check-out", cl.TransactionController.GetAllReadyCheckOut, adminReceptionistMiddleware.CheckToken)
	transaction.GET("/check-out/:transaction-id", cl.TransactionController.GetCheckOut, adminReceptionistMiddleware.CheckToken)
	transaction.PUT("/check-out/:transaction-id", cl.TransactionController.UpdateCheckOut, adminReceptionistMiddleware.CheckToken)

	admin := e.Group("/api/v1/admin", adminMiddleware.CheckToken)
	admin.GET("/user", cl.UserController.AdminGetUserList)
	admin.POST("/user", cl.UserController.AdminRegister)
	admin.GET("/user/:user-email", cl.UserController.AdminGetUser)
	admin.PUT("/user/:user-email", cl.UserController.AdminUpdate)
	admin.DELETE("/user/:user-email", cl.UserController.AdminDelete)
	admin.GET("/room", cl.RoomController.GetAllRoom)
	admin.GET("/room/:room-type", cl.RoomController.GetRoom)
	admin.PUT("/room/:room-type", cl.RoomController.UpdateRoom)
	admin.POST("/room", cl.RoomController.CreateRoom)
	admin.DELETE("/room/:room-type", cl.RoomController.DeleteRoom)
	admin.GET("/transaction", cl.TransactionController.AdminGetAllTransaction)
	admin.GET("/transaction/:transaction-id", cl.TransactionController.AdminGetTransaction)
	admin.PUT("/transaction/:transaction-id", cl.TransactionController.AdminUpdateTransaction)
	admin.DELETE("/transaction/:transaction-id", cl.TransactionController.AdminDeleteTransaction)
}
