package main

import (
	"flag"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"periph.io/x/periph/host"

	"github.com/bodokaiser/beagle/driver"
	"github.com/bodokaiser/beagle/driver/dds/ad99xx"
	"github.com/bodokaiser/beagle/driver/mux"
	"github.com/bodokaiser/beagle/httpd"
	"github.com/bodokaiser/beagle/httpd/handler/device"
	"github.com/bodokaiser/beagle/model"
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

	dh := &device.DDS{
		Devices: model.DefaultDDSDevices,
		Driver: &driver.DDSArray{
			DDS: dds,
			Mux: csel,
		},
	}
	e.GET("/devices/dds", dh.List)
	e.PUT("/devices/dds/:name", dh.Update)

	e.Static("/", "public")

	e.Logger.Fatal(e.Start(c.address))
}
