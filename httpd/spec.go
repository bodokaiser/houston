package httpd

import (
	"net/http"

	"github.com/labstack/echo"
)

type RangeSpec struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

type SweepSpec struct {
	Modes []string `json:"modes"`
}

type DeviceSpec struct {
	Frequency RangeSpec `json:"frequency"`
	Amplitude RangeSpec `json:"amplitude"`
	Modes     []string  `json:"modes"`
	Sweep     SweepSpec `json:"sweep"`
}

var AD9910DeviceSpec = &DeviceSpec{
	Frequency: RangeSpec{0, 400e6},
	Amplitude: RangeSpec{-85, 0},
	Modes:     []string{"Single Tone", "Sweep"},
	Sweep: SweepSpec{
		Modes: []string{"Triangle", "Sawtooth"},
	},
}

// ListSpecsHandler responds the device specifications.
func ListSpecsHandler(ctx echo.Context) error {
	c := ctx.(*Context)

	if c.Accepts(echo.MIMEApplicationJSON) {
		return c.JSON(http.StatusOK, map[string]*DeviceSpec{
			"AD9910": AD9910DeviceSpec,
		})
	}

	return echo.ErrUnsupportedMediaType
}
