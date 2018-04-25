package dds

import "time"

type Config struct {
	RefClock float64 `yaml:"refclock"`
	SysClock float64 `yaml:"sysclock"`
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
	Duration time.Duration
	Param    Param
}