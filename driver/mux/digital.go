package mux

import (
	"errors"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
)

// Digital implements Mux for a digital multiplexer.
type Digital struct {
	pins []gpio.PinIO
}

// NewDigital creates a new Digital multiplexer using the given pins in the
// given order for selection.
func NewDigital(pins []string) (*Digital, error) {
	d := &Digital{}
	d.pins = make([]gpio.PinIO, len(pins))

	for i, n := range pins {
		d.pins[i] = gpioreg.ByName(n)

		if d.pins[i] == nil {
			return nil, errors.New("invalid pin name")
		}
	}

	return d, nil
}

// Select selects the given address.
func (d *Digital) Select(address uint8) (err error) {
	if address > (2 << uint(len(d.pins))) {
		return errors.New("address is out of range")
	}

	for i, p := range d.pins {
		var l gpio.Level

		if address&(1<<uint(i)) > 0 {
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
