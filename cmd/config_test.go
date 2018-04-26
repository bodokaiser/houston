package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	suite.Suite

	c Config
}

func (s *ConfigTestSuite) SetupTest() {
	s.c = Config{
		Filename: "config.yaml",
	}
}

func (s *ConfigTestSuite) TestReadFromFile() {
	err := s.c.ReadFromFile()
	assert.NoError(s.T(), err)

	assert.Equal(s.T(), float64(1e9), s.c.DDS.SysClock)
	assert.Equal(s.T(), float64(1e7), s.c.DDS.RefClock)
	assert.Equal(s.T(), "SPI1.0", s.c.DDS.SPI.Device)
	assert.Equal(s.T(), "65", s.c.DDS.GPIO.Reset)
}

func (s *ConfigTestSuite) TestRender() {
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
