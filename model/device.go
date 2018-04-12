package model

// DDSDevice is a public entity to a digital synthesizer device.
//
// Instead of exposing an entity for every support operation mode we will
// decide what mode to use from the defined properties. For example in
// singletone mode it only makes sense to provide an amplitude and frequency
// property wherein in sweep mode we expect to have a frequency range defined
// over a single frequency.
type DDSDevice struct {
	Name           string     `json:"name"`
	Address        uint8      `json:"-"`
	Amplitude      float64    `json:"amplitude"`
	Frequency      float64    `json:"frequency"`
	FrequencyRange [2]float64 `json:"frequency"`
}

// DefaultDDSDevices are the default direct digital synthesizer available.
var DefaultDDSDevices = []DDSDevice{
	DDSDevice{
		Name:      "DDS0",
		Amplitude: 1.0,
		Frequency: 250e6,
	},
	DDSDevice{
		Name:           "DDS1",
		FrequencyRange: [2]float64{10e6, 20e6},
	},
}
