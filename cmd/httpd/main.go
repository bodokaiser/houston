package main

import (
	"flag"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/bodokaiser/beagle/httpd"
)

type config struct {
	address string
}

func main() {
	c := config{}

	flag.StringVar(&c.address, "address", ":8000", "")
	flag.Parse()

	e := echo.New()
	e.Renderer = httpd.NewTemplate("views/*.html")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", httpd.IndexHandler)
	e.PATCH("/", httpd.PatchHandler)

	e.Static("/stylesheets", "public/stylesheets")
	e.Static("/javascripts", "public/javascripts")

	e.Logger.Fatal(e.Start(c.address))
}
