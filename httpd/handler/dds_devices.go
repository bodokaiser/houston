package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/bodokaiser/houston/driver"
	"github.com/bodokaiser/houston/httpd"
	"github.com/bodokaiser/houston/model"
)

// DDSDevices has HTTP handlers to interact with a DDS array.
//
// The Devices field contains the list of available devices which will be kept
// in memory to store the recent configuration.
// The Driver field contains the interface to the dds array.
type DDSDevices struct {
	Devices model.DDSDevices
	Driver  driver.DDSArray
}

// List handles responds a list of available devices.
func (h *DDSDevices) List(ctx echo.Context) error {
	c := ctx.(*httpd.Context)

	if c.Accepts(echo.MIMEApplicationJSON) {
		return c.JSON(http.StatusOK, h.Devices)
	}
	if c.Accepts(echo.MIMEApplicationXML) {
		return c.XML(http.StatusOK, h.Devices)
	}

	return echo.ErrUnsupportedMediaType
}

// Update updates configuration of specified device.
func (h *DDSDevices) Update(ctx echo.Context) error {
	i := h.Devices.FindByName(ctx.Param("name"))
	if i == -1 {
		return echo.ErrNotFound
	}

	err := ctx.Bind(&h.Devices[i])
	if err != nil {
		return err
	}
	d := h.Devices[i]

	err = h.Driver.Select(d.Address)
	if err != nil {
		return err
	}

	err = h.Driver.SingleTone(d.Frequency)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}
