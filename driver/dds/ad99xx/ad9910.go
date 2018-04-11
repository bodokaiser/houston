package ad9910

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math"
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/spi"
	"periph.io/x/periph/conn/spi/spireg"

	"github.com/bodokaiser/beagle/register/dds/ad99xx"
)

// AD9910Config is the config for the AD9910 driver.
type AD9910Config struct {
	SysClock    float64
	RefClock    float64
	ResetPin    string
	IOUpdatePin string
	SPIDevice   string
	SPIMaxFreq  int64
	SPIMode     spi.Mode
}

// AD9910 drives the AD9910 dds.
type AD9910 struct {
	register    *ad99xx.AD9910
	refClock    float64
	sysClock    float64
	resetPin    gpio.PinIO
	ioUpdatePin gpio.PinIO
	spiConn     spi.Conn
}

// NewAD9910 returns an initialized AD9910 driver.
func NewAD9910(c AD9910Config) (*AD9910, error) {
	d := &AD9910{
		register:    ad99xx.NewAD9910(),
		refClock:    c.RefClock,
		sysClock:    c.SysClock,
		resetPin:    gpioreg.ByName(c.ResetPin),
		ioUpdatePin: gpioreg.ByName(c.IOUpdatePin),
	}

	if d.resetPin == nil {
		return nil, errors.New("failed to find reset GPIO pin")
	}
	if d.ioUpdatePin == nil {
		return nil, errors.New("failed to find I/O update GPIO pin")
	}

	spi, err := spireg.Open(c.SPIDevice)
	if err != nil {
		return nil, err
	}

	err = d.resetPin.Out(gpio.Low)
	if err != nil {
		return nil, err
	}
	err = d.ioUpdatePin.Out(gpio.Low)
	if err != nil {
		return nil, err
	}

	d.spiConn, err = spi.Connect(c.SPIMaxFreq, c.SPIMode, 8)

	return d, err
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

// SingleTone configures the AD9910 to single tone mode at given frequency.
func (d *AD9910) SingleTone(freq float64) error {
	ftw := ad99xx.FrequencyToFTW(d.sysClock, freq)

	d.register.CFR1[2] = ad99xx.FlagManualOSK
	d.register.CFR1[3] = ad99xx.FlagOSKEnable
	d.register.CFR1[4] = ad99xx.FlagSDIOInput

	d.register.CFR2[1] = ad99xx.FlagAmplScaleEnable
	d.register.CFR2[2] = ad99xx.FlagSYNCCLKEnable
	d.register.CFR2[3] = ad99xx.FlagPDCLKEnable
	d.register.CFR2[4] = ad99xx.FlagSyncValidDisable

	d.register.CFR3[1] = ad99xx.ModeDRV0OutputCurrentLow | ad99xx.ModeVCORange5
	d.register.CFR3[2] = ad99xx.ModeChargePumpCurrent387
	d.register.CFR3[3] = ad99xx.FlagREFCLKDivReset | ad99xx.FlagPLLEnable
	d.register.CFR3[4] = d.divider() << 1

	d.register.AuxDAC[4] = 0x7f

	d.register.ASF[3] = 0xff
	d.register.ASF[4] = 0xfc

	d.register.IOUpdateRate[1] = 0xff
	d.register.IOUpdateRate[2] = 0xff
	d.register.IOUpdateRate[3] = 0xff
	d.register.IOUpdateRate[4] = 0xff

	d.register.STProfile0[1] = 0x3f
	d.register.STProfile0[2] = 0xff
	binary.BigEndian.PutUint32(d.register.STProfile0[5:], ftw)

	w := bytes.Join([][]byte{
		d.register.CFR1[:],
		d.register.CFR2[:],
		d.register.CFR3[:],
		d.register.AuxDAC[:],
		d.register.ASF[:],
		d.register.IOUpdateRate[:],
		d.register.STProfile0[:],
	}, []byte{})
	r := make([]byte, len(w))

	err := d.spiConn.Tx(w, r)
	if err != nil {
		return err
	}

	return d.IOUpdate()
}

func (d *AD9910) divider() uint8 {
	return uint8(math.Round(d.sysClock / d.refClock))
}

func strobe(p gpio.PinIO) error {
	err := p.Out(gpio.High)
	if err != nil {
		return err
	}
	time.Sleep(time.Millisecond)

	return p.Out(gpio.Low)
}
