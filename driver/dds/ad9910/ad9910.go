package ad9910

import (
	"bytes"
	"errors"
	"time"

	"periph.io/x/periph/conn"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpiotest"
	"periph.io/x/periph/conn/physic"
	"periph.io/x/periph/conn/spi"
	"periph.io/x/periph/conn/spi/spitest"

	"github.com/bodokaiser/houston/device/dds/ad9910"
	"github.com/bodokaiser/houston/driver/dds"
)

const (
	flagRead  = 0x80
	flagWrite = 0x00
)

const (
	addrCFR1      = 0x00
	addrCFR2      = 0x01
	addrCFR3      = 0x02
	addrAuxDAC    = 0x03
	addrIOUpdate  = 0x04
	addrFTW       = 0x07
	addrPOW       = 0x08
	addrASF       = 0x09
	addrRampLimit = 0x0b
	addrRampStep  = 0x0c
	addrRampRate  = 0x0d
	addrProfile0  = 0x0e
	addrProfile1  = 0x0f
	addrProfile2  = 0x10
	addrProfile3  = 0x11
	addrProfile4  = 0x12
	addrProfile5  = 0x13
	addrProfile6  = 0x14
	addrProfile7  = 0x15
	addrRAM       = 0x16
)

// AD9910 implements DDS interface for the AD9910.
type AD9910 struct {
	ad9910.AD9910
	config    dds.Config
	spiConn   spi.Conn
	spiPort   spi.Port
	resetPin  gpio.PinIO
	updatePin gpio.PinIO
}

// NewAD9910 returns an initialized AD9910 driver.
func NewAD9910(c dds.Config) *AD9910 {
	d := &AD9910{
		config:    c,
		AD9910:    ad9910.NewAD9910(c.Config),
		spiPort:   c.SPIPort,
		resetPin:  c.ResetPin,
		updatePin: c.UpdatePin,
	}

	return d
}

// Init implements driver.Driver interface.
func (d *AD9910) Init() (err error) {
	if d.spiPort == nil {
		return errors.New("failed to find SPI port")
	}
	if d.resetPin == nil {
		return errors.New("failed to find reset GPIO pin")
	}
	if d.updatePin == nil {
		return errors.New("failed to find update GPIO pin")
	}

	if err = d.resetPin.Out(gpio.Low); err != nil {
		return
	}
	if err = d.updatePin.Out(gpio.Low); err != nil {
		return
	}

	d.spiConn, err = d.spiPort.Connect(5*physic.MegaHertz, spi.Mode0, 8)
	if err != nil {
		return
	}

	if d.Debug() {
		d.resetPin = &gpiotest.LogPinIO{PinIO: d.resetPin}
		d.updatePin = &gpiotest.LogPinIO{PinIO: d.updatePin}
		d.spiConn = &spitest.LogConn{Conn: d.spiConn}
	}

	return
}

func strobe(p gpio.PinIO) error {
	if err := p.Out(gpio.High); err != nil {
		return err
	}
	time.Sleep(time.Millisecond)

	return p.Out(gpio.Low)
}

// Reset implements DDS interace.
func (d *AD9910) Reset() error {
	return strobe(d.resetPin)
}

// Update implements DDS interace.
func (d *AD9910) Update() error {
	return strobe(d.updatePin)
}

func prefix(prefix byte, b []byte) []byte {
	return append([]byte{prefix}, b[:]...)
}

// MaxTxSize returns the maximum number of bytes per SPI
// transaction.
func (d *AD9910) MaxTxSize() int {
	if c, ok := d.spiConn.(*spitest.LogConn); ok {
		if l, ok := c.Conn.(conn.Limits); ok {
			return l.MaxTxSize()
		}
	}
	if l, ok := d.spiConn.(conn.Limits); ok {
		return l.MaxTxSize()
	}

	return d.config.Config.SPIMaxTxSize
}

// Exec implements DDS interace.
func (d *AD9910) Exec() error {
	d.CFR1.SetOSKEnabled(false)

	if d.CFR1.RAMEnabled() {
		d.CFR2.SetSTAmplScaleEnabled(false)
	}

	p := [][]byte{
		prefix(addrCFR1, d.CFR1[:]),
		prefix(addrCFR2, d.CFR2[:]),
		prefix(addrCFR3, d.CFR3[:]),
		prefix(addrAuxDAC, d.AuxDAC[:]),
		prefix(addrIOUpdate, d.IOUpdateRate[:]),
	}

	if d.FTW.FreqTuningWord() > 0 {
		p = append(p, prefix(addrFTW, d.FTW[:]))
	}
	if d.POW.PhaseOffsetWord() > 0 {
		p = append(p, prefix(addrPOW, d.POW[:]))
	}
	if d.ASF.AmplScaleFactor() > 0 {
		p = append(p, prefix(addrASF, d.ASF[:]))
	}

	if d.CFR1.RAMEnabled() && len(d.RAM) > 0 {
		p = append(p, prefix(addrProfile0, d.RAMProfile0[:]))
	} else {
		p = append(p, prefix(addrProfile0, d.STProfile0[:]))
	}
	if d.CFR2.RampEnabled() {
		p = append(p, prefix(addrRampLimit, d.RampLimit[:]))
		p = append(p, prefix(addrRampStep, d.RampStep[:]))
		p = append(p, prefix(addrRampRate, d.RampRate[:]))
	}

	if err := d.send(bytes.Join(p, []byte{})); err != nil {
		return err
	}
	if err := d.Update(); err != nil {
		return err
	}

	if d.CFR1.RAMEnabled() && len(d.RAM) > 0 {
		w := []byte{addrRAM}

		for _, b := range d.RAM {
			w = append(w, b[:]...)
		}

		if err := d.send(w); err != nil {
			return err
		}
	}

	return d.Update()
}

func (d *AD9910) send(b []byte) error {
	max := d.MaxTxSize()

	for len(b) != 0 {
		if max > len(b) {
			max = len(b)
		}

		if err := d.spiConn.Tx(b[:max], nil); err != nil {
			return err
		}

		b = b[max:]
	}

	return nil
}
