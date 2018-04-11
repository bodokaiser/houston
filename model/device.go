package model

// DDSDevice is direct digital synthesizer device.
type DDSDevice struct {
	Name           string     `json:"name"`
	Address        uint8      `json:"address"`
	Amplitude      float64    `json:"amplitude"`
	AmplitudeRange [2]float64 `json:"amplitudeRange"`
	Frequency      float64    `json:"frequency"`
	FrequencyRange [2]float64 `json:"frequencyRange"`
}

// DefaultDDSDevices are the default direct digital synthesizer available.
var DefaultDDSDevices = []DDSDevice{
	DDSDevice{
		Name:      "DDS 0",
		Amplitude: 1.0,
		Frequency: 250e6,
	},
	DDSDevice{
		Name:           "DDS 1",
		FrequencyRange: [2]float64{10e6, 20e6},
	},
}
