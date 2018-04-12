package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/bodokaiser/beagle/driver"
	"github.com/bodokaiser/beagle/httpd"
	"github.com/bodokaiser/beagle/model"
)

// DDSDevices exposes HTTP handlers to interact with a DDS array.
type DDSDevices struct {
	Devices []model.DDSDevice
	Driver  driver.DDSArray
}

func (h *DDSDevices) findByName(name string) *model.DDSDevice {
	for _, d := range h.Devices {
		if d.Name == name {
			return &d
		}
	}

	return nil
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
	d := h.findByName(ctx.Param("name"))
	if d == nil {
		return echo.ErrNotFound
	}

	err := ctx.Bind(d)
	if err != nil {
		return err
	}

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
