// http starts a HTTP server with interface to the devices.
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

const (
	defaultSysClock = 1e9
	defaultRefClock = 1e7
)

const (
	defaultSPIDevice  = "SPI1.0"
	defaultSPIMaxFreq = 5e6
	defaultSPIMode    = 0
)

const (
	defaultResetPin    = "65"
	defaultIOUpdatePin = "27"
)

var defaultMuxPins = []string{
	"48", "30", "60", "31", "50",
}

type config struct {
	address string
	devices model.DDSDevices
}

func main() {
	c := config{}

	flag.StringVar(&c.address, "address", ":8000", "address to listen to")
	flag.Var(&c.devices, "devices", "devices to expose")
	flag.Parse()

	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	csel, err := mux.NewDigital(defaultMuxPins)
	if err != nil {
		log.Fatal(err)
	}

	dds, err := ad99xx.NewAD9910(ad99xx.Config{
		SysClock:    defaultSysClock,
		RefClock:    defaultRefClock,
		ResetPin:    defaultResetPin,
		IOUpdatePin: defaultIOUpdatePin,
		SPIDevice:   defaultSPIDevice,
		SPIMaxFreq:  defaultSPIMaxFreq,
		SPIMode:     defaultSPIMode,
	})
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(httpd.WrapContext)
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	dh := &handler.DDSDevices{
		Devices: c.devices,
		Driver: &driver.AD9910DDSArray{
			DDS: dds,
			Mux: csel,
		},
	}
	e.GET("/devices/dds", dh.List)
	e.PUT("/devices/dds/:name", dh.Update)

	e.Logger.Fatal(e.Start(c.address))
}
