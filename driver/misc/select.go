package misc

import (
	"errors"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
)

var selectPinNames = []string{"48", "30", "60", "31", "50"}

// Select is the device to manage chip select between the DDS devices.
type Select struct {
	pins [5]gpio.PinIO
}

// NewSelect returns a new Select.
func NewSelect() (d *Select) {
	for i, n := range selectPinNames {
		d.pins[i] = gpioreg.ByName(n)
	}

	return
}

// Init initializes the Select device.
func (d *Select) Init() error {
	return d.Address(0)
}

// Address configures chip select to address chip numbered with n.
func (d *Select) Address(n uint) (err error) {
	if n > (2 << uint(len(selectPinNames))) {
		return errors.New("chip number is out of range")
	}

	for i, p := range d.pins {
		var l gpio.Level

		if n&(1<<uint(i)) > 0 {
			l = gpio.High
		} else {
			l = gpio.Low
		}

		err = p.Out(l)
		if err != nil {
			return
		}
	}

	return
}
