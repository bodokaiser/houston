package ad99xx

import (
	"bytes"
	"errors"
	"math"
	"time"

	"github.com/bodokaiser/houston/register/dds/ad99xx/ad9910"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/spi"
	"periph.io/x/periph/conn/spi/spireg"
)

// AD9910 implements DDS interface for the AD9910.
type AD9910 struct {
	config      Config
	spiConn     spi.Conn
	resetPin    gpio.PinIO
	ioUpdatePin gpio.PinIO
	cfr1        ad9910.CFR1
	cfr2        ad9910.CFR2
	cfr3        ad9910.CFR3
	ftw         ad9910.FTW
	pow         ad9910.POW
	asf         ad9910.ASF
	rampLimit   ad9910.RampLimit
	rampStep    ad9910.RampStep
	rampRate    ad9910.RampRate
	stProfile0  ad9910.STProfile
	ramProfile0 ad9910.RAMProfile
}

// NewAD9910 returns an initialized AD9910 driver.
func NewAD9910(c Config) *AD9910 {
	d := &AD9910{
		config:      c,
		resetPin:    gpioreg.ByName(c.ResetPin),
		ioUpdatePin: gpioreg.ByName(c.IOUpdatePin),
		cfr1:        ad9910.NewCFR1(),
		cfr2:        ad9910.NewCFR2(),
		cfr3:        ad9910.NewCFR3(),
		ftw:         ad9910.NewFTW(),
		pow:         ad9910.NewPOW(),
		asf:         ad9910.NewASF(),
		rampLimit:   ad9910.NewRampLimit(),
		rampStep:    ad9910.NewRampStep(),
		rampRate:    ad9910.NewRampRate(),
		stProfile0:  ad9910.NewSTProfile(),
		ramProfile0: ad9910.NewRAMProfile(),
	}

	d.cfr1.SetSDIOInputOnly(true)
	d.cfr2.SetSTAmplScaleEnabled(true)
	d.cfr2.SetSyncClockEnabled(true)
	d.cfr2.SetSyncTimingValidationDisabled(true)
	d.cfr3.SetVCORange(ad9910.VCORange5)
	d.cfr3.SetPLLEnabled(true)
	d.cfr3.SetDivider(uint(math.Round(c.SysClock / c.RefClock)))

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

func asfToAmpl(x uint16) float64 {
	return float64(x) / (1<<14 - 1)
}

func (d *AD9910) Amplitude() float64 {
	asf := d.asf.AmplScaleFactor()

	if d.cfr2.STAmplScaleEnabled() && !d.cfr1.RAMEnabled() {
		asf = d.stProfile0.AmplScaleFactor()
	}

	return asfToAmpl(asf)
}

func amplToASF(x float64) uint16 {
	return uint16(math.Round(x * (1<<14 - 1)))
}

func (d *AD9910) SetAmplitude(x float64) {
	if x < 0 || x > 1 {
		panic("amplitude is not in range between 0 and 1")
	}
	asf := amplToASF(x)

	if !d.cfr1.RAMEnabled() {
		d.stProfile0.SetAmplScaleFactor(asf)
	}
	d.asf.SetAmplScaleFactor(asf)
}

func ftwToFreq(x uint32, y float64) float64 {
	return float64(x) * y / (1<<32 - 1)
}

func (d *AD9910) Frequency() float64 {
	if d.cfr1.RAMEnabled() && d.cfr1.RAMDest() == ad9910.RAMDestFrequency {
		panic("frequency is controlled by RAM")
	}
	if d.cfr2.RampEnabled() && d.cfr2.RampDest() == ad9910.RampDestFrequency {
		panic("frequency is controlled by ramp")
	}
	// parallal data port controls frequency

	if d.cfr1.RAMEnabled() {
		return ftwToFreq(d.ftw.FreqTuningWord(), d.config.SysClock)
	}
	return ftwToFreq(d.stProfile0.FreqTuningWord(), d.config.SysClock)
}

func freqToFTW(x float64, y float64) uint32 {
	return uint32(math.Round(x / y * (1<<32 - 1)))
}

func (d *AD9910) SetFrequency(f float64) {
	if f < 0 || f > 420e6 {
		panic("frequency is not in range between 0 and 420 MHz")
	}
	ftw := freqToFTW(f, d.config.SysClock)

	if !d.cfr1.RAMEnabled() {
		d.stProfile0.SetFreqTuningWord(ftw)
	}
	d.ftw.SetFreqTuningWord(ftw)
}

func powToPhase(x uint16) float64 {
	return float64(x) * (2 * math.Pi) / (1<<16 - 1)
}

func (d *AD9910) PhaseOffset() float64 {
	if d.cfr1.RAMEnabled() && (d.cfr1.RAMDest() == ad9910.RAMDestPhase ||
		d.cfr1.RAMDest() == ad9910.RAMDestPolar) {
		panic("phase is controlled by RAM")
	}
	if d.cfr2.RampEnabled() && d.cfr2.RampDest() == ad9910.RampDestPhase {
		panic("phase is controlled by ramp")
	}
	// parallal data port controls phase

	if d.cfr1.RAMEnabled() {
		return powToPhase(d.pow.PhaseOffsetWord())
	}

	return powToPhase(d.stProfile0.PhaseOffsetWord())
}

func phaseToPOW(x float64) uint16 {
	return uint16(math.Round(x / (2 * math.Pi) * (1<<16 - 1)))
}

func (d *AD9910) SetPhaseOffset(p float64) {
	pow := phaseToPOW(math.Mod(p, 2*math.Pi))

	if !d.cfr1.RAMEnabled() {
		d.stProfile0.SetPhaseOffsetWord(pow)
	}
	d.pow.SetPhaseOffsetWord(pow)
}

func writeInstr(x byte, y []byte) []byte {
	return append([]byte{x}, y...)
}

func (d *AD9910) toBytes() []byte {
	b := [][]byte{
		writeInstr(ad9910.AddrCFR1, []byte(d.cfr1)),
		writeInstr(ad9910.AddrCFR2, []byte(d.cfr2)),
		writeInstr(ad9910.AddrCFR3, []byte(d.cfr3)),
		writeInstr(ad9910.AddrFTW, []byte(d.ftw)),
		writeInstr(ad9910.AddrPOW, []byte(d.pow)),
		writeInstr(ad9910.AddrASF, []byte(d.asf)),
	}

	if d.cfr2.RampEnabled() {
		b = append(b, [][]byte{
			writeInstr(ad9910.AddrRampLimit, []byte(d.rampLimit)),
			writeInstr(ad9910.AddrRampStep, []byte(d.rampStep)),
			writeInstr(ad9910.AddrRampRate, []byte(d.rampRate)),
		}...)
	}

	if d.cfr1.RAMEnabled() {
		b = append(b, [][]byte{
			writeInstr(ad9910.AddrProfile0, []byte(d.ramProfile0)),
		}...)
	}
	b = append(b, [][]byte{
		writeInstr(ad9910.AddrProfile0, []byte(d.stProfile0)),
	}...)

	return bytes.Join(b, []byte{})
}

func (d *AD9910) Exec() error {
	w := d.toBytes()
	r := make([]byte, len(w))

	err := d.spiConn.Tx(w, r)
	if err != nil {
		return err
	}

	return d.IOUpdate()
}
