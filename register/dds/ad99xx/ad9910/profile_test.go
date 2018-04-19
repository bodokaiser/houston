package ad9910

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type STProfileTestSuite struct {
	suite.Suite

	r STProfile
}

func (s *STProfileTestSuite) SetupTest() {
	s.r = NewSTProfile()
}

func (s *STProfileTestSuite) TestAmplScaleFactor() {
	assert.Equal(s.T(), uint16(2229), s.r.AmplScaleFactor())

	s.r[0] = 0
	s.r[1] = 0
	assert.Equal(s.T(), uint16(0), s.r.AmplScaleFactor())

	s.r[1] = 5
	assert.Equal(s.T(), uint16(5), s.r.AmplScaleFactor())
}

func (s *STProfileTestSuite) TestSetAmplScaleFactor() {
	s.r.SetAmplScaleFactor(0)
	assert.Equal(s.T(), []byte{0x00, 0x00}, []byte(s.r[0:2]))

	s.r.SetAmplScaleFactor(257)
	assert.Equal(s.T(), []byte{0x01, 0x01}, []byte(s.r[0:2]))
}

func (s *STProfileTestSuite) TestPhaseOffsetWord() {
	assert.Equal(s.T(), uint16(0), s.r.PhaseOffsetWord())

	s.r[2] = 1
	s.r[3] = 1
	assert.Equal(s.T(), uint16(257), s.r.PhaseOffsetWord())
}

func (s *STProfileTestSuite) TestSetPhaseOffsetWord() {
	s.r.SetPhaseOffsetWord(0)
	assert.Equal(s.T(), []byte{0x00, 0x00}, []byte(s.r[2:4]))

	s.r.SetPhaseOffsetWord(257)
	assert.Equal(s.T(), []byte{0x01, 0x01}, []byte(s.r[2:4]))
}

func (s *STProfileTestSuite) TestFreqTuningWord() {
	assert.Equal(s.T(), uint32(0), s.r.FreqTuningWord())

	s.r[7] = 1
	assert.Equal(s.T(), uint32(1), s.r.FreqTuningWord())
}

func (s *STProfileTestSuite) TestSetFreqTuningWord() {
	s.r.SetFreqTuningWord(0)
	assert.Equal(s.T(), []byte{0x00, 0x00, 0x00, 0x00}, []byte(s.r[4:8]))

	s.r.SetFreqTuningWord(1<<32 - 1)
	assert.Equal(s.T(), []byte{0xff, 0xff, 0xff, 0xff}, []byte(s.r[4:8]))
}

func TestSTProfileSuite(t *testing.T) {
	suite.Run(t, new(STProfileTestSuite))
}
