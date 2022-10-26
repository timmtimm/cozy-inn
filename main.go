package main

import (
	"cozy-inn/route"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	route.InitRoute(e)

	e.Logger.Fatal(e.Start(":3000"))
}
