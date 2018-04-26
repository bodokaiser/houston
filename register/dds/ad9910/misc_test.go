package ad9910

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FTWTestSuite struct {
	suite.Suite

	r FTW
}

func (s *FTWTestSuite) SetupTest() {
	s.r = NewFTW()
}

func (s *FTWTestSuite) TestFreqTuningWord() {
	s.r[3] = 0
	assert.Equal(s.T(), uint32(0), s.r.FreqTuningWord())

	s.r[3] = 1
	assert.Equal(s.T(), uint32(1), s.r.FreqTuningWord())

	s.r[2] = 1
	assert.Equal(s.T(), uint32(1<<8+1), s.r.FreqTuningWord())
}

func (s *FTWTestSuite) TestSetFreqTuningWord() {
	s.r.SetFreqTuningWord(100)
	assert.Equal(s.T(), byte(100), s.r[3])
}

func TestFTWSuite(t *testing.T) {
	suite.Run(t, new(FTWTestSuite))
}

type POWTestSuite struct {
	suite.Suite

	r POW
}

func (s *POWTestSuite) SetupTest() {
	s.r = NewPOW()
}

func (s *POWTestSuite) TestPhaseOffsetWord() {
	s.r[1] = 0
	assert.Equal(s.T(), uint16(0), s.r.PhaseOffsetWord())

	s.r[1] = 1
	assert.Equal(s.T(), uint16(1), s.r.PhaseOffsetWord())

	s.r[0] = 1
	assert.Equal(s.T(), uint16(1<<8+1), s.r.PhaseOffsetWord())
}

func (s *POWTestSuite) TestSetPhaseOffsetWord() {
	s.r.SetPhaseOffsetWord(100)
	assert.Equal(s.T(), byte(100), s.r[1])
}

func TestPOWSuite(t *testing.T) {
	suite.Run(t, new(POWTestSuite))
}

type ASFTestSuite struct {
	suite.Suite

	r ASF
}

func (s *ASFTestSuite) SetupTest() {
	s.r = NewASF()
}

func (s *ASFTestSuite) TestAmplScaleFactor() {
	s.r[3] = 0
	assert.Equal(s.T(), uint16(0), s.r.AmplScaleFactor())

	s.r[3] = 3
	assert.Equal(s.T(), uint16(0), s.r.AmplScaleFactor())

	s.r[3] = 255
	assert.Equal(s.T(), uint16(1<<6-1), s.r.AmplScaleFactor())
}

func (s *ASFTestSuite) TestSetAmplScaleFactor() {
	s.r.SetAmplScaleFactor(1)
	assert.Equal(s.T(), byte(1<<2), s.r[3])
}

func TestASFSuite(t *testing.T) {
	suite.Run(t, new(ASFTestSuite))
}
