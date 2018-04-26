package ad9910

import (
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

func (s *AD9910TestSuite) TestFreqToFTW() {
	fout := 41e6     // 41 MHz
	fsys := 122880e3 // 122.88 MHz

	assert.Equal(s.T(), uint32(0x556aaaab), freqToFTW(fout, fsys))
}

func (s *AD9910TestSuite) TestRampClock() {
	assert.Equal(s.T(), 2.5e8, s.d.SysClock()/4)
}

func (s *AD9910TestSuite) TestRampParams() {
	step, rate := s.d.rampParams(1, 1.0)

	assert.Equal(s.T(), uint32(1021704), step)
	assert.Equal(s.T(), uint16(59471), rate)
}

func TestAD9910Suite(t *testing.T) {
	suite.Run(t, new(AD9910TestSuite))
}
