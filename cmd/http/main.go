package main

import (
	"flag"
	"log"

	"periph.io/x/periph/host"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/bodokaiser/houston/driver"
	"github.com/bodokaiser/houston/driver/dds/ad99xx"
	"github.com/bodokaiser/houston/driver/mux"
	"github.com/bodokaiser/houston/httpd"
	"github.com/bodokaiser/houston/httpd/handler"
	"github.com/bodokaiser/houston/model"
)

type config struct {
	address string
}

func main() {
	c := config{}

	flag.StringVar(&c.address, "address", ":8000", "")
	flag.Parse()

	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	csel, err := mux.NewDigital(mux.DefaultDigitalPins)
	if err != nil {
		log.Fatal(err)
	}

	dds, err := ad99xx.NewAD9910(ad99xx.AD9910DefaultConfig)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(httpd.WrapContext)
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	dh := &handler.DDSDevices{
		Devices: model.DefaultDDSDevices,
		Driver: &driver.AD9910DDSArray{
			DDS: dds,
			Mux: csel,
		},
	}
	e.GET("/devices/dds", dh.List)
	e.PUT("/devices/dds/:name", dh.Update)

	e.Logger.Fatal(e.Start(c.address))
}
