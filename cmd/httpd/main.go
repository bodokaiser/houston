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
	e.Use(httpd.WrapContext)
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/specs", httpd.ListSpecsHandler)

	e.GET("/devices", httpd.ListDevicesHandler)
	e.GET("/devices/:device", httpd.ShowDeviceHandler)
	e.PUT("/devices/:device", httpd.UpdateDeviceHandler)

	e.Static("/", "public")

	e.Logger.Fatal(e.Start(c.address))
}
