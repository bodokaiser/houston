// httpdev starts a HTTP server with mocked hardware.
package main

import (
	"flag"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/bodokaiser/houston/driver"
	"github.com/bodokaiser/houston/httpd"
	"github.com/bodokaiser/houston/httpd/handler"
	"github.com/bodokaiser/houston/model"
)

var defaultDDSDevices = []model.DDSDevice{
	model.DDSDevice{
		Name:      "DDS0",
		Address:   0,
		Amplitude: 1.0,
		Frequency: 250e6,
	},
	model.DDSDevice{
		Name:      "DDS1",
		Address:   1,
		Frequency: 10e6,
	},
}

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
	e.HTTPErrorHandler = handler.HTTPError

	dh := &handler.DDSDevices{
		Devices: defaultDDSDevices,
		Driver:  &driver.MockedDDSArray{},
	}
	e.GET("/devices/dds", dh.List)
	e.PUT("/devices/dds/:name", dh.Update)

	e.Logger.Fatal(e.Start(c.address))
}
