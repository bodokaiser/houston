package misc

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
)

// IO control pins to use on the Beaglebone Black.
var (
	IOControlIntUpdatePin = "P8_17"
	IOControlIntResetPin  = "P8_18"
	IOControlExtUpdatePin = "P8_16"
	IOControlExtResetPin  = "P8_15"
)

// IOControlPulseWidth defines the time for a IOControl pin to strobe.
var IOControlPulseWidth = 1 * time.Millisecond

// IOControl is the gobot driver to handle io IOControl.
type IOControl struct {
	name       string
	connection gpio.DigitalWriter
}

// NewIOControl returns a new initialized IOControl driver.
func NewIOControl(c gpio.DigitalWriter) *IOControl {
	return &IOControl{
		name:       gobot.DefaultName("IOControl"),
		connection: c,
	}
}

// Name returns the device name.
func (d *IOControl) Name() string {
	return d.name
}

// SetName sets the device name.
func (d *IOControl) SetName(s string) {
	d.name = s
}

// Start puts external io control pins on high to be ignored.
func (d *IOControl) Start() (err error) {
	err = d.connection.DigitalWrite(IOControlExtResetPin, 1)

	if err != nil {
		return
	}

	return d.connection.DigitalWrite(IOControlExtIOUpdatePin, 1)
}

// Halt does nothing but is required by the Driver interface.
func (d *IOControl) Halt() error {
	return nil
}

// Reset triggers the reset pin.
func (d *IOControl) Reset() (err error) {
	err := d.connection.DigitalWrite(IOControlResetPin, 1)

	if err != nil {
		return
	}

	time.Sleep(IOControlPulseWidth)

	return d.connection.DigitalWrite(IOControlResetPin, 0)
}

// Update triggers the I/O update pin which commands the DDS chips to update
// there parameters from the registers.
func (d *IOControl) Update() (err error) {
	err = d.connection.DigitalWrite(IOControlIOUpdatePin, 1)

	if err != nil {
		return
	}

	time.Sleep(IOControlPulseWidth)

	return d.connection.DigitalWrite(IOControlIOUpdatePin, 0)
}

// Connection returns the gobot.Connection used for digital io.
func (d *IOControl) Connection() gobot.Connection {
	return d.connection.(gobot.Connection)
}
