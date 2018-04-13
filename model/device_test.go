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
	assert.NoError(s.T(), DDSDevice{
		Name:      "DDS0",
		Address:   0,
		Amplitude: 1.0,
		Frequency: 200e6,
	}.Validate())

	assert.Error(s.T(), DDSDevice{
		Name:      "DDS0",
		Address:   0,
		Amplitude: 1.0,
		Frequency: 800e6,
	}.Validate())

	assert.Error(s.T(), DDSDevice{
		Name:      "DDS0",
		Address:   0,
		Amplitude: 1.1,
		Frequency: 200e6,
	}.Validate())
}

func TestDeviceTestSuite(t *testing.T) {
	suite.Run(t, new(DDSDeviceTestSuite))
}

type DDSDevicesTestSuite struct {
	suite.Suite

	d DDSDevices
}

func (s *DDSDevicesTestSuite) SetupTest() {
	s.d = DDSDevices{
		DDSDevice{Name: "DDS0", Address: 0},
		DDSDevice{Name: "DDS3", Address: 3},
	}
}

func (s *DDSDevicesTestSuite) TestFindByName() {
	assert.Equal(s.T(), 0, s.d.FindByName("DDS0"))
	assert.Equal(s.T(), 1, s.d.FindByName("DDS3"))
	assert.Equal(s.T(), -1, s.d.FindByName("DDS2"))
}

func (s *DDSDevicesTestSuite) TestString() {
	assert.Equal(s.T(), "0,3", s.d.String())
}

func (s *DDSDevicesTestSuite) TestSet() {
	d := DDSDevices{}

	assert.NoError(s.T(), d.Set("1,2"))
	assert.Equal(s.T(), "DDS1", d[0].Name)
	assert.Equal(s.T(), "DDS2", d[1].Name)
	assert.Equal(s.T(), uint8(1), d[0].Address)
	assert.Equal(s.T(), uint8(2), d[1].Address)

	assert.Error(s.T(), d.Set(""))
	assert.Error(s.T(), d.Set("x,y"))
}

func TestDevicesTestSuite(t *testing.T) {
	suite.Run(t, new(DDSDevicesTestSuite))
}
