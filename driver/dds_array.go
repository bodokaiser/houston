package driver

import (
	"github.com/bodokaiser/beagle/driver/dds"
	"github.com/bodokaiser/beagle/driver/mux"
)

// DDSArray combines a multiplexer and a direct digital synthesizer driver
// to interface an array of direct digital synthesizer.
type DDSArray interface {
	mux.Mux
	dds.DDS
}

// AD9910DDSArray combines a multiplexer and a direct digital synthesizer driver
// to interface an array of direct digital synthesizer.
type AD9910DDSArray struct {
	DDS dds.DDS
	Mux mux.Mux
}

// Select configures the multiplexer to address given address.
func (d *AD9910DDSArray) Select(address uint8) error {
	return d.Mux.Select(address)
}

// SingleTone configures the addressed dds to run in single tone mode with
// given frequency.
func (d *AD9910DDSArray) SingleTone(frequency float64) error {
	return d.DDS.SingleTone(frequency)
}

// MockedDDSArray mocks a DDSArray. Useful for development and testing.
type MockedDDSArray struct {
	Frequency float64
	Address   uint8
}

// Select implements DDSArray.
func (d *MockedDDSArray) Select(address uint8) error {
	d.Address = address

	return nil
}

// SingleTone implements DDSArray.
func (d *MockedDDSArray) SingleTone(frequency float64) error {
	d.Frequency = frequency

	return nil
}
