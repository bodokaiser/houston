// Package dds provides device drivers for direct digital synthesizer.
package dds

import "errors"

// Common errors for DDS.
var (
	ErrInvalidAmplitude = errors.New("invalid amplitude")
	ErrInvalidFrequency = errors.New("invalid frequency")
	ErrInvalidPhase     = errors.New("invalid phase")
)

// DDS is an implementation of a direct digital synthesizer device.
type DDS interface {
	// SingleTone configures the DDS to run in single tone mode.
	SingleTone(c SingleToneConfig) error

	// DigitalRamp configures the DDS to run in digital ramp mode.
	DigitalRamp(c DigitalRampConfig) error

	// Playback configures the DDS to run in playback (RAM) mode.
	Playback(c PlaybackConfig) error
}

// ControlParam relates to the controllable output parameters.
type ControlParam uint

// Supported controllable output parameters.
const (
	Amplitude ControlParam = iota
	Frequency
	PhaseOffset
)

// SingleToneConfig holds the configuration for single tone mode.
type SingleToneConfig struct {
	Amplitude   float64
	Frequency   float64
	PhaseOffset float64
}

// DigitalRampConfig holds the configuration for digital ramp mode.
type DigitalRampConfig struct {
	SingleToneConfig
	Limits      [2]float64
	StepSize    [2]float64
	SlopeRate   [2]float64
	NoDwellHigh bool
	NoDwellLow  bool
	Destination ControlParam
}

// PlaybackConfig holds the configuration for playback (RAM) mode.
type PlaybackConfig struct {
	SingleToneConfig
	Trigger     bool
	Duplex      bool
	Rate        float64
	Data        []byte
	Destination ControlParam
}
