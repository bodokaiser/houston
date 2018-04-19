package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DDSDeviceTestSuite struct {
	suite.Suite
}

func (s *DDSDeviceTestSuite) TestValidation() {
	d1 := &DDSDevice{}
	d2 := &DDSDevice{ID: 0}
	d3 := &DDSDevice{ID: 0, Name: "DDS0"}
	d4 := &DDSDevice{
		ID:   0,
		Name: "DDS0",
		Amplitude: DDSParam{
			DDSConst: &DDSConst{Value: 1.0},
		},
		Frequency: DDSParam{
			DDSConst: &DDSConst{Value: 200},
		},
		PhaseOffset: DDSParam{
			DDSConst: &DDSConst{Value: 0},
		},
	}

	assert.Error(s.T(), d1.Validate())
	assert.Error(s.T(), d2.Validate())
	assert.Error(s.T(), d3.Validate())
	assert.NoError(s.T(), d4.Validate())
}

func TestDeviceTestSuite(t *testing.T) {
	suite.Run(t, new(DDSDeviceTestSuite))
}
