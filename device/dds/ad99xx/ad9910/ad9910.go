package ad9910

import (
	"bytes"
	"log"
	"math"
	"time"

	"github.com/bodokaiser/houston/device/dds"
	"github.com/bodokaiser/houston/register/dds/ad99xx/ad9910"
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

type AD9910 struct {
	config       dds.Config
	cfr1         ad9910.CFR1
	cfr2         ad9910.CFR2
	cfr3         ad9910.CFR3
	auxDAC       ad9910.AuxDAC
	ioUpdateRate ad9910.IOUpdateRate
	ftw          ad9910.FTW
	pow          ad9910.POW
	asf          ad9910.ASF
	rampLimit    ad9910.RampLimit
	rampStep     ad9910.RampStep
	rampRate     ad9910.RampRate
	stProfile0   ad9910.STProfile
	stProfile1   ad9910.STProfile
	stProfile2   ad9910.STProfile
	stProfile3   ad9910.STProfile
	stProfile4   ad9910.STProfile
	stProfile5   ad9910.STProfile
	stProfile6   ad9910.STProfile
	stProfile7   ad9910.STProfile
	ramProfile0  ad9910.RAMProfile
	ramProfile1  ad9910.RAMProfile
	ramProfile2  ad9910.RAMProfile
	ramProfile3  ad9910.RAMProfile
	ramProfile4  ad9910.RAMProfile
	ramProfile5  ad9910.RAMProfile
	ramProfile6  ad9910.RAMProfile
	ramProfile7  ad9910.RAMProfile
}

func NewAD9910(c dds.Config) AD9910 {
	d := AD9910{
		config:       c,
		cfr1:         ad9910.NewCFR1(),
		cfr2:         ad9910.NewCFR2(),
		cfr3:         ad9910.NewCFR3(),
		auxDAC:       ad9910.NewAuxDAC(),
		ioUpdateRate: ad9910.NewIOUpdateRate(),
		ftw:          ad9910.NewFTW(),
		pow:          ad9910.NewPOW(),
		asf:          ad9910.NewASF(),
		rampLimit:    ad9910.NewRampLimit(),
		rampStep:     ad9910.NewRampStep(),
		rampRate:     ad9910.NewRampRate(),
		stProfile0:   ad9910.NewSTProfile(),
		stProfile1:   ad9910.NewSTProfile(),
		stProfile2:   ad9910.NewSTProfile(),
		stProfile3:   ad9910.NewSTProfile(),
		stProfile4:   ad9910.NewSTProfile(),
		stProfile5:   ad9910.NewSTProfile(),
		stProfile6:   ad9910.NewSTProfile(),
		stProfile7:   ad9910.NewSTProfile(),
		ramProfile0:  ad9910.NewRAMProfile(),
		ramProfile1:  ad9910.NewRAMProfile(),
		ramProfile2:  ad9910.NewRAMProfile(),
		ramProfile3:  ad9910.NewRAMProfile(),
		ramProfile4:  ad9910.NewRAMProfile(),
		ramProfile5:  ad9910.NewRAMProfile(),
		ramProfile6:  ad9910.NewRAMProfile(),
		ramProfile7:  ad9910.NewRAMProfile(),
	}

	d.cfr1.SetSDIOInputOnly(true)
	d.cfr2.SetSTAmplScaleEnabled(true)
	d.cfr2.SetSyncClockEnabled(true)
	d.cfr2.SetSyncTimingValidationDisabled(true)
	d.cfr3.SetVCORange(ad9910.VCORange5)
	d.cfr3.SetPLLEnabled(true)
	d.cfr3.SetDivider(divider(d.config.SysClock, d.config.RefClock))

	return d
}

func divider(x, y float64) uint8 {
	return uint8(math.Round(x / y))
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

func assertAmpl(x float64) {
	if x < 0 || x > 1 {
		panic("amplitude is not in range between 0 and 1")
	}
}

func (d *AD9910) SetAmplitude(x float64) {
	assertAmpl(x)
	asf := amplToASF(x)

	if !d.cfr1.RAMEnabled() {
		d.stProfile0.SetAmplScaleFactor(asf)
	}
	d.asf.SetAmplScaleFactor(asf)
}

func ftwToFreq(x uint32, y float64) float64 {
	return float64(x) * y / (1 << 32)
}

func (d *AD9910) ftwToFreq(x uint32) float64 {
	return ftwToFreq(x, float64(d.config.SysClock))
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
		return d.ftwToFreq(d.ftw.FreqTuningWord())
	}
	return d.ftwToFreq(d.stProfile0.FreqTuningWord())
}

func assertFreq(x float64) {
	if x < 0 || x > 420e6 {
		panic("frequency is not in range between 0 and 420 MHz")
	}
}

func freqToFTW(out float64, sys float64) uint32 {
	return uint32(math.Round((1 << 32) * (out / sys)))
}

func (d *AD9910) freqToFTW(f float64) uint32 {
	return freqToFTW(f, float64(d.config.SysClock))
}

func (d *AD9910) SetFrequency(f float64) {
	log.Printf("set frequency: %v\n", f)
	assertFreq(f)
	ftw := d.freqToFTW(f)
	log.Printf("set ftw: %v\n", ftw)

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

func assertPhase(x float64) {
	if x < 0 || x > 2*math.Pi {
		panic("phase not in range between 0 and 2 pi")
	}
}

func phaseToPOW(x float64) uint16 {
	return uint16(math.Round(x / (2 * math.Pi) * (1<<16 - 1)))
}

func (d *AD9910) SetPhaseOffset(p float64) {
	assertPhase(p)
	pow := phaseToPOW(p)

	if !d.cfr1.RAMEnabled() {
		d.stProfile0.SetPhaseOffsetWord(pow)
	}
	d.pow.SetPhaseOffsetWord(pow)
}

func assertRange(a, b float64) {
	if a >= b {
		panic("lower limit not greater than upper limit")
	}
}

func sweepRate(d time.Duration, sysClock, step float64) uint16 {
	return uint16(math.Round(d.Seconds() * sysClock / (4 * step)))
}

func (d *AD9910) Sweep(c dds.SweepConfig) {
	a, b := c.Limits[0], c.Limits[1]
	assertRange(a, b)

	var l, u uint32

	switch c.Param {
	case dds.ParamAmplitude:
		d.cfr2.SetRampDest(ad9910.RampDestAmplitude)

		assertAmpl(a)
		assertAmpl(b)

		//l = amplToASF(a)
		//u = amplToASF(b)
	case dds.ParamFrequency:
		d.cfr2.SetRampDest(ad9910.RampDestFrequency)

		assertFreq(a)
		assertFreq(b)

		l = d.freqToFTW(a)
		u = d.freqToFTW(b)

		d.rampLimit.SetLowerLimit(l)
		d.rampLimit.SetUpperLimit(u)
	case dds.ParamPhase:
		d.cfr2.SetRampDest(ad9910.RampDestPhase)
	default:
		panic("invalid parameter")
	}

	r := sweepRate(c.Duration, float64(d.config.SysClock), float64(u-l))

	d.cfr2.SetRampEnabled(true)
	d.cfr2.SetRampNoDwellLow(c.NoDwells[0])
	d.cfr2.SetRampNoDwellHigh(c.NoDwells[1])

	d.rampStep.SetDecrStepSize(1)
	d.rampStep.SetIncrStepSize(1)
	d.rampRate.SetNegSlopeRate(r)
	d.rampRate.SetPosSlopeRate(r)
}

func (d *AD9910) Playback(c dds.PlaybackConfig) {

}

func prefix(prefix byte, b []byte) []byte {
	return append([]byte{prefix}, b[:]...)
}

func (d *AD9910) Encode() ([]byte, error) {
	log.Printf("encode:\n%+v\n", d)

	// TODO: only send necessary registers
	p := [][]byte{
		prefix(addrCFR1, d.cfr1[:]),
		prefix(addrCFR2, d.cfr2[:]),
		prefix(addrCFR3, d.cfr3[:]),
		prefix(addrAuxDAC, d.auxDAC[:]),
		prefix(addrIOUpdate, d.ioUpdateRate[:]),
	}

	if d.ftw.FreqTuningWord() > 0 {
		p = append(p, prefix(addrFTW, d.ftw[:]))
	}
	if d.pow.PhaseOffsetWord() > 0 {
		p = append(p, prefix(addrPOW, d.pow[:]))
	}
	if d.asf.AmplScaleFactor() > 0 {
		p = append(p, prefix(addrASF, d.asf[:]))
	}

	if d.cfr1.RAMEnabled() {
		p = append(p, prefix(addrProfile0, d.ramProfile0[:]))
	} else {
		p = append(p, prefix(addrProfile0, d.stProfile0[:]))
	}
	if d.cfr2.RampEnabled() {
		p = append(p, prefix(addrRampLimit, d.rampLimit[:]))
		p = append(p, prefix(addrRampStep, d.rampStep[:]))
		p = append(p, prefix(addrRampRate, d.rampRate[:]))
	}

	b := bytes.Join(p, []byte{})

	log.Printf("encode:\n%+v\n", b)

	return b, nil
}

func (d *AD9910) Decode(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	l := 0

	switch b[0] & ^byte(flagRead) {
	case addrCFR1:
		l = len(d.cfr1)
		copy(d.cfr1[:], b[1:l])
	case addrCFR2:
		l = len(d.cfr2)
		copy(d.cfr2[:], b[1:l])
	case addrCFR3:
		l = len(d.cfr3)
		copy(d.cfr3[:], b[1:l])
	}

	return d.Decode(b[1+l:])
}
