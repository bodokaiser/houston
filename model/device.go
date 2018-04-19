package model

// DDSDevice is a public entity to a digital synthesizer device.
//
// Instead of exposing an entity for every support operation mode we will
// decide what mode to use from the defined properties. For example in
// singletone mode it only makes sense to provide an amplitude and frequency
// property wherein in sweep mode we expect to have a frequency range defined
// over a single frequency.
type DDSDevice struct {
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	Amplitude   DDSParam `json:"amplitude"`
	Frequency   DDSParam `json:"frequency"`
	PhaseOffset DDSParam `json:"phase"`
}

// Validate returns an error if DDSDevice has ambigious configuration.
func (d *DDSDevice) Validate() (err error) {
	return validate.Struct(d)
}
