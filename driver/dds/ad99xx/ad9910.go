package ad99xx

import (
	"errors"
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/spi"
	"periph.io/x/periph/conn/spi/spireg"

	"github.com/bodokaiser/houston/device/dds/ad99xx/ad9910"
)

// AD9910 implements DDS interface for the AD9910.
type AD9910 struct {
	config      Config
	spiConn     spi.Conn
	resetPin    gpio.PinIO
	ioUpdatePin gpio.PinIO
	device      ad9910.AD9910
}

// NewAD9910 returns an initialized AD9910 driver.
func NewAD9910(c Config) *AD9910 {
	d := &AD9910{
		config:      c,
		resetPin:    gpioreg.ByName(c.ResetPin),
		ioUpdatePin: gpioreg.ByName(c.IOUpdatePin),
		device:      ad9910.NewAD9910(c.SysClock, c.RefClock),
	}

	return d
}

func (d *AD9910) Init() (err error) {
	if d.resetPin == nil {
		return errors.New("failed to find reset GPIO pin")
	}
	if d.ioUpdatePin == nil {
		return errors.New("failed to find I/O update GPIO pin")
	}

	spi, err := spireg.Open(d.config.SPIDevice)
	if err != nil {
		return
	}

	err = d.resetPin.Out(gpio.Low)
	if err != nil {
		return
	}
	err = d.ioUpdatePin.Out(gpio.Low)
	if err != nil {
		return
	}

	d.spiConn, err = spi.Connect(d.config.SPIMaxFreq, d.config.SPIMode, 8)

	return
}

func strobe(p gpio.PinIO) error {
	err := p.Out(gpio.High)
	if err != nil {
		return err
	}
	time.Sleep(time.Millisecond)

	return p.Out(gpio.Low)
}

// Reset triggers a reset which commands the connected DDS devices to clear
// the memory and reset the registers to the default values.
func (d *AD9910) Reset() error {
	return strobe(d.resetPin)
}

// IOUpdate triggers an I/O update which commands the connected DDS devices
// to read the updated configuration.
func (d *AD9910) IOUpdate() error {
	return strobe(d.ioUpdatePin)
}
