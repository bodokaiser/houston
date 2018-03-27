package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/bodokaiser/beagle/httpd"
)

func main() {
	e := echo.New()

	e.Static("/", "public")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/foo", httpd.Handler)

	e.Logger.Fatal(e.Start(":8000"))
}
