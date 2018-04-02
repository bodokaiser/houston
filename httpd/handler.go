package httpd

import (
	"net/http"

	"github.com/labstack/echo"
)

// ListDevicesHandler responds a list of available devices.
func ListDevicesHandler(ctx echo.Context) error {
	c := ctx.(*Context)

	if c.Accepts(echo.MIMEApplicationJSON) {
		return c.NoContent(http.StatusNoContent)
	}

	return echo.ErrUnsupportedMediaType
}

// ShowDeviceHandler responds configuration of the specified device.
func ShowDeviceHandler(ctx echo.Context) error {
	c := ctx.(*Context)

	if c.Accepts(echo.MIMEApplicationJSON) {
		return c.NoContent(http.StatusNoContent)
	}

	return echo.ErrUnsupportedMediaType
}

// UpdateDeviceHandler updates configuration of specified device.
func UpdateDeviceHandler(ctx echo.Context) error {
	c := ctx.(*Context)

	return c.NoContent(http.StatusNoContent)
}
