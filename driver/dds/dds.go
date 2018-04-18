// Package dds provides device drivers for direct digital synthesizer.
package dds

import "time"

// DDS is an implementation of a direct digital synthesizer device.
type DDS interface {
	// SingleTone configures the DDS to run in single tone mode.
	SingleTone(c SingleToneConfig) error

	SweepAmplitude(c DigitalRampConfig) error
	SweepFrequency(c DigitalRampConfig) error
	SweepPhase(c DigitalRampConfig) error

	// Playback configures the DDS to run in playback (RAM) mode.
	Playback(c PlaybackConfig) error
}

// ControlParam relates to the controllable output parameters.
type ControlParam uint

// Supported controllable output parameters.
const (
	Frequency ControlParam = iota
	PhaseOffset
	Amplitude
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
	NoDwell     [2]bool
	Duration    time.Duration
	Destination ControlParam
}

// PlaybackConfig holds the configuration for playback (RAM) mode.
type PlaybackConfig struct {
	SingleToneConfig
	Trigger     bool
	Duplex      bool
	Duration    time.Duration
	Data        []float64
	Destination ControlParam
}
