package driver

import (
	"fmt"

	"github.com/bodokaiser/houston/driver/dds"
	"github.com/bodokaiser/houston/driver/mux"
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
func (d *AD9910DDSArray) Select(a uint8) error {
	return d.Mux.Select(a)
}

// SingleTone configures the addressed dds to run in single tone mode with
// given frequency.
func (d *AD9910DDSArray) SingleTone(a float64, f float64, p float64) error {
	return d.DDS.SingleTone(a, f, p)
}

// MockedDDSArray mocks a DDSArray.
//
// As a pointer to MockedDDSArray implements the DDSArray interface it can
// be used as alternative implementation for example when developing on a
// machine which does not have the required sysfs interface or for just
// running some tests.
type MockedDDSArray struct {
	Frequency float64
	Amplitude float64
	Phase     float64
	Address   uint8
}

// Select implements the DDSArray interface.
//
// This will assign the structs Address field value to the given address and
// print the new address to stdout.
func (d *MockedDDSArray) Select(a uint8) error {
	d.Address = a

	fmt.Printf("selected address %v\n", a)

	return nil
}

// SingleTone implements the DDSArray interface.
//
// This will assign the structs Frequency field value to the given address and
// print the new frequency to stdout.
func (d *MockedDDSArray) SingleTone(a float64, f float64, p float64) error {
	d.Phase = p
	d.Amplitude = a
	d.Frequency = f

	fmt.Printf("running single tone at frequency %v\n", f)

	return nil
}
