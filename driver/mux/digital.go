package mux

import (
	"errors"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
)

// Digital implements Mux for a digital multiplexer.
type Digital struct {
	config Config
	pins   []gpio.PinIO
}

// NewDigital creates a new Digital multiplexer using the given pins in the
// given order for selection.
func NewDigital(c Config) *Digital {
	return &Digital{
		config: c,
	}
}

func (d *Digital) Init() error {
	d.pins = make([]gpio.PinIO, len(d.config.GPIO.CS))

	for i, n := range d.config.GPIO.CS {
		d.pins[i] = gpioreg.ByName(n)

		if d.pins[i] == nil {
			return errors.New("invalid pin name")
		}
	}

	return nil
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
