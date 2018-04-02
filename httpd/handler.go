package httpd

import (
	"net/http"

	"github.com/labstack/echo"
)

// Device is a device exposed by the HTTP api.
type Device struct {
	Name string `json:"name"`
	Mode string `json:"mode"`
}

var defaultDevices = []Device{
	Device{"Signal Generator 1a", "Const"},
	Device{"Signal Generator 1b", "Const"},
	Device{"Signal Generator 2a", "Sweep"},
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
