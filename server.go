package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/bodokaiser/beagle/httpd"
)

func main() {
	e := echo.New()
	e.Renderer = httpd.NewTemplate("views/*.html")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", httpd.IndexHandler)

	e.Static("/stylesheets", "public/stylesheets")
	e.Static("/javascripts", "public/javascripts")

	e.Logger.Fatal(e.Start(":8000"))
}
