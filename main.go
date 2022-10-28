package main

import (
	"context"
	_middleware "cozy-inn/app/middleware"
	_route "cozy-inn/app/route"
	_userUseCase "cozy-inn/businesses/users"
	_userController "cozy-inn/controller/users"
	_driverFactory "cozy-inn/driver"
	"fmt"

	firestore "cozy-inn/driver/firestore"
	util "cozy-inn/util"

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

	routeController := _route.ControllerList{
		LoggerMiddleware: configLogger.Init(),
		JWTMiddleware:    configJWT.Init(),
		UserController:   userController,
	}

	routeController.InitRoute(e)

	appPort := fmt.Sprintf(":%s", util.GetConfig("APP_PORT"))

	e.Logger.Fatal(e.Start(appPort))
}
