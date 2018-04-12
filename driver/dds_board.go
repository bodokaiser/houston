package driver

import (
	"github.com/bodokaiser/beagle/driver/dds"
	"github.com/bodokaiser/beagle/driver/mux"
)

// DDSArray combines a multiplexer and a direct digital synthesizer driver
// to interface an array of direct digital synthesizer.
type DDSArray struct {
	DDS dds.DDS
	Mux mux.Mux
}

// Select configures the multiplexer to address given address.
func (d *DDSArray) Select(address uint8) error {
	return d.Mux.Select(address)
}

// SingleTone configures the addressed dds to run in single tone mode with
// given frequency.
func (d *DDSArray) SingleTone(frequency float64) error {
	return d.DDS.SingleTone(frequency)
}
