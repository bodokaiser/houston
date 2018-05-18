package ad9910

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/bodokaiser/houston/device/dds"
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

func (s *AD9910TestSuite) TestAmplitudeToASF() {
	assert.Panics(s.T(), func() {
		AmplitudeToASF(-0.1)
	})
	assert.Panics(s.T(), func() {
		AmplitudeToASF(+1.1)
	})

	assert.Equal(s.T(), uint16(math.MaxUint16>>2), AmplitudeToASF(1.0))
	assert.Equal(s.T(), uint16(0), AmplitudeToASF(0.0))
}

func (s *AD9910TestSuite) TestASFToAmplitude() {
	assert.Equal(s.T(), 1.0, ASFToAmplitude(math.MaxUint16>>2))
	assert.Equal(s.T(), 0.0, ASFToAmplitude(0.0))
}

func (s *AD9910TestSuite) TestFrequencyToFTW() {
	assert.Panics(s.T(), func() {
		FrequencyToFTW(0, 1e9)
	})
	assert.Panics(s.T(), func() {
		FrequencyToFTW(421e6, 1e9)
	})

	assert.Equal(s.T(), uint32(0x556aaaab), FrequencyToFTW(41e6, 122880e3))
}

func (s *AD9910TestSuite) TestFTWToFrequency() {
	assert.Equal(s.T(), 41e6, FTWToFrequency(0x556aaaab, 122880e3))
}

func (s *AD9910TestSuite) TestPhaseToPOW() {
	assert.Panics(s.T(), func() {
		PhaseToPOW(-0.01)
	})
	assert.Panics(s.T(), func() {
		PhaseToPOW(2.1 * math.Pi)
	})

	assert.Equal(s.T(), uint16(0), PhaseToPOW(0))
	assert.Equal(s.T(), uint16(math.MaxUint16), PhaseToPOW(2*math.Pi))
}

func (s *AD9910TestSuite) TestPOWToPhase() {
	assert.Equal(s.T(), 0.0, POWToPhase(0))
	assert.Equal(s.T(), 2*math.Pi, POWToPhase(math.MaxUint16))
}

func (s *AD9910TestSuite) TestRampClock() {
	assert.Equal(s.T(), 2.5e8, s.d.SysClock()/4)
}

func (s *AD9910TestSuite) TestRampParams() {
	step, rate := s.d.rampParams(1, 1.0)

	assert.Equal(s.T(), uint32(1021704), step)
	assert.Equal(s.T(), uint16(59471), rate)
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
