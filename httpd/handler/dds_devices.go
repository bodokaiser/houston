package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/bodokaiser/houston/driver/dds"
	"github.com/bodokaiser/houston/driver/mux"
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
	Mux     mux.Mux
	DDS     dds.DDS
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
	d := model.DDSDevice{}

	i := h.Devices.FindByName(ctx.Param("name"))
	if i == -1 {
		return echo.ErrNotFound
	}

	err := ctx.Bind(d)
	if err != nil {
		return err
	}

	err = d.Validate()
	if err != nil {
		return err
	}
	d.ID = h.Devices[i].ID
	h.Devices[i] = d

	if err := h.Mux.Select(d.ID); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (h *DDSDevices) Delete(ctx echo.Context) error {
	i := h.Devices.FindByName(ctx.Param("name"))
	if i == -1 {
		return echo.ErrNotFound
	}

	d := h.Devices[i]

	if err := h.Mux.Select(d.ID); err != nil {
		return err
	}
	if err := h.DDS.Reset(d.ID); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}
