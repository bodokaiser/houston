package model

import validator "gopkg.in/go-playground/validator.v9"

// DDSDevice is a public entity to a digital synthesizer device.
//
// Instead of exposing an entity for every support operation mode we will
// decide what mode to use from the defined properties. For example in
// singletone mode it only makes sense to provide an amplitude and frequency
// property wherein in sweep mode we expect to have a frequency range defined
// over a single frequency.
type DDSDevice struct {
	ID          uint8    `json:"id"        validate:"max=31"`
	Name        string   `json:"name"      validate:"required"`
	Amplitude   DDSParam `json:"amplitude"`
	Frequency   DDSParam `json:"frequency"`
	PhaseOffset DDSParam `json:"phase"`
}

// Validate returns an error if DDSDevice has ambigious configuration.
func (d *DDSDevice) Validate() (err error) {
	return validate.Struct(d)
}

// DDSDeviceValidation implements struct level validation.
func DDSDeviceValidation(sl validator.StructLevel) {
	d := sl.Current().Interface().(DDSDevice)

	a, b := 0, 0

	switch d.Amplitude.Mode {
	case ModeSweep:
		a++
	case ModePlayback:
		b++
	}
	switch d.Frequency.Mode {
	case ModeSweep:
		a++
	case ModePlayback:
		b++
	}
	switch d.PhaseOffset.Mode {
	case ModeSweep:
		a++
	case ModePlayback:
		b++
	}

	if a > 1 {
		sl.ReportError(d, "DDSParam", "", "toomanysweeps",
			"only one parameter can be in sweep mode")
	}
	if b > 1 {
		sl.ReportError(d, "DDSParam", "", "toomanyplaybacks",
			"only one parameter can be in playback mode")
	}
}
