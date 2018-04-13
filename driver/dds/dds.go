// Package dds provides device drivers for direct digital synthesizer.
package dds

import "errors"

// DDS is an implementation of a direct digital synthesizer device.
type DDS interface {
	// SingleTone configures the DDS to run in single tone mode with given
	// frequency.
	SingleTone(amplitude float64, frequency float64, phaseOffset float64) error
}

// Common errors for DDS.
var (
	ErrInvalidAmplitude = errors.New("invalid amplitude")
	ErrInvalidFrequency = errors.New("invalid frequency")
	ErrInvalidPhase     = errors.New("invalid phase")
)
