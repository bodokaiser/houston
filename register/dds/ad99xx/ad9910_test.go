package ad99xx

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AD9910TestSuite struct {
	suite.Suite

	device *AD9910
}

func (s *AD9910TestSuite) SetupTest() {
	s.device = NewAD9910()
}

func (s *AD9910TestSuite) TestNew() {
	assert.Equal(s.T(), [5]byte{0x00, 0x00, 0x00, 0x00, 0x00}, s.device.CFR1)
	assert.Equal(s.T(), [5]byte{0x01, 0x00, 0x40, 0x08, 0x20}, s.device.CFR2)
	assert.Equal(s.T(), [5]byte{0x02, 0x17, 0x38, 0x40, 0x00}, s.device.CFR3)

	assert.Equal(s.T(), [5]byte{0x03, 0x00, 0x00, 0x00, 0x7f}, s.device.AuxDAC)
	assert.Equal(s.T(), [5]byte{0x04, 0xff, 0xff, 0xff, 0xff}, s.device.IOUpdateRate)
	assert.Equal(s.T(), [5]byte{0x07}, s.device.FTW)
	assert.Equal(s.T(), [5]byte{0x08}, s.device.POW)
	assert.Equal(s.T(), [5]byte{0x09}, s.device.ASF)
	assert.Equal(s.T(), [5]byte{0x0a}, s.device.MultiChip)
	assert.Equal(s.T(), [9]byte{0x0b}, s.device.DRampLimit)
	assert.Equal(s.T(), [9]byte{0x0c}, s.device.DRampStepSize)
	assert.Equal(s.T(), [9]byte{0x0d}, s.device.DRampRate)
	assert.Equal(s.T(), [9]byte{0x0e, 0x08, 0xb5}, s.device.STProfile0)
	assert.Equal(s.T(), [9]byte{0x0f}, s.device.STProfile1)
	assert.Equal(s.T(), [9]byte{0x10}, s.device.STProfile2)
	assert.Equal(s.T(), [9]byte{0x11}, s.device.STProfile3)
	assert.Equal(s.T(), [9]byte{0x12}, s.device.STProfile4)
	assert.Equal(s.T(), [9]byte{0x13}, s.device.STProfile5)
	assert.Equal(s.T(), [9]byte{0x14}, s.device.STProfile6)
	assert.Equal(s.T(), [9]byte{0x15}, s.device.STProfile7)
	assert.Equal(s.T(), [9]byte{0x0e}, s.device.RAMProfile0)
	assert.Equal(s.T(), [9]byte{0x0f}, s.device.RAMProfile1)
	assert.Equal(s.T(), [9]byte{0x10}, s.device.RAMProfile2)
	assert.Equal(s.T(), [9]byte{0x11}, s.device.RAMProfile3)
	assert.Equal(s.T(), [9]byte{0x12}, s.device.RAMProfile4)
	assert.Equal(s.T(), [9]byte{0x13}, s.device.RAMProfile5)
	assert.Equal(s.T(), [9]byte{0x14}, s.device.RAMProfile6)
	assert.Equal(s.T(), [9]byte{0x15}, s.device.RAMProfile7)
}

func TestAD9910Suite(t *testing.T) {
	suite.Run(t, new(AD9910TestSuite))
}
