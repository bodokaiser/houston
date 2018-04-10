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
	"github.com/bodokaiser/beagle/model"
)

var defaultDevices = []model.Device{
	model.Device{
		ID:   0,
		Name: "DDS 0",
		Mode: "Sweep",
		SingleTone: model.SingleTone{
			Amplitude: 1.0,
			Frequency: 250e6,
		},
		Sweep: model.Sweep{
			StartFrequency: 100e6,
			StopFrequency:  200e6,
			Interval:       1,
			Waveform:       "Triangle",
		}},
	model.Device{
		ID:   1,
		Name: "DDS 1",
		Mode: "Single Tone",
		SingleTone: model.SingleTone{
			Amplitude: 1.0,
			Frequency: 30e6,
		},
		Sweep: model.Sweep{
			StartFrequency: 10e6,
			StopFrequency:  20e6,
			Interval:       .5,
			Waveform:       "Triangle",
		}},
}

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

	e.GET("/specs", httpd.ListSpecsHandler)

	dh := &httpd.DeviceHandler{
		Devices:     defaultDevices,
		ChipSelect:  s,
		Synthesizer: d,
	}

	e.GET("/devices", dh.List)
	e.PUT("/devices/:device", dh.Update)

	e.Static("/", "public")

	e.Logger.Fatal(e.Start(c.address))
}
