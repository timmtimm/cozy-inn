package main

import (
	"context"
	"cozy-inn/util"
	"fmt"

	_middleware "cozy-inn/app/middleware"
	_route "cozy-inn/app/route"
	_driverFactory "cozy-inn/driver"
	firestore "cozy-inn/driver/firestore"

	_userUseCase "cozy-inn/businesses/users"
	_userController "cozy-inn/controller/users"

	_roomUseCase "cozy-inn/businesses/rooms"
	_roomController "cozy-inn/controller/rooms"

	_transactionUseCase "cozy-inn/businesses/transactions"
	_transactionController "cozy-inn/controller/transactions"

	firebase "firebase.google.com/go"
	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	ctx := context.Background()

	configLogger := _middleware.ConfigLogger{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       util.GetConfig("JWT_SECRET_KEY"),
		ExpiresDuration: 1,
	}

	FirestoreConfig := &firebase.Config{
		ProjectID: util.GetFirebaseEnv("project_id"),
	}

	firestore := firestore.InitFirestore(FirestoreConfig)

	userRepository := _driverFactory.NewUserRepository(firestore, ctx)
	userUsecase := _userUseCase.NewUserUsecase(userRepository, &configJWT)
	userController := _userController.NewUserController(userUsecase)

	RoomRepository := _driverFactory.NewRoomRepository(firestore, ctx)
	RoomUsecase := _roomUseCase.NewRoomUsecase(RoomRepository)
	RoomController := _roomController.NewRoomController(RoomUsecase)

	TransactionRepository := _driverFactory.NewTransactionRepository(firestore, ctx)
	TransactionUsecase := _transactionUseCase.NewTransactionUsecase(TransactionRepository)
	TransactionController := _transactionController.NewTransactionController(TransactionUsecase)

	routeController := _route.ControllerList{
		LoggerMiddleware:      configLogger.Init(),
		JWTMiddleware:         configJWT.Init(),
		UserController:        userController,
		RoomController:        RoomController,
		TransactionController: TransactionController,
	}

	routeController.Init(e)

	appPort := fmt.Sprintf(":%s", util.GetConfig("APP_PORT"))

	e.Logger.Fatal(e.Start(appPort))
}
