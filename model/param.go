package model

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	validator "gopkg.in/go-playground/validator.v9"
)

type Mode int

const (
	ModeConst Mode = iota
	ModeSweep
	ModePlayback
)

func (m Mode) String() string {
	if m == ModeConst {
		return "const"
	}
	if m == ModeSweep {
		return "sweep"
	}
	if m == ModePlayback {
		return "playback"
	}

	panic("unknown mode")
}

func (m Mode) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}

func (m *Mode) UnmarshalJSON(b []byte) error {
	var s string

	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	switch strings.ToLower(s) {
	case "const":
		*m = ModeConst
	case "sweep":
		*m = ModeSweep
	case "playback":
		*m = ModePlayback
	default:
		return errors.New("invalid value")
	}

	return nil
}

// DDSParam is the DDS mode in which a DDS controllable parameter runs.
//
// Note that usually only one of the embedded structs will be not nil as it
// does not make sense to have a constant frequency which we is also swept.
type DDSParam struct {
	Mode     Mode         `json:"mode"`
	Const    *DDSConst    `json:"const,omitempty"`
	Sweep    *DDSSweep    `json:"sweep,omitempty"`
	Playback *DDSPlayback `json:"playback,omitempty"`
}

// DDSConst is the DDS mode where a DDS controllable parameter is constant.
type DDSConst struct {
	Value float64 `json:"value" validate:"gte=0"`
}

// DDSSweep is the DDS mode where a DDS controllable parameter is swept.
type DDSSweep struct {
	Limits   [2]float64    `json:"limits" validate:"len=2,range,dive,gte=0"`
	NoDwells [2]bool       `json:"nodwell"`
	Duration time.Duration `json:"duration" validate:"required,gt=0"`
}

// DDSPlayback is the DDS mode where a DDS controllable parameter is playbed
// back from memory.
type DDSPlayback struct {
	Trigger  bool          `json:"trigger"`
	Duplex   bool          `json:"duplex"`
	Interval time.Duration `json:"interval" validate:"gt=0"`
	Data     []float64     `json:"data" validate:"required"`
}

// DDSParamValidation implements struct level validation.
func DDSParamValidation(sl validator.StructLevel) {
	p := sl.Current().Interface().(DDSParam)

	switch p.Mode {
	case ModeConst:
		if p.Const == nil {
			sl.ReportError(p, "DDSParam", "", "required",
				"const mode was specified but not provided")
		} else {
			if err := sl.Validator().Struct(p.Const); err != nil {
				sl.ReportValidationErrors("Const.", "Const.", err.(validator.ValidationErrors))
			}
		}
	case ModeSweep:
		if p.Sweep == nil {
			sl.ReportError(p, "DDSParam", "", "required",
				"sweep mode was specified but not provided")
		} else {
			if err := sl.Validator().Struct(p.Sweep); err != nil {
				sl.ReportValidationErrors("Sweep.", "Sweep.", err.(validator.ValidationErrors))
			}
		}
	case ModePlayback:
		if p.Playback == nil {
			sl.ReportError(p, "DDSParam", "", "required",
				"playback mode was specified but not provided")
		} else {
			if err := sl.Validator().Struct(p.Playback); err != nil {
				sl.ReportValidationErrors("Playback.", "Playback.", err.(validator.ValidationErrors))
			}
		}
	}
}
