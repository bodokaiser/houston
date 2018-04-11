package model

// Device is a generic device.
type Device struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// DDSDevice is direct digital synthesizer device.
type DDSDevice struct {
	Device
	Amplitude      float64    `json:"amplitude"`
	AmplitudeRange [2]float64 `json:"amplitudeRange"`
	Frequency      float64    `json:"frequency"`
	FrequencyRange [2]float64 `json:"frequencyRange"`
}
