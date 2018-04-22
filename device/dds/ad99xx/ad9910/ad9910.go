package ad9910

import (
	"bytes"
	"math"

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
	SysClock     float64
	RefClock     float64
	cfr1         ad9910.CFR1
	cfr2         ad9910.CFR2
	cfr3         ad9910.CFR3
	auxDAC       ad9910.AuxDAC
	ioUpdataRate ad9910.IOUpdateRate
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

func NewAD9910(sysClock, refClock uint32) AD9910 {
	d := AD9910{
		SysClock:    float64(sysClock),
		RefClock:    float64(refClock),
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
		stProfile1:  ad9910.NewSTProfile(),
		stProfile2:  ad9910.NewSTProfile(),
		stProfile3:  ad9910.NewSTProfile(),
		stProfile4:  ad9910.NewSTProfile(),
		stProfile5:  ad9910.NewSTProfile(),
		stProfile6:  ad9910.NewSTProfile(),
		stProfile7:  ad9910.NewSTProfile(),
		ramProfile0: ad9910.NewRAMProfile(),
		ramProfile1: ad9910.NewRAMProfile(),
		ramProfile2: ad9910.NewRAMProfile(),
		ramProfile3: ad9910.NewRAMProfile(),
		ramProfile4: ad9910.NewRAMProfile(),
		ramProfile5: ad9910.NewRAMProfile(),
		ramProfile6: ad9910.NewRAMProfile(),
		ramProfile7: ad9910.NewRAMProfile(),
	}

	d.cfr1.SetSDIOInputOnly(true)
	d.cfr2.SetSTAmplScaleEnabled(true)
	d.cfr2.SetSyncClockEnabled(true)
	d.cfr2.SetSyncTimingValidationDisabled(true)
	d.cfr3.SetVCORange(ad9910.VCORange5)
	d.cfr3.SetPLLEnabled(true)
	d.cfr3.SetDivider(uint(math.Round(d.SysClock / d.RefClock)))

	return d
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
		return ftwToFreq(d.ftw.FreqTuningWord(), d.SysClock)
	}
	return ftwToFreq(d.stProfile0.FreqTuningWord(), d.SysClock)
}

func freqToFTW(x float64, y float64) uint32 {
	return uint32(math.Round(x / y * (1<<32 - 1)))
}

func (d *AD9910) SetFrequency(f float64) {
	if f < 0 || f > 420e6 {
		panic("frequency is not in range between 0 and 420 MHz")
	}
	ftw := freqToFTW(f, d.SysClock)

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

func (d *AD9910) Encode() ([]byte, error) {
	// TODO: only send necessary registers
	b := bytes.Join([][]byte{
		append([]byte{addrCFR1}, d.cfr1[:]...),
		append([]byte{addrCFR2}, d.cfr2[:]...),
		append([]byte{addrCFR3}, d.cfr3[:]...),
		append([]byte{addrAuxDAC}, d.auxDAC[:]...),
		append([]byte{addrIOUpdate}, d.ioUpdataRate[:]...),
		append([]byte{addrFTW}, d.ftw[:]...),
		append([]byte{addrASF}, d.asf[:]...),
		append([]byte{addrPOW}, d.pow[:]...),
		append([]byte{addrRampLimit}, d.rampLimit[:]...),
		append([]byte{addrRampStep}, d.rampStep[:]...),
		append([]byte{addrRampRate}, d.rampRate[:]...),
	}, nil)

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
