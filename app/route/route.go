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
	allMiddleware := _middleware.RoleMiddleware{Role: []string{"user", "receptionist", "admin"}}
	adminReceptionistMiddleware := _middleware.RoleMiddleware{Role: []string{"receptionist", "admin"}}

	apiV1 := e.Group("/api/v1")

	apiV1.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Hello World!",
		})
	})

	user := apiV1.Group("/user")
	user.POST("/register", cl.UserController.UserRegister)
	user.POST("/login", cl.UserController.Login)
	user.GET("/profile", cl.UserController.GetUserProfile, allMiddleware.CheckToken)
	user.PUT("/profile", cl.UserController.UpdateUserProfile, allMiddleware.CheckToken)

	room := apiV1.Group("/room")
	room.GET("/", cl.RoomController.GetAllRoom)
	room.GET("/:room-type", cl.RoomController.GetRoom, adminReceptionistMiddleware.CheckToken)
	room.PUT("/:room-type", cl.RoomController.UpdateRoom, adminReceptionistMiddleware.CheckToken)
	room.POST("/check-availability", cl.TransactionController.CheckAvailabilityAllRoom)

	transaction := apiV1.Group("/transaction")
	transaction.GET("", cl.TransactionController.GetAllTransaction, userMiddleware.CheckToken)
	transaction.POST("/", cl.TransactionController.CreateTransaction, userMiddleware.CheckToken)
	transaction.GET("/:transaction-id", cl.TransactionController.GetTransaction, userMiddleware.CheckToken)
	transaction.PUT("/:transaction-id", cl.TransactionController.UpdatePayment, userMiddleware.CheckToken)
	transaction.PUT("/cancel/:transaction-id", cl.TransactionController.CancelTransaction, userMiddleware.CheckToken)
	transaction.DELETE("/:transaction-id", cl.TransactionController.AdminDelete, adminMiddleware.CheckToken)

	transactionVerification := transaction.Group("/verification")
	transactionVerification.GET("", cl.TransactionController.GetAllPaymentNotVerified, adminReceptionistMiddleware.CheckToken)
	transactionVerification.GET("/:transaction-id", cl.TransactionController.GetTransactionOnVerification, adminReceptionistMiddleware.CheckToken)
	transactionVerification.PUT("/:transaction-id", cl.TransactionController.UpdateVerification, adminReceptionistMiddleware.CheckToken)

	transactionCheckIn := transaction.Group("/check-in")
	transactionCheckIn.GET("", cl.TransactionController.GetAllReadyCheckIn, adminReceptionistMiddleware.CheckToken)
	transactionCheckIn.GET("/:transaction-id", cl.TransactionController.GetCheckIn, adminReceptionistMiddleware.CheckToken)
	transactionCheckIn.PUT("/:transaction-id", cl.TransactionController.UpdateCheckIn, adminReceptionistMiddleware.CheckToken)

	transactionCheckOut := transaction.Group("/check-out")
	transactionCheckOut.GET("", cl.TransactionController.GetAllReadyCheckOut, adminReceptionistMiddleware.CheckToken)
	transactionCheckOut.GET("/:transaction-id", cl.TransactionController.GetCheckOut, adminReceptionistMiddleware.CheckToken)
	transactionCheckOut.PUT("/:transaction-id", cl.TransactionController.UpdateCheckOut, adminReceptionistMiddleware.CheckToken)

	receptionist := apiV1.Group("/receptionist")
	receptionist.POST("/transaction", cl.TransactionController.ReceptionistCreateTransaction, receptionistMiddleware.CheckToken)

	admin := apiV1.Group("/admin", adminMiddleware.CheckToken)

	adminUser := admin.Group("/user")
	adminUser.GET("", cl.UserController.AdminGetUserList)
	adminUser.POST("/", cl.UserController.AdminRegister)
	adminUser.GET("/:user-email", cl.UserController.AdminGetUser)
	adminUser.PUT("/:user-email", cl.UserController.AdminUpdate)
	adminUser.DELETE("/:user-email", cl.UserController.AdminDelete)

	adminRoom := admin.Group("/room")
	adminRoom.GET("", cl.RoomController.GetAllRoom)
	adminRoom.GET("/:room-type", cl.RoomController.GetRoom)
	adminRoom.PUT("/:room-type", cl.RoomController.UpdateRoom)
	adminRoom.POST("", cl.RoomController.CreateRoom)
	adminRoom.DELETE("/:room-type", cl.RoomController.DeleteRoom)

	adminTransaction := admin.Group("/transaction")
	adminTransaction.GET("", cl.TransactionController.AdminGetAllTransaction)
	adminTransaction.GET("/:transaction-id", cl.TransactionController.AdminGetTransaction)
	adminTransaction.PUT("/:transaction-id", cl.TransactionController.AdminUpdateTransaction)
	adminTransaction.DELETE("/:transaction-id", cl.TransactionController.AdminDeleteTransaction)
}
