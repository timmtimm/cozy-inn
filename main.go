package main

import (
	_middleware "cozy-inn/app/middleware"
	_route "cozy-inn/app/route"
	_driverFactory "cozy-inn/driver"

	firestore "cozy-inn/driver/firestore"
	util "cozy-inn/util"

	firebase "firebase.google.com/go"
	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	FirestoreConfig := &firebase.Config{ProjectID: util.GetFirebaseEnv("project_id")}

	firestore := firestore.InitFirestore(FirestoreConfig)
	UserRepository := _driverFactory.NewUserRepository(firestore)

	configLogger := _middleware.ConfigLogger{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	routeController := _route.ControllerList{
		LoggerMiddleware: configLogger.InitLogger(),
		UserRepository:   UserRepository,
	}

	routeController.InitRoute(e)

	e.Logger.Fatal(e.Start(":3000"))
}
