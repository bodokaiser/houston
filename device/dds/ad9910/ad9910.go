package ad9910

import (
	"fmt"
	"math"

	"github.com/bodokaiser/approx"
	"github.com/bodokaiser/houston/device/dds"
	"github.com/bodokaiser/houston/register/dds/ad9910"
)

type AD9910 struct {
	config       dds.Config
	CFR1         ad9910.CFR1
	CFR2         ad9910.CFR2
	CFR3         ad9910.CFR3
	AuxDAC       ad9910.AuxDAC
	IOUpdateRate ad9910.IOUpdateRate
	FTW          ad9910.FTW
	POW          ad9910.POW
	ASF          ad9910.ASF
	RampLimit    ad9910.RampLimit
	RampStep     ad9910.RampStep
	RampRate     ad9910.RampRate
	STProfile0   ad9910.STProfile
	STProfile1   ad9910.STProfile
	STProfile2   ad9910.STProfile
	STProfile3   ad9910.STProfile
	STProfile4   ad9910.STProfile
	STProfile5   ad9910.STProfile
	STProfile6   ad9910.STProfile
	STProfile7   ad9910.STProfile
	RAMProfile0  ad9910.RAMProfile
	RAMProfile1  ad9910.RAMProfile
	RAMProfile2  ad9910.RAMProfile
	RAMProfile3  ad9910.RAMProfile
	RAMProfile4  ad9910.RAMProfile
	RAMProfile5  ad9910.RAMProfile
	RAMProfile6  ad9910.RAMProfile
	RAMProfile7  ad9910.RAMProfile
	RAM          []ad9910.RAM
}

func NewAD9910(c dds.Config) AD9910 {
	d := AD9910{
		config:       c,
		CFR1:         ad9910.NewCFR1(),
		CFR2:         ad9910.NewCFR2(),
		CFR3:         ad9910.NewCFR3(),
		AuxDAC:       ad9910.NewAuxDAC(),
		IOUpdateRate: ad9910.NewIOUpdateRate(),
		FTW:          ad9910.NewFTW(),
		POW:          ad9910.NewPOW(),
		ASF:          ad9910.NewASF(),
		RampLimit:    ad9910.NewRampLimit(),
		RampStep:     ad9910.NewRampStep(),
		RampRate:     ad9910.NewRampRate(),
		STProfile0:   ad9910.NewSTProfile(),
		STProfile1:   ad9910.NewSTProfile(),
		STProfile2:   ad9910.NewSTProfile(),
		STProfile3:   ad9910.NewSTProfile(),
		STProfile4:   ad9910.NewSTProfile(),
		STProfile5:   ad9910.NewSTProfile(),
		STProfile6:   ad9910.NewSTProfile(),
		STProfile7:   ad9910.NewSTProfile(),
		RAMProfile0:  ad9910.NewRAMProfile(),
		RAMProfile1:  ad9910.NewRAMProfile(),
		RAMProfile2:  ad9910.NewRAMProfile(),
		RAMProfile3:  ad9910.NewRAMProfile(),
		RAMProfile4:  ad9910.NewRAMProfile(),
		RAMProfile5:  ad9910.NewRAMProfile(),
		RAMProfile6:  ad9910.NewRAMProfile(),
		RAMProfile7:  ad9910.NewRAMProfile(),
	}

	d.CFR1.SetSDIOInputOnly(true)
	d.CFR2.SetSTAmplScaleEnabled(true)
	d.CFR2.SetSyncClockEnabled(true)
	d.CFR2.SetSyncTimingValidationDisabled(true)

	if c.PLL && c.RefClock > 0 {
		div := divider(d.config.SysClock, d.config.RefClock)

		d.CFR3.SetPLLEnabled(true)
		d.CFR3.SetDivider(div)

		switch {
		case c.SysClock >= 400e6 && c.SysClock < 460e6:
			d.CFR3.SetVCORange(ad9910.VCORange0)
		case c.SysClock >= 455e6 && c.SysClock < 530e6:
			d.CFR3.SetVCORange(ad9910.VCORange1)
		case c.SysClock >= 530e6 && c.SysClock < 615e6:
			d.CFR3.SetVCORange(ad9910.VCORange2)
		case c.SysClock >= 650e6 && c.SysClock < 790e6:
			d.CFR3.SetVCORange(ad9910.VCORange3)
		case c.SysClock >= 760e6 && c.SysClock < 875e6:
			d.CFR3.SetVCORange(ad9910.VCORange4)
		case c.SysClock >= 920e6 && c.SysClock < 1030e6:
			d.CFR3.SetVCORange(ad9910.VCORange5)
		default:
			panic("sys clock with PLL out of VCO ranges")
		}
	}

	return d
}

func divider(x, y float64) uint8 {
	return uint8(math.Round(x / y))
}

func (d *AD9910) SysClock() float64 {
	return d.config.SysClock
}

func (d *AD9910) RefClock() float64 {
	return d.config.SysClock
}

func asfToAmpl(x uint16) float64 {
	return float64(x) / (1<<14 - 1)
}

func (d *AD9910) Amplitude() float64 {
	asf := d.ASF.AmplScaleFactor()

	if d.CFR2.STAmplScaleEnabled() && !d.CFR1.RAMEnabled() {
		asf = d.STProfile0.AmplScaleFactor()
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

	if !d.CFR1.RAMEnabled() {
		d.STProfile0.SetAmplScaleFactor(asf)
	}
	d.ASF.SetAmplScaleFactor(asf)
}

func ftwToFreq(x uint32, y float64) float64 {
	return float64(x) * y / (1 << 32)
}

func (d *AD9910) ftwToFreq(x uint32) float64 {
	return ftwToFreq(x, float64(d.SysClock()))
}

func (d *AD9910) Frequency() float64 {
	if d.CFR1.RAMEnabled() && d.CFR1.RAMDest() == ad9910.RAMDestFrequency {
		panic("frequency is controlled by RAM")
	}
	if d.CFR2.RampEnabled() && d.CFR2.RampDest() == ad9910.RampDestFrequency {
		panic("frequency is controlled by ramp")
	}
	// parallal data port controls frequency

	if d.CFR1.RAMEnabled() {
		return d.ftwToFreq(d.FTW.FreqTuningWord())
	}
	return d.ftwToFreq(d.STProfile0.FreqTuningWord())
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
	return freqToFTW(f, float64(d.SysClock()))
}

func (d *AD9910) SetFrequency(f float64) {
	assertFreq(f)
	ftw := d.freqToFTW(f)

	if !d.CFR1.RAMEnabled() {
		d.STProfile0.SetFreqTuningWord(ftw)
	}
	d.FTW.SetFreqTuningWord(ftw)
}

func powToPhase(x uint16) float64 {
	return float64(x) * (2 * math.Pi) / (1<<16 - 1)
}

func (d *AD9910) PhaseOffset() float64 {
	if d.CFR1.RAMEnabled() && (d.CFR1.RAMDest() == ad9910.RAMDestPhase ||
		d.CFR1.RAMDest() == ad9910.RAMDestPolar) {
		panic("phase is controlled by RAM")
	}
	if d.CFR2.RampEnabled() && d.CFR2.RampDest() == ad9910.RampDestPhase {
		panic("phase is controlled by ramp")
	}
	// parallal data port controls phase

	if d.CFR1.RAMEnabled() {
		return powToPhase(d.POW.PhaseOffsetWord())
	}

	return powToPhase(d.STProfile0.PhaseOffsetWord())
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

	if !d.CFR1.RAMEnabled() {
		d.STProfile0.SetPhaseOffsetWord(pow)
	}
	d.POW.SetPhaseOffsetWord(pow)
}

func assertRange(a, b float64) {
	if a >= b {
		panic("lower limit not greater than upper limit")
	}
}

const maxRate = (1 << 16) - 1

func (d *AD9910) rampClock() float64 {
	return d.config.SysClock / 4
}

func (d *AD9910) rampParams(T, dx float64) (uint32, uint16) {
	m := T * d.rampClock() / (math.MaxUint32 * dx)
	r, s := approx.RatioConstr2(m, math.MaxUint32, math.MaxUint16)

	return uint32(s), uint16(r)
}

func (d *AD9910) Sweep(c dds.SweepConfig) {
	a, b := c.Limits[0], c.Limits[1]
	assertRange(a, b)

	scale := 1.0

	switch c.Param {
	case dds.ParamAmplitude:
		assertAmpl(a)
		assertAmpl(b)

		d.CFR2.SetRampDest(ad9910.RampDestAmplitude)
		d.RampLimit.SetLowerASF(amplToASF(a))
		d.RampLimit.SetUpperASF(amplToASF(b))
	case dds.ParamFrequency:
		assertFreq(a)
		assertFreq(b)
		scale = d.SysClock()

		d.CFR2.SetRampDest(ad9910.RampDestFrequency)
		d.RampLimit.SetLowerFTW(d.freqToFTW(a))
		d.RampLimit.SetUpperFTW(d.freqToFTW(b))
	case dds.ParamPhase:
		assertPhase(a)
		assertPhase(b)
		scale = 2 * math.Pi

		d.CFR2.SetRampDest(ad9910.RampDestPhase)
		d.RampLimit.SetLowerASF(phaseToPOW(a))
		d.RampLimit.SetUpperASF(phaseToPOW(b))
	default:
		panic("invalid parameter")
	}
	s, r := d.rampParams(c.Duration.Seconds(),
		float64(b-a)/scale)

	d.CFR2.SetRampEnabled(true)
	d.CFR2.SetRampNoDwellLow(c.NoDwells[0])
	d.CFR2.SetRampNoDwellHigh(c.NoDwells[1])

	d.RampStep.SetDecrStepSize(s)
	d.RampStep.SetIncrStepSize(s)
	d.RampRate.SetNegSlopeRate(r)
	d.RampRate.SetPosSlopeRate(r)
}

func (d *AD9910) playbackClock() float64 {
	return d.SysClock() / 4
}

func (d *AD9910) playbackParams(T float64) uint16 {
	return uint16(math.Round(T * d.playbackClock()))
}

func (d *AD9910) Playback(c dds.PlaybackConfig) {
	t := c.Duration.Seconds()
	tmin := 1 / d.playbackClock()
	tmax := tmin * math.MaxUint16

	if t < tmin || t > tmax {
		panic(fmt.Sprintf("interval %v out of range (%v, %v)", t, tmin, tmax))
	}
	l := uint16(len(c.Data)) - 1
	r := d.playbackParams(t)

	d.RAM = []ad9910.RAM{}
	d.CFR1.SetRAMEnabled(true)
	d.RAMProfile0.SetNoDwellHigh(true)
	d.RAMProfile0.SetAddrStepRate(r)
	d.RAMProfile0.SetWaveformStartAddr(0)
	d.RAMProfile0.SetWaveformEndAddr(l)

	switch c.Param {
	case dds.ParamAmplitude:
		d.CFR1.SetRAMDest(ad9910.RAMDestAmplitude)

		for _, v := range c.Data {
			assertAmpl(v)

			r := ad9910.NewRAM()
			r.SetAmplScaleFactor(amplToASF(v))

			d.RAM = append(d.RAM, r)
		}
	case dds.ParamFrequency:
		d.CFR1.SetRAMDest(ad9910.RAMDestFrequency)

		for _, v := range c.Data {
			assertFreq(v)

			r := ad9910.NewRAM()
			r.SetFreqTuningWord(d.freqToFTW(v))

			d.RAM = append(d.RAM, r)
		}
	case dds.ParamPhase:
		d.CFR1.SetRAMDest(ad9910.RAMDestPhase)

		for _, v := range c.Data {
			assertPhase(v)

			r := ad9910.NewRAM()
			r.SetPhaseOffsetWord(phaseToPOW(v))

			d.RAM = append(d.RAM, r)
		}
	}

	if c.Trigger && !c.Duplex {
		d.RAMProfile0.SetRAMControlMode(ad9910.RAMControlModeRampUp)
	}
	if c.Trigger && c.Duplex {
		d.RAMProfile0.SetRAMControlMode(ad9910.RAMControlModeBiRampUp)
	}
	if !c.Trigger && c.Duplex {
		d.RAMProfile0.SetRAMControlMode(ad9910.RAMControlModeContBiRampUp)
	}
	if !c.Trigger && !c.Duplex {
		d.RAMProfile0.SetRAMControlMode(ad9910.RAMControlModeContRecirculate)
	}
}
