package ad9910

import (
	"bytes"
	"errors"
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/spi"
	"periph.io/x/periph/conn/spi/spireg"
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

	spiDev, err := spireg.Open(d.config.SPI.Device)
	if err != nil {
		return
	}

	if err = d.resetPin.Out(gpio.Low); err != nil {
		return
	}
	if err = d.updatePin.Out(gpio.Low); err != nil {
		return
	}

	d.spiConn, err = spiDev.Connect(d.config.SPI.MaxFreq, spi.Mode0, 8)
	if err != nil {
		return
	}

	if d.Debug() {
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

func (d *AD9910) Reset() error {
	return strobe(d.resetPin)
}

func (d *AD9910) Update() error {
	return strobe(d.updatePin)
}

func prefix(prefix byte, b []byte) []byte {
	return append([]byte{prefix}, b[:]...)
}

func (d *AD9910) Exec() error {
	// we cannot write to RAM if RAM is enabled
	re := d.CFR1.RAMEnabled()
	d.CFR1.SetOSKEnabled(false)

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

	if re && len(d.RAM) > 0 {
		p = append(p, prefix(addrProfile0, d.RAMProfile0[:]))
	} else {
		p = append(p, prefix(addrProfile0, d.STProfile0[:]))
	}
	if d.CFR2.RampEnabled() {
		p = append(p, prefix(addrRampLimit, d.RampLimit[:]))
		p = append(p, prefix(addrRampStep, d.RampStep[:]))
		p = append(p, prefix(addrRampRate, d.RampRate[:]))
	}

	w := bytes.Join(p, []byte{})
	r := make([]byte, len(w))

	if err := d.spiConn.Tx(w, r); err != nil {
		return err
	}
	if err := d.Update(); err != nil {
		return err
	}

	if re && len(d.RAM) > 0 {
		w = []byte{addrRAM}
		for _, ram := range d.RAM {
			w = append(w, ram[:]...)
		}
		r = make([]byte, len(w))

		if err := d.spiConn.Tx(w, r); err != nil {
			return err
		}

		d.CFR1.SetRAMEnabled(true)
		w = prefix(addrCFR1, d.CFR1[:])
		r = make([]byte, len(w))
		if err := d.spiConn.Tx(w, r); err != nil {
			return err
		}
	}

	return d.Update()
}