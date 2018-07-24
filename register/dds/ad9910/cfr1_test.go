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

	s.r[0] = 0x80
	assert.True(s.T(), s.r.RAMEnabled())
}

func (s *CFR1TestSuite) TestSetRamEnabled() {
	s.r[0] = 0

	s.r.SetRAMEnabled(true)
	assert.Equal(s.T(), []byte{0x80}, s.r[0:1])

	s.r.SetRAMEnabled(false)
	assert.False(s.T(), s.r[0]&1<<7 > 0)
}

func (s *CFR1TestSuite) TestRAMDest() {
	s.r[0] = 0x80
	assert.Equal(s.T(), RAMDestFrequency, s.r.RAMDest())

	s.r[0] = 0xa0
	assert.Equal(s.T(), RAMDestPhase, s.r.RAMDest())

	s.r[0] = 0xc0
	assert.Equal(s.T(), RAMDestAmplitude, s.r.RAMDest())

	s.r[0] = 0xe0
	assert.Equal(s.T(), RAMDestPolar, s.r.RAMDest())
}

func (s *CFR1TestSuite) TestSetRAMDest() {
	s.r[0] = 0x00
	s.r.SetRAMDest(RAMDestFrequency)
	assert.Equal(s.T(), []byte{0x00}, s.r[0:1])
	s.r.SetRAMDest(RAMDestPhase)
	assert.Equal(s.T(), []byte{0x20}, s.r[0:1])
	s.r.SetRAMDest(RAMDestAmplitude)
	assert.Equal(s.T(), []byte{0x40}, s.r[0:1])
	s.r.SetRAMDest(RAMDestPolar)
	assert.Equal(s.T(), []byte{0x60}, s.r[0:1])

	s.r[0] = 0x80
	s.r.SetRAMDest(RAMDestFrequency)
	assert.Equal(s.T(), []byte{0x80}, s.r[0:1])
	s.r.SetRAMDest(RAMDestPhase)
	assert.Equal(s.T(), []byte{0xa0}, s.r[0:1])
	s.r.SetRAMDest(RAMDestAmplitude)
	assert.Equal(s.T(), []byte{0xc0}, s.r[0:1])
	s.r.SetRAMDest(RAMDestPolar)
	assert.Equal(s.T(), []byte{0xe0}, s.r[0:1])
}

func (s *CFR1TestSuite) TestOSKEnabled() {
	s.r[2] = 0
	assert.False(s.T(), s.r.OSKEnabled())

	s.r[2] = 1 << 1
	assert.True(s.T(), s.r.OSKEnabled())
}

func (s *CFR1TestSuite) TestSetOSKEnabled() {
	s.r.SetOSKEnabled(true)
	assert.Equal(s.T(), []byte{0x02}, s.r[2:3])

	s.r.SetOSKEnabled(false)
	assert.Equal(s.T(), []byte{0x00}, s.r[2:3])
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
	assert.Equal(s.T(), []byte{0x80}, s.r[1:2])

	s.r.SetOSKManual(false)
	assert.Equal(s.T(), []byte{0x00}, s.r[1:2])
}

func (s *CFR1TestSuite) TestOSKAuto() {
	s.r[2] = 0x00
	assert.False(s.T(), s.r.OSKAuto())

	s.r[2] = 0x01
	assert.True(s.T(), s.r.OSKAuto())
}

func (s *CFR1TestSuite) TestSetOSKAuto() {
	s.r.SetOSKAuto(true)
	assert.Equal(s.T(), []byte{0x01}, s.r[2:3])

	s.r.SetOSKAuto(false)
	assert.Equal(s.T(), []byte{0x00}, s.r[2:3])
}

func (s *CFR1TestSuite) TestSDIOInputOnly() {
	s.r[3] = 0x00
	assert.False(s.T(), s.r.SDIOInputOnly())

	s.r[3] = 0x02
	assert.True(s.T(), s.r.SDIOInputOnly())
}

func (s *CFR1TestSuite) TestSetSDIOInputOnly() {
	s.r.SetSDIOInputOnly(true)
	assert.Equal(s.T(), []byte{0x02}, s.r[3:4])

	s.r.SetSDIOInputOnly(false)
	assert.Equal(s.T(), []byte{0x00}, s.r[3:4])
}

func (s *CFR1TestSuite) TestInverseSincFilter() {
	s.r[1] = 0x80
	assert.False(s.T(), s.r.InverseSincFilter())

	s.r[1] = 0x40
	assert.True(s.T(), s.r.InverseSincFilter())
}

func (s *CFR1TestSuite) TestSetInverseSincFilter() {
	s.r.SetInverseSincFilter(true)
	assert.Equal(s.T(), []byte{0x40}, s.r[1:2])

	s.r.SetInverseSincFilter(false)
	assert.Equal(s.T(), []byte{0x00}, s.r[1:2])
}

func TestCFR1Suite(t *testing.T) {
	suite.Run(t, new(CFR1TestSuite))
}
