package driver

import (
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

// SingleTone implements dds.DDS.SingleTone.
func (d *AD9910DDSArray) SingleTone(c dds.SingleToneConfig) error {
	return d.DDS.SingleTone(c)
}

// MockedDDSArray mocks a DDSArray.
//
// As a pointer to MockedDDSArray implements the DDSArray interface it can
// be used as alternative implementation for example when developing on a
// machine which does not have the required sysfs interface or for just
// running some tests.
type MockedDDSArray struct {
	Address           uint8
	SingleToneConfig  dds.SingleToneConfig
	DigitalRampConfig dds.DigitalRampConfig
	PlaybackConfig    dds.PlaybackConfig
}

// Select implements the DDSArray interface.
//
// This will assign the structs Address field value to the given address and
// print the new address to stdout.
func (d *MockedDDSArray) Select(a uint8) error {
	d.Address = a

	return nil
}

// SingleTone implements the DDSArray interface.
//
// This will assign the structs Frequency field value to the given address and
// print the new frequency to stdout.
func (d *MockedDDSArray) SingleTone(c dds.SingleToneConfig) error {
	d.SingleToneConfig = c

	return nil
}

// DigitalRamp implements the DDSArray interface.
func (d *MockedDDSArray) DigitalRamp(c dds.DigitalRampConfig) error {
	d.DigitalRampConfig = c

	return nil
}

// Playback implements DDSArray interace.
func (d *MockedDDSArray) Playback(c dds.PlaybackConfig) error {
	d.PlaybackConfig = c

	return nil
}
