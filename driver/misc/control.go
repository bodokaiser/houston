package misc

import (
	"errors"
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
)

var (
	intIOUpdatePinName = "27"
	extIOUpdatePinName = "46"
	intResetPinName    = "65"
	extResetPinName    = "47"
)

var defaultPulseWidth = 1 * time.Millisecond

// Control is the device to control IO updates and resets.
type Control struct {
	extIOUpdatePin gpio.PinIO
	extResetPin    gpio.PinIO
	intIOUpdatePin gpio.PinIO
	intResetPin    gpio.PinIO
}

// NewControl returns a new Control.
func NewControl() *Control {
	return &Control{
		extIOUpdatePin: gpioreg.ByName(extIOUpdatePinName),
		extResetPin:    gpioreg.ByName(extResetPinName),
		intIOUpdatePin: gpioreg.ByName(intIOUpdatePinName),
		intResetPin:    gpioreg.ByName(intResetPinName),
	}
}

// Init initializes the Control device.
func (d *Control) Init() (err error) {
	if d.extIOUpdatePin == nil {
		return errors.New("failed to find GPIO pin for external I/O update")
	}
	if d.extResetPin == nil {
		return errors.New("failed to find GPIO pin for external reset")
	}
	if d.intIOUpdatePin == nil {
		return errors.New("failed to find GPIO pin for internal I/O update")
	}
	if d.intResetPin == nil {
		return errors.New("failed to find GPIO pin for internal reset")
	}

	err = d.extIOUpdatePin.Out(gpio.High)
	if err != nil {
		return
	}

	err = d.extResetPin.Out(gpio.High)
	if err != nil {
		return
	}

	err = d.intIOUpdatePin.Out(gpio.Low)
	if err != nil {
		return
	}

	return d.intResetPin.Out(gpio.Low)
}

// IOUpdate triggers an I/O update which commands the connected DDS devices
// to read the updated configuration.
func (d *Control) IOUpdate() (err error) {
	return strobe(d.intIOUpdatePin, defaultPulseWidth)
}

// Reset triggers a reset which commands the connected DDS devices to clear
// the memory and reset the registers to the default values.
func (d *Control) Reset() (err error) {
	return strobe(d.intResetPin, defaultPulseWidth)
}

func strobe(p gpio.PinIO, d time.Duration) error {
	err := p.Out(gpio.High)
	if err != nil {
		return err
	}
	time.Sleep(d)

	return p.Out(gpio.Low)
}
