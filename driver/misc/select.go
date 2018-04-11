package misc

import (
	"errors"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
)

var pinNames = []string{"48", "30", "60", "31", "50"}

// Select is the device to manage chip select between the DDS devices.
type Select struct {
	pins [5]gpio.PinIO
}

// NewSelect returns a new Select.
func NewSelect() (*Select, error) {
	d := &Select{}

	for i, n := range pinNames {
		d.pins[i] = gpioreg.ByName(n)
	}

	err := d.Address(0)
	if err != nil {
		return nil, err
	}

	return d, nil
}

// Address configures chip select to address chip numbered with n.
func (d *Select) Address(n uint8) (err error) {
	if n > (2 << uint(len(pinNames))) {
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
