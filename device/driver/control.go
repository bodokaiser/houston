package driver

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
)

// Control pins to use on the Beaglebone Black.
var (
	ControlIOUpdatePin    = "P8_17"
	ControlResetPin       = "P8_18"
	ControlExtIOUpdatePin = "P8_16"
	ControlExtRSTenPin    = "P8_15"
)

// ControlPulseWidth defines the time for a control pin to strobe.
var ControlPulseWidth = 1 * time.Millisecond

// Control is the gobot driver to handle io control.
type Control struct {
	name       string
	connection gpio.DigitalWriter
}

// NewControl returns a new initialized Control driver.
func NewControl(c gpio.DigitalWriter) *Control {
	return &Control{
		name:       gobot.DefaultName("Control"),
		connection: c,
	}
}

// Name returns the device name.
func (d *Control) Name() string {
	return d.name
}

// SetName sets the device name.
func (d *Control) SetName(s string) {
	d.name = s
}

// Start puts external control pins on high.
func (d *Control) Start() error {
	if err := d.connection.DigitalWrite(ControlExtRSTenPin, 1); err != nil {
		return err
	}

	return d.connection.DigitalWrite(ControlExtIOUpdatePin, 1)
}

// Halt does nothing but is required by the Driver interface.
func (d *Control) Halt() error {
	return nil
}

// Reset triggers the reset pin.
func (d *Control) Reset() error {
	if err := d.connection.DigitalWrite(ControlResetPin, 1); err != nil {
		return err
	}
	time.Sleep(ControlPulseWidth)

	return d.connection.DigitalWrite(ControlResetPin, 0)
}

// IOUpdate triggers the I/O update pin.
func (d *Control) IOUpdate() error {
	if err := d.connection.DigitalWrite(ControlIOUpdatePin, 1); err != nil {
		return err
	}
	time.Sleep(ControlPulseWidth)

	return d.connection.DigitalWrite(ControlIOUpdatePin, 0)
}

// Connection returns the gobot.Connection used for digital io.
func (d *Control) Connection() gobot.Connection {
	return d.connection.(gobot.Connection)
}
