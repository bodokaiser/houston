package ad99xx

import (
	"errors"
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/spi"
	"periph.io/x/periph/conn/spi/spireg"

	"github.com/bodokaiser/houston/device/dds/ad99xx/ad9910"
	"github.com/bodokaiser/houston/driver/dds"
)

// AD9910 implements DDS interface for the AD9910.
type AD9910 struct {
	ad9910.AD9910
	config    dds.Config
	spiConn   spi.Conn
	resetPin  gpio.PinIO
	updatePin gpio.PinIO
}

// NewAD9910 returns an initialized AD9910 driver.
func NewAD9910(c dds.Config) *AD9910 {
	d := &AD9910{
		config: c,
		AD9910: ad9910.NewAD9910(c.Config),
	}

	return d
}

func (d *AD9910) Init() (err error) {
	d.resetPin = gpioreg.ByName(d.config.GPIO.Reset)
	d.updatePin = gpioreg.ByName(d.config.GPIO.Update)

	if d.resetPin == nil {
		return errors.New("failed to find reset GPIO pin")
	}
	if d.updatePin == nil {
		return errors.New("failed to find I/O update GPIO pin")
	}

	spi, err := spireg.Open(d.config.SPI.Device)
	if err != nil {
		return
	}

	err = d.resetPin.Out(gpio.Low)
	if err != nil {
		return
	}
	err = d.updatePin.Out(gpio.Low)
	if err != nil {
		return
	}

	d.spiConn, err = spi.Connect(d.config.SPI.MaxFreq, d.config.SPI.Mode, 8)

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

func (d *AD9910) Reset() error {
	return strobe(d.resetPin)
}

func (d *AD9910) Update() error {
	return strobe(d.updatePin)
}

func (d *AD9910) WriteToDev() error {
	w, err := d.Encode()
	if err != nil {
		return err
	}
	r := make([]byte, len(w))

	return d.spiConn.Tx(w, r)
}

func (d *AD9910) ReadFromDev() error {
	return nil
}
