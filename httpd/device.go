package httpd

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/bodokaiser/beagle/driver/dds"
	"github.com/bodokaiser/beagle/driver/misc"
	"github.com/bodokaiser/beagle/model"
)

type DeviceHandler struct {
	Devices     []model.Device
	ChipSelect  *misc.Select
	Synthesizer *dds.AD9910
}

// ListDevicesHandler responds a list of available devices.
func (h *DeviceHandler) List(ctx echo.Context) error {
	c := ctx.(*Context)

	if c.Accepts(echo.MIMEApplicationJSON) {
		return c.JSON(http.StatusOK, h.Devices)
	}

	return echo.ErrUnsupportedMediaType
}

type update struct {
	Frequency float64 `json:"frequency"`
}

// UpdateDeviceHandler updates configuration of specified device.
func (h *DeviceHandler) Update(ctx echo.Context) error {
	n, err := strconv.Atoi(ctx.Param("device"))
	if err != nil {
		return err
	}

	u := new(update)
	err = ctx.Bind(u)
	if err != nil {
		return err
	}

	err = h.ChipSelect.Address(uint(n))
	if err != nil {
		return err
	}

	err = h.Synthesizer.SingleTone(u.Frequency)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}
