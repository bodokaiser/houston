package dds

import "time"

type Config struct {
	Debug    bool    `yaml:"debug"`
	RefClock float64 `yaml:"refclock"`
	SysClock float64 `yaml:"sysclock"`
	PLL      bool    `yaml:"pll"`
	SPI3Wire bool    `yaml:"spi3wire"`
}

type Param int

const (
	ParamAmplitude Param = iota
	ParamFrequency
	ParamPhase
)

type SweepConfig struct {
	Limits   [2]float64
	NoDwells [2]bool
	Duration time.Duration
	Param    Param
}

type PlaybackConfig struct {
	Data     []float64
	Trigger  bool
	Duplex   bool
	Interval time.Duration
	Param    Param
}
