package ad99xx

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/spi"
	"periph.io/x/periph/conn/spi/spireg"

	"github.com/bodokaiser/houston/driver/dds"
	"github.com/bodokaiser/houston/register/dds/ad99xx"
)

// AD9910 implements DDS interface for the AD9910.
type AD9910 struct {
	register    *ad99xx.AD9910
	refClock    float64
	sysClock    float64
	resetPin    gpio.PinIO
	ioUpdatePin gpio.PinIO
	spiConn     spi.Conn
}

// NewAD9910 returns an initialized AD9910 driver.
func NewAD9910(c Config) (*AD9910, error) {
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

// SingleTone implements DDS interface for AD9910.
func (d *AD9910) SingleTone(c dds.SingleToneConfig) error {
	if c.Amplitude < 0 || c.Amplitude > 1 {
		return errors.New("amplitude not between 0.0 and 1.0")
	}
	if c.Frequency < 1 || c.Frequency > 500e6 {
		return errors.New("frequency not between 1 Hz and 500 MHz")
	}
	if c.PhaseOffset < 0 || c.PhaseOffset > 2*math.Pi {
		return errors.New("phase offset not between 0 and 2π")
	}

	d.register.CFR1[2] = ad99xx.FlagManualOSK
	d.register.CFR1[3] = ad99xx.FlagOSKEnable
	d.register.CFR1[4] = ad99xx.FlagSDIOInput

	d.register.CFR2[1] = ad99xx.FlagAmplScaleEnable
	d.register.CFR2[2] = ad99xx.FlagSYNCCLKEnable
	d.register.CFR2[3] = ad99xx.FlagPDCLKEnable
	d.register.CFR2[4] = ad99xx.FlagSyncValidDisable

	// TODO: ModeVCORangeX should be inferred from SysClock.
	d.register.CFR3[1] = ad99xx.ModeDRV0OutputCurrentLow | ad99xx.ModeVCORange5
	d.register.CFR3[2] = ad99xx.ModeChargePumpCurrent387
	d.register.CFR3[3] = ad99xx.FlagREFCLKDivReset | ad99xx.FlagPLLEnable
	d.register.CFR3[4] = d.divider() << 1

	pow := ad99xx.PhaseToPOW(c.PhaseOffset)
	asf := ad99xx.AmplitudeToASF(c.Amplitude)
	ftw := ad99xx.FrequencyToFTW(d.sysClock, c.Frequency)

	// for some reason ASF register has to be written and cannot be omitted
	// as done with FTW, POW
	binary.BigEndian.PutUint16(d.register.ASF[3:], asf<<2)
	binary.BigEndian.PutUint16(d.register.STProfile0[1:3], asf)
	binary.BigEndian.PutUint16(d.register.STProfile0[3:5], pow)
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

// DigitalRamp implements DDS interface.
func (d *AD9910) DigitalRamp(c dds.DigitalRampConfig) error {
	d.register.CFR1[2] = ad99xx.FlagManualOSK
	d.register.CFR1[3] = ad99xx.FlagOSKEnable
	d.register.CFR1[4] = ad99xx.FlagSDIOInput

	d.register.CFR2[2] = ad99xx.FlagSYNCCLKEnable
	d.register.CFR2[3] = ad99xx.FlagPDCLKEnable
	d.register.CFR2[4] = ad99xx.FlagSyncValidDisable

	// TODO: ModeVCORangeX should be inferred from SysClock.
	d.register.CFR3[1] = ad99xx.ModeDRV0OutputCurrentLow | ad99xx.ModeVCORange5
	d.register.CFR3[2] = ad99xx.ModeChargePumpCurrent387
	d.register.CFR3[3] = ad99xx.FlagREFCLKDivReset | ad99xx.FlagPLLEnable
	d.register.CFR3[4] = d.divider() << 1

	d.register.CFR2[2] |= ad99xx.FlagDRampEnable
	if c.NoDwell[0] {
		d.register.CFR2[2] |= ad99xx.FlagDRampNoDwellLow
	}
	if c.NoDwell[1] {
		d.register.CFR2[2] |= ad99xx.FlagDRampNoDwellHigh
	}
	d.register.CFR2[2] |= byte(c.Destination << 4)

	pow := ad99xx.PhaseToPOW(c.PhaseOffset)
	asf := ad99xx.AmplitudeToASF(c.Amplitude)
	ftw := ad99xx.FrequencyToFTW(d.sysClock, c.Frequency)

	binary.BigEndian.PutUint16(d.register.ASF[3:], asf<<2)
	binary.BigEndian.PutUint32(d.register.FTW[1:], ftw)
	binary.BigEndian.PutUint16(d.register.POW[1:], pow)

	switch c.Destination {
	case dds.Amplitude:
		u := ad99xx.AmplitudeToASF(c.Limits[1])
		l := ad99xx.AmplitudeToASF(c.Limits[0])
		s := uint16(math.Round(c.Duration.Seconds() * d.sysClock / (4 * float64(u-l))))

		binary.BigEndian.PutUint32(d.register.DRampLimit[1:5], uint32(u)<<18)
		binary.BigEndian.PutUint32(d.register.DRampLimit[5:], uint32(l)<<18)
		binary.BigEndian.PutUint16(d.register.DRampRate[1:3], s)
		binary.BigEndian.PutUint16(d.register.DRampRate[3:], s)
	case dds.Frequency:
		u := ad99xx.FrequencyToFTW(d.sysClock, c.Limits[1])
		l := ad99xx.FrequencyToFTW(d.sysClock, c.Limits[0])
		s := uint16(math.Round(c.Duration.Seconds() * d.sysClock / (4 * float64(u-l))))

		binary.BigEndian.PutUint32(d.register.DRampLimit[1:5], u)
		binary.BigEndian.PutUint32(d.register.DRampLimit[5:], l)
		binary.BigEndian.PutUint16(d.register.DRampRate[1:3], s)
		binary.BigEndian.PutUint16(d.register.DRampRate[3:], s)
	case dds.PhaseOffset:
		u := ad99xx.PhaseToPOW(c.Limits[1])
		l := ad99xx.PhaseToPOW(c.Limits[0])
		s := uint16(math.Round(c.Duration.Seconds() * d.sysClock / (4 * float64(u-l))))

		binary.BigEndian.PutUint32(d.register.DRampLimit[1:5], uint32(u)<<16)
		binary.BigEndian.PutUint32(d.register.DRampLimit[5:], uint32(l)<<16)
		binary.BigEndian.PutUint16(d.register.DRampRate[1:3], s)
		binary.BigEndian.PutUint16(d.register.DRampRate[3:], s)
	}
	// the smallest possible step size gives us the highest resolution
	binary.BigEndian.PutUint32(d.register.DRampStepSize[1:5], uint32(1))
	binary.BigEndian.PutUint32(d.register.DRampStepSize[5:], uint32(1))

	fmt.Printf("%+v\n", d.register)

	w := bytes.Join([][]byte{
		d.register.CFR1[:],
		d.register.CFR2[:],
		d.register.CFR3[:],
		d.register.AuxDAC[:],
		d.register.IOUpdateRate[:],
		d.register.FTW[:],
		d.register.POW[:],
		d.register.ASF[:],
		d.register.DRampLimit[:],
		d.register.DRampStepSize[:],
		d.register.DRampRate[:],
	}, []byte{})
	r := make([]byte, len(w))

	err := d.spiConn.Tx(w, r)
	if err != nil {
		return err
	}

	return d.IOUpdate()
}

// Playback implements DDS interace.
func (d *AD9910) Playback(c dds.PlaybackConfig) error {
	return nil
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
