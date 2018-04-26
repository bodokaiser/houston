package ad9910

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bodokaiser/houston/driver/dds"
)

type AD9910TestSuite struct {
	suite.Suite

	d *AD9910
}

func (s *AD9910TestSuite) SetupTest() {
	c := dds.Config{}
	c.SysClock = 1e9
	c.RefClock = 1e7

	s.d = NewAD9910(c)
}

func (s *AD9910TestSuite) TestSingleTone() {
}

func TestAD9910Suite(t *testing.T) {
	suite.Run(t, new(AD9910TestSuite))
}
