package main

import (
	_middleware "cozy-inn/middleware"
	_route "cozy-inn/route"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	configLogger := _middleware.ConfigLogger{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	routeController := _route.ControllerList{
		LoggerMiddleware: configLogger.InitLogger(),
	}

	routeController.InitRoute(e)

	e.Logger.Fatal(e.Start(":3000"))
}
