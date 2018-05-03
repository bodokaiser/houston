// http starts a HTTP server with interface to the devices.
package main

import (
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"periph.io/x/periph/host"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/bodokaiser/houston/cmd"
	"github.com/bodokaiser/houston/driver/dds/ad9910"
	"github.com/bodokaiser/houston/driver/mux"
	"github.com/bodokaiser/houston/httpd"
	"github.com/bodokaiser/houston/httpd/handler"
	"github.com/bodokaiser/houston/model"
)

var config = cmd.Config{}

var devices model.DDSDevices

var address string

func main() {
	kingpin.Flag("address", "").Default(":8000").StringVar(&address)
	kingpin.Flag("config", "").Default("config.yaml").ExistingFileVar(&config.Filename)
	kingpin.Flag("debug", "").Default("false").BoolVar(&config.DDS.Debug)
	kingpin.Parse()

	config.Mux.Debug = config.DDS.Debug

	kingpin.FatalIfError(config.ReadFromFile(), "config")

	if _, err := host.Init(); err != nil {
		kingpin.FatalIfError(err, "host initialization")
	}

	h := &handler.DDSDevices{
		Devices: devices,
		DDS:     ad9910.NewAD9910(config.DDS),
		Mux:     mux.NewDigital(config.Mux),
	}

	kingpin.FatalIfError(h.DDS.Init(), "mux initialization")
	kingpin.FatalIfError(h.Mux.Init(), "dds initialization")

	e := echo.New()
	e.Use(httpd.WrapContext)
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = handler.HTTPError

	e.GET("/devices/dds", h.List)
	e.PUT("/devices/dds/:id", h.Update)
	e.DELETE("/devices/dds/:id", h.Delete)

	e.Static("/", "public")

	e.Logger.Fatal(e.Start(address))
}
