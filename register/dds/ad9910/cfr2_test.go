package ad9910

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CFR2TestSuite struct {
	suite.Suite

	r CFR2
}

func (s *CFR2TestSuite) SetupTest() {
	s.r = NewCFR2()
}

func (s *CFR2TestSuite) TestSTAmplScaleEnabled() {
	s.r[0] = 0x00
	assert.False(s.T(), s.r.STAmplScaleEnabled())

	s.r[0] = 0x01
	assert.True(s.T(), s.r.STAmplScaleEnabled())
}

func (s *CFR2TestSuite) TestSetSTAmplScaleEnabled() {
	s.r.SetSTAmplScaleEnabled(true)
	assert.Equal(s.T(), []byte{0x01}, s.r[0:1])

	s.r.SetSTAmplScaleEnabled(false)
	assert.Equal(s.T(), []byte{0x00}, s.r[0:1])
}

func (s *CFR2TestSuite) TestSyncClockEnabled() {
	s.r[1] = 0
	assert.False(s.T(), s.r.SyncClockEnabled())

	s.r[1] = 1 << 6
	assert.True(s.T(), s.r.SyncClockEnabled())
}

func (s *CFR2TestSuite) TestSetSyncClockEnabled() {
	s.r.SetSyncClockEnabled(true)
	assert.Equal(s.T(), []byte{1 << 6}, s.r[1:2])

	s.r.SetSyncClockEnabled(false)
	assert.Equal(s.T(), []byte{0}, s.r[1:2])
}

func (s *CFR2TestSuite) TestSyncTimingValidationDisabled() {
	s.r[3] = 0
	assert.False(s.T(), s.r.SyncTimingValidationDisabled())

	s.r[3] = 1 << 5
	assert.True(s.T(), s.r.SyncTimingValidationDisabled())
}

func (s *CFR2TestSuite) TestSetSyncTimingValidationDisabled() {
	s.r[3] = 0

	s.r.SetSyncTimingValidationDisabled(true)
	assert.Equal(s.T(), []byte{1 << 5}, s.r[3:4])

	s.r.SetSyncTimingValidationDisabled(false)
	assert.Equal(s.T(), []byte{0}, s.r[3:4])
}

func (s *CFR2TestSuite) TestRampEnabled() {
	s.r[1] = 0
	assert.False(s.T(), s.r.RampEnabled())

	s.r[1] = 1 << 3
	assert.True(s.T(), s.r.RampEnabled())
}

func (s *CFR2TestSuite) TestSetRampEnabled() {
	s.r[1] = 0

	s.r.SetRampEnabled(true)
	assert.Equal(s.T(), []byte{1 << 3}, s.r[1:2])

	s.r.SetRampEnabled(false)
	assert.Equal(s.T(), []byte{0x00}, s.r[1:2])
}

func (s *CFR2TestSuite) TestRampDest() {
	s.r[1] = 0
	assert.Equal(s.T(), RampDestFrequency, s.r.RampDest())

	s.r[1] = 1 << 4
	assert.Equal(s.T(), RampDestPhase, s.r.RampDest())

	s.r[1] = 1 << 5
	assert.Equal(s.T(), RampDestAmplitude, s.r.RampDest())

	s.r[1] = (1 << 5) + (1 << 4)
	assert.Equal(s.T(), RampDestAmplitude, s.r.RampDest())
}

func (s *CFR2TestSuite) TestSetRampDest() {
	s.r[1] = 0
	s.r.SetRampDest(RampDestFrequency)
	assert.Equal(s.T(), []byte{0}, s.r[1:2])

	s.r[1] = 0
	s.r.SetRampDest(RampDestPhase)
	assert.Equal(s.T(), []byte{1 << 4}, s.r[1:2])

	s.r[1] = 0
	s.r.SetRampDest(RampDestAmplitude)
	assert.Equal(s.T(), []byte{1 << 5}, s.r[1:2])

	s.r[1] = 0x80
	s.r.SetRampDest(RampDestAmplitude)
	assert.Equal(s.T(), []byte{0xa0}, s.r[1:2])
}

func (s *CFR2TestSuite) TestRampNoDwellLow() {
	s.r[1] = 0
	assert.False(s.T(), s.r.RampNoDwellLow())

	s.r[1] = 1 << 1
	assert.True(s.T(), s.r.RampNoDwellLow())
}

func (s *CFR2TestSuite) TestSetRampNoDwellLow() {
	s.r[1] = 0x00

	s.r.SetRampNoDwellLow(true)
	assert.Equal(s.T(), []byte{0x02}, s.r[1:2])

	s.r.SetRampNoDwellLow(false)
	assert.Equal(s.T(), []byte{0x00}, s.r[1:2])
}

func (s *CFR2TestSuite) TestRampNoDwellHigh() {
	s.r[1] = 0
	assert.False(s.T(), s.r.RampNoDwellHigh())

	s.r[1] = 1 << 2
	assert.True(s.T(), s.r.RampNoDwellHigh())
}

func (s *CFR2TestSuite) TestSetRampNoDwellHigh() {
	s.r[1] = 0

	s.r.SetRampNoDwellHigh(true)
	assert.Equal(s.T(), []byte{1 << 2}, s.r[1:2])

	s.r.SetRampNoDwellHigh(false)
	assert.Equal(s.T(), []byte{0x00}, s.r[1:2])
}

func TestCFR2Suite(t *testing.T) {
	suite.Run(t, new(CFR2TestSuite))
}
