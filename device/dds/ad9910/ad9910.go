// Package ad9910 provides a structure to mirror the AD9910 DDS chip.
package ad9910

import (
	"fmt"
	"math"

	"github.com/bodokaiser/approx"
	"github.com/bodokaiser/houston/device/dds"
	"github.com/bodokaiser/houston/register/dds/ad9910"
)

// AD9910 mirrors a AD9910 DDS device as described by the datasheet.
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

// NewAD9910 returns an initialized AD9910 from config.
//
// In particular this will initialize present registers with defaults defined
// in the config.
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

	if c.SPI3Wire {
		d.CFR1.SetSDIOInputOnly(true)
	}

	d.CFR2.SetSTAmplScaleEnabled(true)
	d.CFR2.SetSyncTimingValidationDisabled(true)

	if c.PLL && c.RefClock > 0 {
		div := divider(d.config.SysClock, d.config.RefClock)

		d.CFR3.SetPLLEnabled(true)
		d.CFR2.SetSyncClockEnabled(true)
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

// Debug returns true if debug mode is enabled.
func (d *AD9910) Debug() bool {
	return d.config.Debug
}

// SysClock returns the system clock frequency.
func (d *AD9910) SysClock() float64 {
	return d.config.SysClock
}

// RefClock returns the reference clock frequency.
func (d *AD9910) RefClock() float64 {
	return d.config.SysClock
}

// AmplitudeToASF converts a relative amplitude scale from 0 to 1 to an ASF word.
//
// If the amplitude scales resides outside of the range 0 to 1 this function will panic.
func AmplitudeToASF(amplitude float64) uint16 {
	if amplitude < 0 || amplitude > 1 {
		panic("amplitude is not in range between 0 and 1")
	}

	return uint16(math.Round(amplitude * (math.MaxUint16 >> 2)))
}

// ASFToAmplitude converts an ASF word to a relative amplitude scale from 0 to 1.
func ASFToAmplitude(asf uint16) float64 {
	return float64(asf) / (math.MaxUint16 >> 2)
}

// Amplitude returns the relative amplitude.
//
// Depending on the device configuration the amplitude will originate from
// the ASF or STProfile0 register even if sweep or playback is configured for
// the amplitude.
func (d *AD9910) Amplitude() float64 {
	asf := d.ASF.AmplScaleFactor()

	if d.CFR1.RAMEnabled() && d.CFR1.RAMDest() == ad9910.RAMDestAmplitude {
		panic("amplitude controlled by RAM")
	}
	if d.CFR2.RampEnabled() && d.CFR2.RampDest() == ad9910.RampDestAmplitude {
		panic("amplitude controlled by ramp")
	}

	if d.CFR2.STAmplScaleEnabled() && !d.CFR1.RAMEnabled() {
		asf = d.STProfile0.AmplScaleFactor()
	}

	return ASFToAmplitude(asf)
}

// SetAmplitude sets the relative amplitude.
//
// See Amplitude() for details.
func (d *AD9910) SetAmplitude(x float64) {
	asf := AmplitudeToASF(x)

	if !d.CFR1.RAMEnabled() {
		d.STProfile0.SetAmplScaleFactor(asf)
	}
	d.ASF.SetAmplScaleFactor(asf)
}

// FrequencyToFTW converts an output frequency in Hz to a FTW word.
//
// If the frequency resides outside of the range 0 to 420 MHz this function will panic.
func FrequencyToFTW(frequency float64, sysClock float64) uint32 {
	if frequency <= 0 || frequency > 420e6 {
		panic("frequency is not in range between 0 and 420 MHz")
	}

	return uint32(math.Round(math.MaxUint32*(frequency/sysClock))) + 1
}

// FTWToFrequency converts a FTW word to a output frequency in Hz.
func FTWToFrequency(ftw uint32, sysClock float64) float64 {
	return math.Round(float64(ftw) * sysClock / math.MaxUint32)
}

// Frequency returns the output frequency in Hz.
//
// Depending on the device configuration this method will panic if frequency
// is controlled by RAM or the digital ramp.
func (d *AD9910) Frequency() float64 {
	if d.CFR1.RAMEnabled() && d.CFR1.RAMDest() == ad9910.RAMDestFrequency {
		panic("frequency is controlled by RAM")
	}
	if d.CFR2.RampEnabled() && d.CFR2.RampDest() == ad9910.RampDestFrequency {
		panic("frequency is controlled by ramp")
	}
	// parallal data port controls frequency

	if d.CFR1.RAMEnabled() {
		return FTWToFrequency(d.FTW.FreqTuningWord(), d.SysClock())
	}
	return FTWToFrequency(d.STProfile0.FreqTuningWord(), d.SysClock())
}

// SetFrequency sets the output frequency.
//
// See Frequency() for details.
func (d *AD9910) SetFrequency(x float64) {
	ftw := FrequencyToFTW(x, d.SysClock())

	if !d.CFR1.RAMEnabled() {
		d.STProfile0.SetFreqTuningWord(ftw)
	}
	d.FTW.SetFreqTuningWord(ftw)
}

// PhaseToPOW converts a phase in radiants to a POW word.
//
// If the phase resides outside of the range 0 to 2π this function will panic.
func PhaseToPOW(phase float64) uint16 {
	if phase < 0 || phase > 2*math.Pi {
		panic("phase not in range between 0 and 2π")
	}

	return uint16(math.Round(phase / (2 * math.Pi) * math.MaxUint16))
}

// POWToPhase converts a POW word to a phase in radiants.
func POWToPhase(pow uint16) float64 {
	return float64(pow) * (2 * math.Pi) / math.MaxUint16
}

// PhaseOffset returns the phase offset in rads.
//
// Depending on the device configuration this method will panic if phase
// is controlled by RAM or the digital ramp.
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
		return POWToPhase(d.POW.PhaseOffsetWord())
	}

	return POWToPhase(d.STProfile0.PhaseOffsetWord())
}

// SetPhaseOffset sets the output phase offset.
//
// See PhaseOffset() for details.
func (d *AD9910) SetPhaseOffset(x float64) {
	pow := PhaseToPOW(x)

	if !d.CFR1.RAMEnabled() {
		d.STProfile0.SetPhaseOffsetWord(pow)
	}
	d.POW.SetPhaseOffsetWord(pow)
}

// RampClock returns the digital ramp clock in Hz.
func (d *AD9910) RampClock() float64 {
	return d.SysClock() / 4
}

func (d *AD9910) rampParams(T, dx float64) (uint32, uint16) {
	m := T * d.RampClock() / (math.MaxUint32 * dx)
	r, s := approx.RatioConstr2(m, math.MaxUint32, math.MaxUint16)

	return uint32(s), uint16(r)
}

// Sweep returns the configured sweep as SweepConfig.
//
// Will panic as not implemented.
func (d *AD9910) Sweep() dds.SweepConfig {
	panic("not implemented")
}

// SetSweep configures a sweep as defined in SweepConfig.
//
// Will panic on invalid sweep parameters.
func (d *AD9910) SetSweep(c dds.SweepConfig) {
	a, b := c.Limits[0], c.Limits[1]
	if a >= b {
		panic("lower limit not greater than upper limit")
	}

	scale := 1.0

	switch c.Param {
	case dds.ParamAmplitude:
		d.CFR2.SetRampDest(ad9910.RampDestAmplitude)
		d.RampLimit.SetLowerASF(AmplitudeToASF(a))
		d.RampLimit.SetUpperASF(AmplitudeToASF(b))
	case dds.ParamFrequency:
		scale = d.SysClock()

		d.CFR2.SetRampDest(ad9910.RampDestFrequency)
		d.RampLimit.SetLowerFTW(FrequencyToFTW(a, scale))
		d.RampLimit.SetUpperFTW(FrequencyToFTW(b, scale))
	case dds.ParamPhase:
		scale = 2 * math.Pi

		d.CFR2.SetRampDest(ad9910.RampDestPhase)
		d.RampLimit.SetLowerASF(PhaseToPOW(a))
		d.RampLimit.SetUpperASF(PhaseToPOW(b))
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

// Playback returns the configured playback as PlaybackConfig.
//
// Will panic as not implemented.
func (d *AD9910) Playback() dds.PlaybackConfig {
	panic("not implemented")
}

// SetPlayback configures a playback as defined in PlaybackConfig.
//
// Will panic on invalid playback parameters.
func (d *AD9910) SetPlayback(c dds.PlaybackConfig) {
	t := c.Interval.Seconds()
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
			r := ad9910.NewRAM()
			r.SetAmplScaleFactor(AmplitudeToASF(v))

			d.RAM = append(d.RAM, r)
		}
	case dds.ParamFrequency:
		d.CFR1.SetRAMDest(ad9910.RAMDestFrequency)

		for _, v := range c.Data {
			r := ad9910.NewRAM()
			r.SetFreqTuningWord(FrequencyToFTW(v, d.SysClock()))

			d.RAM = append(d.RAM, r)
		}
	case dds.ParamPhase:
		d.CFR1.SetRAMDest(ad9910.RAMDestPhase)

		for _, v := range c.Data {
			r := ad9910.NewRAM()
			r.SetPhaseOffsetWord(PhaseToPOW(v))

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
