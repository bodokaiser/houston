// Package dds provides structures to mirror direct digital synthesizer devices.
package dds

import "time"

// Config holds the configuration of a generic direct digital synthesizer
// device.
type Config struct {
	Debug    bool    `yaml:"debug"`
	RefClock float64 `yaml:"refclock"`
	SysClock float64 `yaml:"sysclock"`
	PLL      bool    `yaml:"pll"`
	SPI3Wire bool    `yaml:"spi3wire"`
}

// Param defines a signal parameter controlled by the DDS device.
type Param int

// Common DDS controllable parameters.
const (
	ParamAmplitude Param = iota
	ParamFrequency
	ParamPhase
)

// SweepConfig holds the configuration for sweeping a DDS controllable
// parameter via the digital ramp of a DDS device.
type SweepConfig struct {
	Limits   [2]float64
	NoDwells [2]bool
	Duration time.Duration
	Param    Param
}

// PlaybackConfig holds the configuration for reading a DDS controllable
// parameter from memory.
type PlaybackConfig struct {
	Data     []float64
	Trigger  bool
	Duplex   bool
	Interval time.Duration
	Param    Param
}
