package httpd

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/bodokaiser/beagle/model"
)

// ListSpecsHandler responds the device specifications.
func ListSpecsHandler(ctx echo.Context) error {
	c := ctx.(*Context)

	if c.Accepts(echo.MIMEApplicationJSON) {
		return c.JSON(http.StatusOK, map[string]*model.DeviceSpec{
			"AD9910": model.AD9910DeviceSpec,
		})
	}

	return echo.ErrUnsupportedMediaType
}
