package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/bodokaiser/beagle/httpd"
	"github.com/bodokaiser/beagle/model"
)

type Spec struct {
	Specs map[string]model.DDSSpec
}

// List responds the device specifications.
func (h *Spec) List(ctx echo.Context) error {
	c := ctx.(*httpd.Context)

	if c.Accepts(echo.MIMEApplicationJSON) {
		return c.JSON(http.StatusOK, model.DefaultDDSSpecs)
	}
	if c.Accepts(echo.MIMEApplicationXML) {
		return c.XML(http.StatusOK, model.DefaultDDSSpecs)
	}

	return echo.ErrUnsupportedMediaType
}
