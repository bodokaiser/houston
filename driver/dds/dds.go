// Package dds provides device drivers for direct digital synthesizer.
package dds

import (
	"github.com/bodokaiser/houston/device/dds"
	"github.com/bodokaiser/houston/driver"
	"github.com/bodokaiser/houston/driver/spi"
)

// Config extends dds.Config with hardware specific configuration.
type Config struct {
	dds.Config `yaml:",inline"`
	SPI        spi.Config `yaml:"spi"`
	GPIO       GPIOConfig `yaml:"gpio"`
}

// GPIOConfig defines the GPIO pins to as update and reset lines.
type GPIOConfig struct {
	Reset  string `yaml:"reset"`
	Update string `yaml:"update"`
}

// Param alias dds.Param.
type Param = dds.Param

// Params alias dds.Params.
var (
	ParamAmplitude = dds.ParamAmplitude
	ParamFrequency = dds.ParamFrequency
	ParamPhase     = dds.ParamPhase
)

// SweepConfig aliases dds.SweepConfig.
type SweepConfig = dds.SweepConfig

// PlaybackConfig aliases dds.PlaybackConfig.
type PlaybackConfig = dds.PlaybackConfig

// DDS interface represents a direct digital synthesizer device.
type DDS interface {
	driver.Driver

	Reset() error

	Amplitude() float64
	SetAmplitude(float64)

	Frequency() float64
	SetFrequency(float64)

	PhaseOffset() float64
	SetPhaseOffset(float64)

	Sweep() SweepConfig
	SetSweep(SweepConfig)

	Playback() PlaybackConfig
	SetPlayback(PlaybackConfig)

	Exec() error
}
