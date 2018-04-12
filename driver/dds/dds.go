package dds

// DDS is an implementation of a direct digital synthesizer device.
type DDS interface {
	// SingleTone configures the DDS to run in single tone mode with frequency.
	SingleTone(frequency float64) error
}
