package main

import (
	"context"
	_middleware "cozy-inn/app/middleware"
	_route "cozy-inn/app/route"
	_userUseCase "cozy-inn/businesses/users"
	_userController "cozy-inn/controller/users"
	_driverFactory "cozy-inn/driver"

	firestore "cozy-inn/driver/firestore"
	util "cozy-inn/util"

	firebase "firebase.google.com/go"
	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	ctx := context.Background()

	FirestoreConfig := &firebase.Config{
		ProjectID: util.GetFirebaseEnv("project_id"),
	}

	firestore := firestore.InitFirestore(FirestoreConfig)

	userRepository := _driverFactory.NewUserRepository(firestore, ctx)
	userUsecase := _userUseCase.NewUserUsecase(userRepository)
	userController := _userController.NewUserController(userUsecase)

	configLogger := _middleware.ConfigLogger{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	routeController := _route.ControllerList{
		LoggerMiddleware: configLogger.InitLogger(),
		UserController:   userController,
	}

	routeController.InitRoute(e)

	e.Logger.Fatal(e.Start(":3000"))
}
