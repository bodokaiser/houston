package ad99xx

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AD9910TestSuite struct {
	suite.Suite

	d *AD9910
}

func (s *AD9910TestSuite) SetupTest() {
	s.d = NewAD9910(Config{
		SysClock: 1e9,
		RefClock: 1e7,
	})
}

func (s *AD9910TestSuite) TestSingleTone() {
	assert.Equal(s.T(), []byte{
		// CFR1
		0x00, 0x00, 0x00, 0x00, 0x02,
		// CFR2
		0x01, 0x01, 0x40, 0x00, 0x00,
		// CFR3
		0x02, 0x00, 0x00, 0x00, 0x20,
		// STProfile0
		0x0e, 0x08, 0xb5, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}, s.d.toBytes())
}

func TestAD9910Suite(t *testing.T) {
	suite.Run(t, new(AD9910TestSuite))
}
