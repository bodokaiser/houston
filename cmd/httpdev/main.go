package main

import (
	"flag"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/bodokaiser/beagle/driver"
	"github.com/bodokaiser/beagle/httpd"
	"github.com/bodokaiser/beagle/httpd/handler"
	"github.com/bodokaiser/beagle/model"
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

	dh := &handler.DDSDevices{
		Devices: model.DefaultDDSDevices,
		Driver:  &driver.MockedDDSArray{},
	}
	e.GET("/devices/dds", dh.List)
	e.PUT("/devices/dds/:name", dh.Update)

	e.Logger.Fatal(e.Start(c.address))
}
