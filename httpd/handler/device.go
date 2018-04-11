package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/bodokaiser/beagle/driver/dds"
	"github.com/bodokaiser/beagle/driver/misc"
	"github.com/bodokaiser/beagle/httpd"
	"github.com/bodokaiser/beagle/model"
)

type Device struct {
	Devices     []model.DDSDevice
	ChipSelect  *misc.Select
	Synthesizer *dds.AD9910
}

func (h *Device) findByName(name string) *model.DDSDevice {
	for _, d := range h.Devices {
		if d.Name == name {
			return &d
		}
	}

	return nil
}

// List handles responds a list of available devices.
func (h *Device) List(ctx echo.Context) error {
	c := ctx.(*httpd.Context)

	if c.Accepts(echo.MIMEApplicationJSON) {
		return c.JSON(http.StatusOK, h.Devices)
	}
	if c.Accepts(echo.MIMEApplicationXML) {
		return c.XML(http.StatusOK, h.Devices)
	}

	return echo.ErrUnsupportedMediaType
}

// UpdateDeviceHandler updates configuration of specified device.
func (h *Device) Update(ctx echo.Context) error {
	d := h.findByName(ctx.Param("name"))
	if d == nil {
		return echo.ErrNotFound
	}

	err := ctx.Bind(d)
	if err != nil {
		return err
	}

	err = h.ChipSelect.Address(d.Address)
	if err != nil {
		return err
	}

	err = h.Synthesizer.SingleTone(d.Frequency)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}
