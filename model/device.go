package model

// Device is a device exposed by the HTTP api.
type Device struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Mode       string     `json:"mode"`
	SingleTone SingleTone `json:"singleTone"`
	Sweep      Sweep      `json:"sweep"`
}

// SingleTone contains single tone configuration parameters.
type SingleTone struct {
	Amplitude float32 `json:"amplitude"`
	Frequency float64 `json:"frequency"`
}

// Sweep contains sweep configuration parameters.
type Sweep struct {
	StartFrequency float32 `json:"startFrequency"`
	StopFrequency  float32 `json:"stopFrequency"`
	Interval       float32 `json:"interval"`
	Waveform       string  `json:"waveform"`
}
