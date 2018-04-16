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
	d1 := &DDSDevice{
		ID:   0,
		Name: "DDS0",
	}
	d2 := &DDSDevice{
		ID:   0,
		Name: "DDS0",
	}
	d3 := &DDSDevice{
		ID:   0,
		Name: "DDS0",
	}

	assert.NoError(s.T(), d1.Validate())
	assert.NoError(s.T(), d2.Validate())
	assert.NoError(s.T(), d3.Validate())
}

func TestDeviceTestSuite(t *testing.T) {
	suite.Run(t, new(DDSDeviceTestSuite))
}
