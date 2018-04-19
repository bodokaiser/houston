package ad9910

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CFR3TestSuite struct {
	suite.Suite

	r CFR3
}

func (s *CFR3TestSuite) SetupTest() {
	s.r = NewCFR3()
}

func (s *CFR3TestSuite) TestVCORange() {
	s.r[0] = 0
	assert.EqualValues(s.T(), VCORange0, s.r.VCORange())

	s.r[0] = 1
	assert.EqualValues(s.T(), VCORange1, s.r.VCORange())

	s.r[0] = 2
	assert.EqualValues(s.T(), VCORange2, s.r.VCORange())

	s.r[0] = 3
	assert.EqualValues(s.T(), VCORange3, s.r.VCORange())

	s.r[0] = 4
	assert.EqualValues(s.T(), VCORange4, s.r.VCORange())

	s.r[0] = 5
	assert.EqualValues(s.T(), VCORange5, s.r.VCORange())

	s.r[0] = 6
	assert.EqualValues(s.T(), VCORangeByPassed, s.r.VCORange())
}

func (s *CFR3TestSuite) TestSetVCORange() {
	s.r[0] = 0
	s.r.SetVCORange(VCORange0)
	assert.EqualValues(s.T(), 0, s.r[0])

	s.r[0] = 0
	s.r.SetVCORange(VCORange1)
	assert.EqualValues(s.T(), 1, s.r[0])

	s.r[0] = 0
	s.r.SetVCORange(VCORange2)
	assert.EqualValues(s.T(), 2, s.r[0])

	s.r[0] = 0
	s.r.SetVCORange(VCORange3)
	assert.EqualValues(s.T(), 3, s.r[0])

	s.r[0] = 0
	s.r.SetVCORange(VCORange4)
	assert.EqualValues(s.T(), 4, s.r[0])

	s.r[0] = 0
	s.r.SetVCORange(VCORange5)
	assert.EqualValues(s.T(), 5, s.r[0])

	s.r[0] = 0
	s.r.SetVCORange(VCORangeByPassed)
	assert.EqualValues(s.T(), 6, s.r[0])
}

func (s *CFR3TestSuite) TestPLLEnabled() {
	s.r[2] = 0
	assert.False(s.T(), s.r.PLLEnabled())

	s.r[2] = 1
	assert.True(s.T(), s.r.PLLEnabled())
}

func (s *CFR3TestSuite) TestSetPLLEnabled() {
	s.r[2] = 0

	s.r.SetPLLEnabled(true)
	assert.EqualValues(s.T(), 1, s.r[2])

	s.r.SetPLLEnabled(false)
	assert.EqualValues(s.T(), 0, s.r[2])
}

func (s *CFR3TestSuite) TestDivider() {
	s.r[3] = 0
	assert.EqualValues(s.T(), 0, s.r.Divider())

	s.r[3] = 100 << 1
	assert.EqualValues(s.T(), 100, s.r.Divider())
}

func (s *CFR3TestSuite) TestSetDivider() {
	s.r[3] = 0

	s.r.SetDivider(100)
	assert.EqualValues(s.T(), 100<<1, s.r[3])

	s.r.SetDivider(0)
	assert.EqualValues(s.T(), 0, s.r[3])
}

func TestCFR3Suite(t *testing.T) {
	suite.Run(t, new(CFR3TestSuite))
}
