// HTTP interface to devices.
package main

import (
	"periph.io/x/periph/host"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

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
	kingpin.Flag("address", "").Default(":8000").StringVar(&cmd.Address)
	kingpin.Flag("config", "").ExistingFileVar(&cmd.Filename)
	kingpin.Flag("debug", "").Default("false").BoolVar(&cmd.Debug)
	kingpin.Parse()

	cmd.Ensure()
	cmd.ReadFromBox(cmd.Filename)

	if _, err := host.Init(); err != nil {
		kingpin.FatalIfError(err, "host initialization")
	}

	h := &handler.DDSDevices{
		Devices: cmd.Devices,
		DDS:     ad9910.NewAD9910(cmd.DDS),
		Mux:     mux.NewDigital(cmd.Mux),
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

	e.Logger.Fatal(e.Start(cmd.Address))
}
