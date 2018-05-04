package model

import (
	"encoding/json"
	"errors"
	"math"
	"strings"
	"time"

	validator "gopkg.in/go-playground/validator.v9"
)

// Mode is a DDS parameter mode.
type Mode int

// Available DDS parameter modes.
const (
	ModeConst Mode = iota
	ModeSweep
	ModePlayback
)

// String returns the string representation of mode.
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

// MarshalJSON implements json.Marshaller interface.
func (m Mode) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}

// UnmarshalJSON implements json.Unmarshaller interface.
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

// Seconds represents a time duration in seconds.
//
// Actually we would use time.Duration for this case but unfortunately
// time.Duration states strange json requirements which we do not support.
type Seconds float64

// Duration returns a time.Duration with the same duration as seconds.
func (s Seconds) Duration() time.Duration {
	return time.Duration(math.Round(float64(time.Second) * float64(s)))
}

// String returns the string representation of seconds.
func (s Seconds) String() string {
	return s.Duration().String()
}

// Set sets the seconds from a duration time string.
func (s *Seconds) Set(v string) error {
	d, err := time.ParseDuration(v)
	if err != nil {
		return err
	}

	*s = Seconds(d.Seconds())

	return nil
}

// DDSParam is the DDS mode in which a DDS controllable parameter runs.
//
// Note that usually only one of the embedded structs will be not nil as it
// does not make sense to have a constant frequency which we is also swept.
type DDSParam struct {
	Mode     Mode        `json:"mode"`
	Const    DDSConst    `json:"const" validate:"structonly"`
	Sweep    DDSSweep    `json:"sweep" validate:"structonly"`
	Playback DDSPlayback `json:"playback" validate:"structonly"`
}

// DDSConst is the DDS mode where a DDS controllable parameter is constant.
type DDSConst struct {
	Value float64 `json:"value" validate:"gte=0"`
}

// DDSSweep is the DDS mode where a DDS controllable parameter is swept.
type DDSSweep struct {
	Limits   [2]float64 `json:"limits" validate:"len=2,range,dive,gte=0"`
	NoDwells [2]bool    `json:"nodwells"`
	Duration Seconds    `json:"duration" validate:"required,gt=0"`
}

// DDSPlayback is the DDS mode where a DDS controllable parameter is playbed
// back from memory.
type DDSPlayback struct {
	Trigger  bool      `json:"trigger"`
	Duplex   bool      `json:"duplex"`
	Interval Seconds   `json:"interval" validate:"gt=0"`
	Data     []float64 `json:"data" validate:"required"`
}

// DDSParamValidation implements struct level validation.
func DDSParamValidation(sl validator.StructLevel) {
	p := sl.Current().Interface().(DDSParam)

	switch p.Mode {
	case ModeConst:
		if err := sl.Validator().Struct(p.Const); err != nil {
			sl.ReportValidationErrors("Const.", "Const.", err.(validator.ValidationErrors))
		}
	case ModeSweep:
		if err := sl.Validator().Struct(p.Sweep); err != nil {
			sl.ReportValidationErrors("Sweep.", "Sweep.", err.(validator.ValidationErrors))
		}
	case ModePlayback:
		if err := sl.Validator().Struct(p.Playback); err != nil {
			sl.ReportValidationErrors("Playback.", "Playback.", err.(validator.ValidationErrors))
		}
	}
}
