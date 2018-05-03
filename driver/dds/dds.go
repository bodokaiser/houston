// Package dds provides device drivers for direct digital synthesizer.
package dds

import (
	"github.com/bodokaiser/houston/device/dds"
	"github.com/bodokaiser/houston/driver"
	"github.com/bodokaiser/houston/driver/spi"
)

type Config struct {
	dds.Config `yaml:",inline"`
	SPI        spi.Config `yaml:"spi"`
	GPIO       GPIOConfig `yaml:"gpio"`
}

type GPIOConfig struct {
	Reset  string `yaml:"reset"`
	Update string `yaml:"update"`
}

type Param = dds.Param

var (
	ParamAmplitude = dds.ParamAmplitude
	ParamFrequency = dds.ParamFrequency
	ParamPhase     = dds.ParamPhase
)

type SweepConfig = dds.SweepConfig
type PlaybackConfig = dds.PlaybackConfig

type DDS interface {
	driver.Driver

	Reset(uint8) error

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
}
