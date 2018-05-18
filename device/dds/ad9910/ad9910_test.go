package ad9910

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/bodokaiser/houston/device/dds"
	"github.com/bodokaiser/houston/register/dds/ad9910"
)

type AD9910TestSuite struct {
	suite.Suite

	d AD9910
}

func (s *AD9910TestSuite) SetupTest() {
	s.d = NewAD9910(dds.Config{
		SysClock: 1e9,
		RefClock: 1e7,
	})
}

func (s *AD9910TestSuite) TestRampClock() {
	assert.Equal(s.T(), 2.5e8, s.d.SysClock()/4)
}

func (s *AD9910TestSuite) TestRampParams() {
	step, rate := s.d.rampParams(1, 1.0)

	assert.Equal(s.T(), uint32(1021704), step)
	assert.Equal(s.T(), uint16(59471), rate)
}

func (s *AD9910TestSuite) TestAmplitude() {
	s.d.CFR1.SetRAMEnabled(true)
	s.d.CFR1.SetRAMDest(ad9910.RAMDestAmplitude)
	assert.Panics(s.T(), func() {
		s.d.Amplitude()
	})
	s.d.CFR1.SetRAMEnabled(false)

	s.d.CFR2.SetRampEnabled(true)
	s.d.CFR2.SetRampDest(ad9910.RampDestAmplitude)
	assert.Panics(s.T(), func() {
		s.d.Amplitude()
	})
	s.d.CFR2.SetRampEnabled(false)

	s.d.STProfile0.SetAmplScaleFactor(AmplitudeToASF(1.0))
	assert.Equal(s.T(), s.d.Amplitude(), 1.0)

	s.d.CFR1.SetRAMEnabled(true)
	s.d.CFR1.SetRAMDest(ad9910.RAMDestPhase)
	s.d.STProfile0.SetAmplScaleFactor(AmplitudeToASF(0.0))
	s.d.ASF.SetAmplScaleFactor(AmplitudeToASF(1.0))
	assert.Equal(s.T(), s.d.Amplitude(), 1.0)
}

func (s *AD9910TestSuite) TestFrequency() {
	s.d.CFR1.SetRAMEnabled(true)
	s.d.CFR1.SetRAMDest(ad9910.RAMDestFrequency)
	assert.Panics(s.T(), func() {
		s.d.Frequency()
	})
	s.d.CFR1.SetRAMEnabled(false)

	s.d.CFR2.SetRampEnabled(true)
	s.d.CFR2.SetRampDest(ad9910.RampDestFrequency)
	assert.Panics(s.T(), func() {
		s.d.Frequency()
	})
	s.d.CFR2.SetRampEnabled(false)

	s.d.STProfile0.SetFreqTuningWord(FrequencyToFTW(12e6, s.d.SysClock()))
	assert.Equal(s.T(), s.d.Frequency(), 12e6)

	s.d.CFR1.SetRAMEnabled(true)
	s.d.CFR1.SetRAMDest(ad9910.RAMDestPhase)
	s.d.FTW.SetFreqTuningWord(FrequencyToFTW(10e6, s.d.SysClock()))
	assert.Equal(s.T(), s.d.Frequency(), 10e6)
}

func (s *AD9910TestSuite) TestPhaseOffset() {
	s.d.CFR1.SetRAMEnabled(true)
	s.d.CFR1.SetRAMDest(ad9910.RAMDestPhase)
	assert.Panics(s.T(), func() {
		s.d.PhaseOffset()
	})
	s.d.CFR1.SetRAMEnabled(false)

	s.d.CFR2.SetRampEnabled(true)
	s.d.CFR2.SetRampDest(ad9910.RampDestPhase)
	assert.Panics(s.T(), func() {
		s.d.PhaseOffset()
	})
	s.d.CFR2.SetRampEnabled(false)

	s.d.STProfile0.SetPhaseOffsetWord(PhaseToPOW(1.23))
	assert.InEpsilon(s.T(), s.d.PhaseOffset(), 1.23, 0.001)

	s.d.CFR1.SetRAMEnabled(true)
	s.d.CFR1.SetRAMDest(ad9910.RAMDestAmplitude)
	s.d.POW.SetPhaseOffsetWord(PhaseToPOW(1.79))
	assert.InEpsilon(s.T(), s.d.PhaseOffset(), 1.79, 0.001)
}

func (s *AD9910TestSuite) TestSetAmplitude() {
	s.d.SetAmplitude(0.742)
}

func (s *AD9910TestSuite) TestSetFrequency() {

}

func (s *AD9910TestSuite) TestSetPhaseOffset() {

}

func TestAD9910Suite(t *testing.T) {
	suite.Run(t, new(AD9910TestSuite))
}

func TestAmplitudeToASF(t *testing.T) {
	assert.Panics(t, func() {
		AmplitudeToASF(-0.1)
	})
	assert.Panics(t, func() {
		AmplitudeToASF(+1.1)
	})

	assert.Equal(t, uint16(math.MaxUint16>>2), AmplitudeToASF(1.0))
	assert.Equal(t, uint16(0), AmplitudeToASF(0.0))
}

func TestASFToAmplitude(t *testing.T) {
	assert.Equal(t, 1.0, ASFToAmplitude(math.MaxUint16>>2))
	assert.Equal(t, 0.0, ASFToAmplitude(0.0))
}

func TestFrequencyToFTW(t *testing.T) {
	assert.Panics(t, func() {
		FrequencyToFTW(0, 1e9)
	})
	assert.Panics(t, func() {
		FrequencyToFTW(421e6, 1e9)
	})

	assert.Equal(t, uint32(0x556aaaab), FrequencyToFTW(41e6, 122880e3))
}

func TestFTWToFrequency(t *testing.T) {
	assert.Equal(t, 41e6, FTWToFrequency(0x556aaaab, 122880e3))
}

func TestPhaseToPOW(t *testing.T) {
	assert.Panics(t, func() {
		PhaseToPOW(-0.01)
	})
	assert.Panics(t, func() {
		PhaseToPOW(2.1 * math.Pi)
	})

	assert.Equal(t, uint16(0), PhaseToPOW(0))
	assert.Equal(t, uint16(math.MaxUint16), PhaseToPOW(2*math.Pi))
}

func TestPOWToPhase(t *testing.T) {
	assert.Equal(t, 0.0, POWToPhase(0))
	assert.Equal(t, 2*math.Pi, POWToPhase(math.MaxUint16))
}
