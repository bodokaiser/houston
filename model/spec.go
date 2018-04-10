package model

// RangeSpec defines a specification consistent of minimum and maximum.
type RangeSpec struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

// SweepSpec defines allowed waveforms in sweep mode.
type SweepSpec struct {
	Modes []string `json:"modes"`
}

// DeviceSpec defines device specification.
type DeviceSpec struct {
	Frequency RangeSpec `json:"frequency"`
	Amplitude RangeSpec `json:"amplitude"`
	Modes     []string  `json:"modes"`
	Sweep     SweepSpec `json:"sweep"`
}

// AD9910DeviceSpec defines specifications of AD9910.
var AD9910DeviceSpec = &DeviceSpec{
	Frequency: RangeSpec{0, 400e6},
	Amplitude: RangeSpec{0.0, 1.0},
	Modes:     []string{"Single Tone", "Sweep"},
	Sweep: SweepSpec{
		Modes: []string{"Triangle", "Sawtooth"},
	},
}
