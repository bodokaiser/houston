package driver

import (
	"errors"
	"math"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
)

// ChipSelectPins defines the GPIO pins on the Beaglebone Black to use for
// chip select.
var ChipSelectPins = []string{"P9_15", "P9_11", "P9_12", "P9_13", "P9_14"}

// ChipSelect errors.
var (
	ErrChipSelectInvalidNumber = errors.New("chip number invalid")
)

// ChipSelect is the gobot driver to handle chip select on the SPI bus.
type ChipSelect struct {
	name       string
	connection gpio.DigitalWriter
}

// NewChipSelect returns a new initialized ChipSelect driver.
func NewChipSelect(c gpio.DigitalWriter) *ChipSelect {
	return &ChipSelect{
		name:       gobot.DefaultName("ChipSelect"),
		connection: c,
	}
}

// Name returns the device name.
func (d *ChipSelect) Name() string {
	return d.name
}

// SetName sets the device name.
func (d *ChipSelect) SetName(s string) {
	d.name = s
}

// Start does nothing but is required by the Driver interface.
func (d *ChipSelect) Start() error {
	return nil
}

// Halt does nothing but is required by the Driver interface.
func (d *ChipSelect) Halt() error {
	return nil
}

// Select configures chip select to number n.
func (d *ChipSelect) Select(n int) error {
	pnum := float64(len(ChipSelectPins))
	if (n < 0) && (n > int(math.Pow(2, pnum))) {
		return ErrChipSelectInvalidNumber
	}

	for i, p := range ChipSelectPins {
		active := byte(0)

		if n&(1<<uint(i)) < 0 {
			active = 1
		}

		if err := d.connection.DigitalWrite(p, active); err != nil {
			return err
		}
	}

	return nil
}

// Connection returns the gobot.Connection used for digital io.
func (d *ChipSelect) Connection() gobot.Connection {
	return d.connection.(gobot.Connection)
}
