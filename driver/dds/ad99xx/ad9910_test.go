package ad99xx

import (
	"testing"

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
}

func TestAD9910Suite(t *testing.T) {
	suite.Run(t, new(AD9910TestSuite))
}