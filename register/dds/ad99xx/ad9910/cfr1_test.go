package ad9910

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CFR1TestSuite struct {
	suite.Suite

	r CFR1
}

func (s *CFR1TestSuite) SetupTest() {
	s.r = NewCFR1()
}

func (s *CFR1TestSuite) TestRAMEnabled() {
	s.r[0] = 0
	assert.False(s.T(), s.r.RAMEnabled())

	s.r[0] = 1 << 7
	assert.True(s.T(), s.r.RAMEnabled())
}

func (s *CFR1TestSuite) TestSetRamEnabled() {
	s.r[0] = 0

	s.r.SetRAMEnabled(true)
	assert.EqualValues(s.T(), 0x80, s.r[0])

	s.r.SetRAMEnabled(false)
	assert.False(s.T(), s.r[0]&1<<7 > 0)
}

func (s *CFR1TestSuite) TestRAMDest() {
	s.r[0] = 0x00
	assert.Equal(s.T(), RAMDestFrequency, s.r.RAMDest())

	s.r[0] = 0x20
	assert.Equal(s.T(), RAMDestPhase, s.r.RAMDest())

	s.r[0] = 0x40
	assert.Equal(s.T(), RAMDestAmplitude, s.r.RAMDest())

	s.r[0] = 0x60
	assert.Equal(s.T(), RAMDestPolar, s.r.RAMDest())
}

func (s *CFR1TestSuite) TestSetRAMDest() {
	s.r[0] = 0x00
	s.r.SetRAMDest(RAMDestFrequency)
	assert.EqualValues(s.T(), 0x00, s.r[0])
}

func (s *CFR1TestSuite) TestOSKEnabled() {
	s.r[2] = 0
	assert.False(s.T(), s.r.OSKEnabled())

	s.r[2] = 1 << 1
	assert.True(s.T(), s.r.OSKEnabled())
}

func (s *CFR1TestSuite) TestSetOSKEnabled() {
	s.r[2] = 0

	s.r.SetOSKEnabled(true)
	assert.EqualValues(s.T(), 0x2, s.r[2])

	s.r.SetOSKEnabled(false)
	assert.EqualValues(s.T(), 0x00, s.r[2])
}

func (s *CFR1TestSuite) TestOSKManual() {
	s.r[1] = 0
	assert.False(s.T(), s.r.OSKManual())

	s.r[1] = 1 << 7
	assert.True(s.T(), s.r.OSKManual())
}

func (s *CFR1TestSuite) TestSetOSKManual() {
	s.r[1] = 0

	s.r.SetOSKManual(true)
	assert.EqualValues(s.T(), 0x80, s.r[1])

	s.r.SetOSKManual(false)
	assert.EqualValues(s.T(), 0x00, s.r[1])
}

func (s *CFR1TestSuite) TestOSKAuto() {
	s.r[2] = 0x00
	assert.False(s.T(), s.r.OSKAuto())

	s.r[2] = 0x01
	assert.True(s.T(), s.r.OSKAuto())
}

func (s *CFR1TestSuite) TestSetOSKAuto() {
	s.r[2] = 0

	s.r.SetOSKAuto(true)
	assert.EqualValues(s.T(), 1, s.r[2])

	s.r.SetOSKAuto(false)
	assert.EqualValues(s.T(), 0, s.r[2])
}

func (s *CFR1TestSuite) TestSDIOInputOnly() {
	s.r[3] = 0
	assert.False(s.T(), s.r.SDIOInputOnly())

	s.r[3] = 1 << 1
	assert.True(s.T(), s.r.SDIOInputOnly())
}

func (s *CFR1TestSuite) TestSetSDIOInputOnly() {
	s.r[3] = 0

	s.r.SetSDIOInputOnly(true)
	assert.EqualValues(s.T(), 1<<1, s.r[3])

	s.r.SetSDIOInputOnly(false)
	assert.EqualValues(s.T(), 0, s.r[3])
}

func TestCFR1Suite(t *testing.T) {
	suite.Run(t, new(CFR1TestSuite))
}
