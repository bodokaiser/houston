package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	suite.Suite
}

func (s *ConfigTestSuite) TestReadFromFile() {
	c := Config{}

	err := c.ReadFromFile("beagle.yaml")
	assert.NoError(s.T(), err)

	assert.Equal(s.T(), float64(1e9), c.DDS.SysClock)
	assert.Equal(s.T(), float64(1e7), c.DDS.RefClock)
	assert.Equal(s.T(), "SPI1.0", c.DDS.SPI.Device)
	assert.Equal(s.T(), "65", c.DDS.GPIO.Reset)
	assert.Equal(s.T(), "27", c.DDS.GPIO.Update)
}

func (s *ConfigTestSuite) TestReadFromBox() {
	c := Config{}

	err := c.ReadFromBox("beagle.yaml")
	assert.NoError(s.T(), err)

	assert.Equal(s.T(), float64(1e9), c.DDS.SysClock)
	assert.Equal(s.T(), float64(1e7), c.DDS.RefClock)
	assert.Equal(s.T(), "SPI1.0", c.DDS.SPI.Device)
	assert.Equal(s.T(), "65", c.DDS.GPIO.Reset)
	assert.Equal(s.T(), "27", c.DDS.GPIO.Update)
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
