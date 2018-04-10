package dds

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math"
	"time"

	"github.com/bodokaiser/beagle/device/dds/ad99xx"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/spi"
	"periph.io/x/periph/conn/spi/spireg"
)

const (
	intIOUpdatePinName = "27"
	extIOUpdatePinName = "46"
	intResetPinName    = "65"
	extResetPinName    = "47"
)

const spiDevName = "SPI1.0"

const pulseWidth = time.Millisecond

// Default clock values.
const (
	DefaultSysClock = 1e9
	DefaultRefClock = 1e7
)

// AD9910 images a driver to the AD9910.
type AD9910 struct {
	device         *ad99xx.AD9910
	sysClock       float64
	refClock       float64
	conn           spi.Conn
	extIOUpdatePin gpio.PinIO
	intIOUpdatePin gpio.PinIO
	extResetPin    gpio.PinIO
	intResetPin    gpio.PinIO
}

// NewAD9910 returns an initialized AD9910 driver.
func NewAD9910(sys float64, ref float64) (*AD9910, error) {
	d := &AD9910{
		sysClock:       sys,
		refClock:       ref,
		device:         ad99xx.NewAD9910(),
		extResetPin:    gpioreg.ByName(extResetPinName),
		intResetPin:    gpioreg.ByName(intResetPinName),
		extIOUpdatePin: gpioreg.ByName(extIOUpdatePinName),
		intIOUpdatePin: gpioreg.ByName(intIOUpdatePinName),
	}

	if d.extIOUpdatePin == nil {
		return nil, errors.New("failed to find GPIO pin for external I/O update")
	}
	if d.extResetPin == nil {
		return nil, errors.New("failed to find GPIO pin for external reset")
	}
	if d.intIOUpdatePin == nil {
		return nil, errors.New("failed to find GPIO pin for internal I/O update")
	}
	if d.intResetPin == nil {
		return nil, errors.New("failed to find GPIO pin for internal reset")
	}

	err := d.extIOUpdatePin.Out(gpio.High)
	if err != nil {
		return nil, err
	}

	err = d.extResetPin.Out(gpio.High)
	if err != nil {
		return nil, err
	}

	err = d.intIOUpdatePin.Out(gpio.Low)
	if err != nil {
		return nil, err
	}

	spi, err := spireg.Open(spiDevName)
	if err != nil {
		return nil, err
	}

	d.conn, err = spi.Connect(5e6, 0, 8)

	return d, err
}

// Reset triggers a reset which commands the connected DDS devices to clear
// the memory and reset the registers to the default values.
func (d *AD9910) Reset() error {
	return strobe(d.intResetPin, pulseWidth)
}

// IOUpdate triggers an I/O update which commands the connected DDS devices
// to read the updated configuration.
func (d *AD9910) IOUpdate() error {
	return strobe(d.intIOUpdatePin, pulseWidth)
}

// SingleTone configures the AD9910 to single tone mode at given frequency.
func (d *AD9910) SingleTone(freq float64) error {
	ftw := ad99xx.FrequencyToFTW(d.sysClock, freq)

	d.device.CFR1[2] = ad99xx.FlagManualOSK
	d.device.CFR1[3] = ad99xx.FlagOSKEnable
	d.device.CFR1[4] = ad99xx.FlagSDIOInput

	d.device.CFR2[1] = ad99xx.FlagAmplScaleEnable
	d.device.CFR2[2] = ad99xx.FlagSYNCCLKEnable
	d.device.CFR2[3] = ad99xx.FlagPDCLKEnable
	d.device.CFR2[4] = ad99xx.FlagSyncValidDisable

	d.device.CFR3[1] = ad99xx.ModeDRV0OutputCurrentLow | ad99xx.ModeVCORange5
	d.device.CFR3[2] = ad99xx.ModeChargePumpCurrent387
	d.device.CFR3[3] = ad99xx.FlagREFCLKDivReset | ad99xx.FlagPLLEnable
	d.device.CFR3[4] = d.divider() << 1

	d.device.AuxDAC[4] = 0x7f

	d.device.ASF[3] = 0xff
	d.device.ASF[4] = 0xfc

	d.device.IOUpdateRate[1] = 0xff
	d.device.IOUpdateRate[2] = 0xff
	d.device.IOUpdateRate[3] = 0xff
	d.device.IOUpdateRate[4] = 0xff

	d.device.STProfile0[1] = 0x3f
	d.device.STProfile0[2] = 0xff
	binary.BigEndian.PutUint32(d.device.STProfile0[5:], ftw)

	w := bytes.Join([][]byte{
		d.device.CFR1[:],
		d.device.CFR2[:],
		d.device.CFR3[:],
		d.device.AuxDAC[:],
		d.device.ASF[:],
		d.device.IOUpdateRate[:],
		d.device.STProfile0[:],
	}, []byte{})
	r := make([]byte, len(w))

	err := d.conn.Tx(w, r)
	if err != nil {
		return err
	}

	return d.IOUpdate()
}

func (d *AD9910) divider() uint8 {
	return uint8(math.Round(d.sysClock / d.refClock))
}

func strobe(p gpio.PinIO, d time.Duration) error {
	err := p.Out(gpio.High)
	if err != nil {
		return err
	}
	time.Sleep(d)

	return p.Out(gpio.Low)
}
