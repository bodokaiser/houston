package handler

import (
	"net/http"
	"time"

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

	i := h.Devices.FindByIDString(ctx.Param("id"))
	if i == -1 {
		return echo.ErrNotFound
	}

	if err := ctx.Bind(&d); err != nil {
		return err
	}

	if err := d.Validate(); err != nil {
		return err
	}
	h.Devices[i] = d

	if err := h.Mux.Select(d.ID); err != nil {
		return err
	}

	switch d.Amplitude.Mode {
	case model.ModeConst:
		h.DDS.SetAmplitude(d.Amplitude.Const.Value)
	case model.ModeSweep:
		h.DDS.SetSweep(dds.SweepConfig{
			Limits:   d.Amplitude.Sweep.Limits,
			NoDwells: d.Amplitude.Sweep.NoDwells,
			Duration: d.Amplitude.Sweep.Duration.Duration(),
			Param:    dds.ParamAmplitude,
		})
	case model.ModePlayback:
		h.DDS.SetPlayback(dds.PlaybackConfig{
			Trigger:  d.Amplitude.Playback.Trigger,
			Duplex:   d.Amplitude.Playback.Duplex,
			Interval: d.Amplitude.Playback.Interval.Duration(),
			Data:     d.Amplitude.Playback.Data,
			Param:    dds.ParamAmplitude,
		})
	}
	switch d.Frequency.Mode {
	case model.ModeConst:
		h.DDS.SetFrequency(d.Frequency.Const.Value)
	case model.ModeSweep:
		h.DDS.SetSweep(dds.SweepConfig{
			Limits:   d.Frequency.Sweep.Limits,
			NoDwells: d.Frequency.Sweep.NoDwells,
			Duration: d.Frequency.Sweep.Duration.Duration(),
			Param:    dds.ParamFrequency,
		})
	case model.ModePlayback:
		h.DDS.SetPlayback(dds.PlaybackConfig{
			Trigger:  d.Frequency.Playback.Trigger,
			Duplex:   d.Frequency.Playback.Duplex,
			Interval: d.Frequency.Playback.Interval.Duration(),
			Data:     d.Frequency.Playback.Data,
			Param:    dds.ParamFrequency,
		})
	}
	switch d.PhaseOffset.Mode {
	case model.ModeConst:
		h.DDS.SetPhaseOffset(d.PhaseOffset.Const.Value)
	case model.ModeSweep:
		h.DDS.SetSweep(dds.SweepConfig{
			Limits:   d.PhaseOffset.Sweep.Limits,
			NoDwells: d.PhaseOffset.Sweep.NoDwells,
			Duration: d.PhaseOffset.Sweep.Duration.Duration(),
			Param:    dds.ParamPhase,
		})
	case model.ModePlayback:
		h.DDS.SetPlayback(dds.PlaybackConfig{
			Trigger:  d.PhaseOffset.Playback.Trigger,
			Duplex:   d.PhaseOffset.Playback.Duplex,
			Interval: time.Duration(d.PhaseOffset.Playback.Interval),
			Data:     d.PhaseOffset.Playback.Data,
			Param:    dds.ParamPhase,
		})
	}

	if err := h.DDS.Exec(); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (h *DDSDevices) Delete(ctx echo.Context) error {
	i := h.Devices.FindByIDString(ctx.Param("id"))
	if i == -1 {
		return echo.ErrNotFound
	}
	d := h.Devices[i]

	if err := h.Mux.Select(d.ID); err != nil {
		return err
	}
	if err := h.DDS.Reset(); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}
