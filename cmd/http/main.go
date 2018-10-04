// HTTP interface to devices.
package main

import (
	"flag"
	"log"

	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/host"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/bodokaiser/houston/config"
	"github.com/bodokaiser/houston/driver/dds/ad9910"
	"github.com/bodokaiser/houston/driver/mux"
	"github.com/bodokaiser/houston/httpd"
	"github.com/bodokaiser/houston/httpd/handler"
	"github.com/bodokaiser/houston/model"
)

type options struct {
	config.Config

	Devices  model.DDSDevices
	Address  string
	Filename string
}

var cmd = options{}

func main() {
	flag.StringVar(&cmd.Address, "address", ":6200", "")
	flag.StringVar(&cmd.Filename, "config", cmd.Filename, "")
	flag.BoolVar(&cmd.Debug, "debug", false, "")
	flag.Var(&cmd.Devices, "devices", "")
	flag.Parse()

	cmd.Ensure()
	cmd.ReadFromBox(cmd.Filename)

	if _, err := host.Init(); err != nil {
		log.Fatalf("host init: %s", err)
	}

	cmd.DDS.Config.SPIPort, _ = spireg.Open(cmd.DDS.SPI.Device)
	cmd.DDS.Config.ResetPin = gpioreg.ByName(cmd.DDS.GPIO.Reset)
	cmd.DDS.Config.UpdatePin = gpioreg.ByName(cmd.DDS.GPIO.Update)

	h := &handler.DDSDevices{
		Devices: cmd.Devices,
		DDS:     ad9910.NewAD9910(cmd.DDS.Config),
		Mux:     mux.NewDigital(cmd.Mux),
		Debug:   cmd.DDS.Config.Debug,
	}

	if err := h.DDS.Init(); err != nil {
		log.Fatalf("dds init: %s", err)
	}
	if err := h.Mux.Init(); err != nil {
		log.Fatalf("mux init: %s", err)
	}

	e := echo.New()
	e.Use(httpd.WrapContext)
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = handler.HTTPError

	e.GET("/devices/dds", h.List)
	e.PUT("/devices/dds/:id", h.Update)
	e.DELETE("/devices/dds/:id", h.Delete)
	e.POST("/devices/dds/:id/trigger", h.Trigger)

	e.Static("/", "public")
	e.File("/devices", "public/index.html")

	e.Logger.Fatal(e.Start(cmd.Address))
}
