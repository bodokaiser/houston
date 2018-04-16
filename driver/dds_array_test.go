package driver

import (
	"testing"

	"github.com/bodokaiser/houston/driver/dds"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DDSArrayTestSuite struct {
	suite.Suite

	d DDSArray
}

func (s *DDSArrayTestSuite) TestMocked() {
	s.d = &MockedDDSArray{}

	sc := dds.SingleToneConfig{
		Amplitude: 1.0,
		Frequency: 10e6,
	}
	dc := dds.DigitalRampConfig{
		SingleToneConfig: dds.SingleToneConfig{
			Amplitude: 1.0,
		},
		Limits:      [2]float64{10e6, 100e6},
		StepSize:    [2]float64{1e6, 1e6},
		SlopeRate:   [2]float64{1, 1},
		NoDwellHigh: true,
		NoDwellLow:  true,
		Destination: dds.Frequency,
	}
	pc := dds.PlaybackConfig{
		SingleToneConfig: dds.SingleToneConfig{
			Frequency: 20e6,
		},
		Rate:        1e3,
		Data:        []byte{1, 2, 3, 4, 5, 6, 7, 8, 9},
		Destination: dds.Amplitude,
	}

	assert.NoError(s.T(), s.d.Select(2))
	assert.NoError(s.T(), s.d.SingleTone(sc))
	assert.NoError(s.T(), s.d.DigitalRamp(dc))
	assert.NoError(s.T(), s.d.Playback(pc))

	d := s.d.(*MockedDDSArray)
	assert.EqualValues(s.T(), 2, d.Address)
	assert.EqualValues(s.T(), sc, d.SingleToneConfig)
	assert.EqualValues(s.T(), dc, d.DigitalRampConfig)
	assert.EqualValues(s.T(), pc, d.PlaybackConfig)
}

func TestDDSArrayTestSuite(t *testing.T) {
	suite.Run(t, new(DDSArrayTestSuite))
}
