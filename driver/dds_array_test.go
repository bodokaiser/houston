package driver

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DDSArrayTestSuite struct {
	suite.Suite

	d DDSArray
}

func (s *DDSArrayTestSuite) TestMocked() {
	s.d = &MockedDDSArray{}

	assert.NoError(s.T(), s.d.Select(2))
	assert.NoError(s.T(), s.d.SingleTone(1.0, 10e6, 0))

	d := s.d.(*MockedDDSArray)
	assert.EqualValues(s.T(), 2, d.Address)
	assert.EqualValues(s.T(), 0, d.Phase)
	assert.EqualValues(s.T(), 1.0, d.Amplitude)
	assert.EqualValues(s.T(), 10e6, d.Frequency)
}

func TestDDSArrayTestSuite(t *testing.T) {
	suite.Run(t, new(DDSArrayTestSuite))
}
