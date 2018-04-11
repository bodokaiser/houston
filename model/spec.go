package model

// DDSSpec is the specifications of a direct digital synthesizer device.
type DDSSpec struct {
	Frequency [2]float64 `json:"frequency"`
	Amplitude [2]float64 `json:"amplitude"`
}

// DefaultDDSSpecs contains default specs.
var DefaultDDSSpecs = map[string]DDSSpec{
	"AD9910": DDSSpec{
		Frequency: [2]float64{0, 400e6},
		Amplitude: [2]float64{0, 1},
	},
}
