package main

import (
	"flag"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"periph.io/x/periph/host"

	"github.com/bodokaiser/beagle/driver/dds"
	"github.com/bodokaiser/beagle/driver/misc"
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

	err := misc.DefaultConfig.Exec()
	if err != nil {
		log.Fatal(err)
	}

	_, err = host.Init()
	if err != nil {
		log.Fatal(err)
	}

	s, err := misc.NewSelect()
	if err != nil {
		log.Fatal(err)
	}

	d, err := dds.NewAD9910(dds.DefaultSysClock, dds.DefaultRefClock)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(httpd.WrapContext)
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	sh := &handler.Spec{
		Specs: model.DefaultDDSSpecs,
	}
	e.GET("/specs", sh.List)

	dh := &handler.Device{
		Devices:     model.DefaultDDSDevices,
		ChipSelect:  s,
		Synthesizer: d,
	}
	e.GET("/devices", dh.List)
	e.PUT("/devices/:device", dh.Update)

	e.Static("/", "public")

	e.Logger.Fatal(e.Start(c.address))
}
