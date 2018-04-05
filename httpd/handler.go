package httpd

import (
	"net/http"

	"github.com/labstack/echo"
)

var defaultDevices = []*Device{
	&Device{0, "DDS 0", "Sweep", SingleTone{0, 250e6}, Sweep{100e6, 200e6, 1, "Triangle"}},
	&Device{1, "DDS 1", "Single Tone", SingleTone{-80, 30e6}, Sweep{10e6, 20e6, .5, "Triangle"}},
}

// Device is a device exposed by the HTTP api.
type Device struct {
	Id         int        `json:"id"`
	Name       string     `json:"name"`
	Mode       string     `json:"mode"`
	SingleTone SingleTone `json:"singleTone"`
	Sweep      Sweep      `json:"sweep"`
}

type SingleTone struct {
	Amplitude float32 `json:"amplitude"`
	Frequency float64 `json:"frequency"`
}

type Sweep struct {
	StartFrequency float32 `json:"startFrequency"`
	StopFrequency  float32 `json:"stopFrequency"`
	Interval       float32 `json:"interval"`
	Waveform       string  `json:"waveform"`
}

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

// ListDevicesHandler responds a list of available devices.
func ListDevicesHandler(ctx echo.Context) error {
	c := ctx.(*Context)

	if c.Accepts(echo.MIMEApplicationJSON) {
		return c.JSON(http.StatusOK, defaultDevices)
	}

	return echo.ErrUnsupportedMediaType
}

// ShowDeviceHandler responds configuration of the specified device.
func ShowDeviceHandler(ctx echo.Context) error {
	c := ctx.(*Context)

	if c.Accepts(echo.MIMEApplicationJSON) {
		return c.JSON(http.StatusOK, defaultDevices[0])
	}

	return echo.ErrUnsupportedMediaType
}

// UpdateDeviceHandler updates configuration of specified device.
func UpdateDeviceHandler(ctx echo.Context) error {
	c := ctx.(*Context)

	return c.NoContent(http.StatusNoContent)
}
