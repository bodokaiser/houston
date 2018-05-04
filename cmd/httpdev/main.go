// HTTP interface to mockup devices.
package main

import (
	"periph.io/x/periph/host"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/bodokaiser/houston/driver/dds"
	"github.com/bodokaiser/houston/driver/mux"
	"github.com/bodokaiser/houston/httpd"
	"github.com/bodokaiser/houston/httpd/handler"
	"github.com/bodokaiser/houston/model"
)

type options struct {
	Devices model.DDSDevices
	Address string
	Debug   bool
}

var cmd = options{
	Devices: model.DDSDevices{
		model.DDSDevice{
			Name: "Champ",
			Amplitude: model.DDSParam{
				Const: model.DDSConst{Value: 1.0},
			},
			Frequency: model.DDSParam{
				Const: model.DDSConst{Value: 200e6},
			},
		},
	},
}

func main() {
	kingpin.Flag("debug", "").Default("true").BoolVar(&cmd.Debug)
	kingpin.Flag("address", "").Default(":8000").StringVar(&cmd.Address)
	kingpin.Parse()

	if _, err := host.Init(); err != nil {
		kingpin.FatalIfError(err, "host initialization")
	}

	h := &handler.DDSDevices{
		Devices: cmd.Devices,
		DDS:     &dds.Mockup{Debug: cmd.Debug},
		Mux:     &mux.Mockup{Debug: cmd.Debug},
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
