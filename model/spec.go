package model

// DDSSpec is the specifications of a direct digital synthesizer device.
type DDSSpec struct {
	Frequency [2]float64 `json:"frequency"`
	Amplitude [2]float64 `json:"amplitude"`
}

// AD9910DDSSpec is the specification of the AD9910 direct digital synthesizer.
var AD9910DDSSpec = &DDSSpec{
	Frequency: [2]float64{0, 400e6},
	Amplitude: [2]float64{0, 1},
}
