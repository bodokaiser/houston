package ad9910

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RampTestSuite struct {
	suite.Suite

	l RampLimit
	s RampStep
	r RampRate
}

func (s *RampTestSuite) SetupTest() {
	s.l = NewRampLimit()
	s.s = NewRampStep()
	s.r = NewRampRate()
}

func (s *RampTestSuite) TestRampLimit() {
	assert.EqualValues(s.T(), 0, s.l.UpperLimit())
	assert.EqualValues(s.T(), 0, s.l.LowerLimit())

	s.l[3] = 255
	s.l[6] = 1

	assert.EqualValues(s.T(), 255, s.l.UpperLimit())
	assert.EqualValues(s.T(), 256, s.l.LowerLimit())
}

func (s *RampTestSuite) TestSetRampLimit() {
	s.l.SetUpperLimit(10)
	s.l.SetLowerLimit(256)

	assert.EqualValues(s.T(), []byte{0, 0, 0, 10, 0, 0, 1, 0}, s.l[:])
}

func (s *RampTestSuite) TestRampStepSize() {
	assert.EqualValues(s.T(), 0, s.s.DecrStepSize())
	assert.EqualValues(s.T(), 0, s.s.IncrStepSize())

	s.s[3] = 255
	s.s[6] = 1

	assert.EqualValues(s.T(), 255, s.s.DecrStepSize())
	assert.EqualValues(s.T(), 256, s.s.IncrStepSize())
}

func (s *RampTestSuite) TestSetRampStepSize() {
	s.s.SetDecrStepSize(10)
	s.s.SetIncrStepSize(256)

	assert.EqualValues(s.T(), []byte{0, 0, 0, 10, 0, 0, 1, 0}, s.s[:])
}

func (s *RampTestSuite) TestRampSlopeRate() {
	assert.EqualValues(s.T(), 0, s.r.NegSlopeRate())
	assert.EqualValues(s.T(), 0, s.r.PosSlopeRate())

	s.r[1] = 255
	s.r[2] = 1

	assert.EqualValues(s.T(), 255, s.r.NegSlopeRate())
	assert.EqualValues(s.T(), 256, s.r.PosSlopeRate())
}

func (s *RampTestSuite) TestSetRampSlopeRate() {
	s.r.SetNegSlopeRate(10)
	s.r.SetPosSlopeRate(256)

	assert.EqualValues(s.T(), []byte{0, 10, 1, 0}, s.r[:])
}

func TestRampSuite(t *testing.T) {
	suite.Run(t, new(RampTestSuite))
}
