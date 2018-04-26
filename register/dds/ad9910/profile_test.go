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

type RAMProfileTestSuite struct {
	suite.Suite

	r RAMProfile
}

func (s *RAMProfileTestSuite) SetupTest() {
	s.r = NewRAMProfile()
}

func (s *RAMProfileTestSuite) TestAddrStepRate() {
	s.r[2] = 60

	assert.Equal(s.T(), uint16(60), s.r.AddrStepRate())
}

func (s *RAMProfileTestSuite) TestSetAddrStepRate() {
	s.r.SetAddrStepRate(256)

	assert.Equal(s.T(), []byte{1, 0}, s.r[1:3])
}

func (s *RAMProfileTestSuite) TestWaveformEndAddr() {
	s.r[3] = 0x00
	s.r[4] = 0x0a
	assert.Equal(s.T(), uint16(0), s.r.WaveformEndAddr())

	s.r[3] = 0x00
	s.r[4] = 0x40
	assert.Equal(s.T(), uint16(1), s.r.WaveformEndAddr())

	s.r[3] = 0x01
	s.r[4] = 0x00
	assert.Equal(s.T(), uint16(4), s.r.WaveformEndAddr())
}

func (s *RAMProfileTestSuite) TestSetWaveformEndAddr() {
	s.r.SetWaveformEndAddr(4)
	assert.Equal(s.T(), []byte{0x01, 0x00}, s.r[3:5])

	s.r.SetWaveformEndAddr(1)
	assert.Equal(s.T(), []byte{0x00, 0x40}, s.r[3:5])
}

func (s *RAMProfileTestSuite) TestWaveformStartAddr() {
	s.r[5] = 0x00
	s.r[6] = 0x0a
	assert.Equal(s.T(), uint16(0), s.r.WaveformStartAddr())

	s.r[5] = 0x00
	s.r[6] = 0x40
	assert.Equal(s.T(), uint16(1), s.r.WaveformStartAddr())

	s.r[5] = 0x01
	s.r[6] = 0x00
	assert.Equal(s.T(), uint16(4), s.r.WaveformStartAddr())
}

func (s *RAMProfileTestSuite) TestSetWaveformStartAddr() {
	s.r.SetWaveformStartAddr(4)
	assert.Equal(s.T(), []byte{0x01, 0x00}, s.r[5:7])

	s.r.SetWaveformStartAddr(1)
	assert.Equal(s.T(), []byte{0x00, 0x40}, s.r[5:7])
}

func (s *RAMProfileTestSuite) TestRAMContorlMode() {
	s.r[7] = 0
	assert.Equal(s.T(), RAMControlModeDirectSwitch, s.r.RAMControlMode())

	s.r[7] = 1
	assert.Equal(s.T(), RAMControlModeRampUp, s.r.RAMControlMode())

	s.r[7] = 2
	assert.Equal(s.T(), RAMControlModeBiRampUp, s.r.RAMControlMode())

	s.r[7] = 3
	assert.Equal(s.T(), RAMControlModeContBiRampUp, s.r.RAMControlMode())

	s.r[7] = 4
	assert.Equal(s.T(), RAMControlModeContRecirculate, s.r.RAMControlMode())

	s.r[7] = 5
	assert.Equal(s.T(), RAMControlModeDirectSwitch, s.r.RAMControlMode())

	s.r[7] = 6
	assert.Equal(s.T(), RAMControlModeDirectSwitch, s.r.RAMControlMode())

	s.r[7] = 7
	assert.Equal(s.T(), RAMControlModeDirectSwitch, s.r.RAMControlMode())
}

func (s *RAMProfileTestSuite) TestSetRAMContorlMode() {
	s.r[7] = 0x80
	s.r.SetRAMControlMode(RAMControlModeDirectSwitch)
	assert.Equal(s.T(), []byte{0x80}, s.r[7:8])

	s.r.SetRAMControlMode(RAMControlModeRampUp)
	assert.Equal(s.T(), []byte{0x81}, s.r[7:8])

	s.r.SetRAMControlMode(RAMControlModeBiRampUp)
	assert.Equal(s.T(), []byte{0x82}, s.r[7:8])

	s.r.SetRAMControlMode(RAMControlModeContBiRampUp)
	assert.Equal(s.T(), []byte{0x83}, s.r[7:8])

	s.r.SetRAMControlMode(RAMControlModeContRecirculate)
	assert.Equal(s.T(), []byte{0x84}, s.r[7:8])
}

func (s *RAMProfileTestSuite) TestNoDwellHigh() {
	s.r[7] = 0x20
	assert.True(s.T(), s.r.NoDwellHigh())

	s.r[7] = 0x00
	assert.False(s.T(), s.r.NoDwellHigh())
}

func (s *RAMProfileTestSuite) TestSetNoDwellHigh() {
	s.r[7] = 0xa0
	s.r.SetNoDwellHigh(false)
	assert.Equal(s.T(), []byte{0x80}, s.r[7:8])

	s.r[7] = 0x80
	s.r.SetNoDwellHigh(true)
	assert.Equal(s.T(), []byte{0xa0}, s.r[7:8])
}

func (s *RAMProfileTestSuite) TestZeroCrossing() {
	s.r[7] = 0x08
	assert.True(s.T(), s.r.ZeroCrossing())

	s.r[7] = 0x00
	assert.False(s.T(), s.r.ZeroCrossing())
}

func (s *RAMProfileTestSuite) TestSetZeroCrossing() {
	s.r[7] = 0x88
	s.r.SetZeroCrossing(false)
	assert.Equal(s.T(), []byte{0x80}, s.r[7:8])

	s.r[7] = 0x80
	s.r.SetZeroCrossing(true)
	assert.Equal(s.T(), []byte{0x88}, s.r[7:8])
}

func TestRAMProfileSuite(t *testing.T) {
	suite.Run(t, new(RAMProfileTestSuite))
}
