package ad9910

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bodokaiser/houston/device/dds"
)

type AD9910TestSuite struct {
	suite.Suite

	d AD9910
}

func (s *AD9910TestSuite) SetupTest() {
	s.d = NewAD9910(dds.Config{1e9, 1e7})
}

func (s *AD9910TestSuite) TestSingleTone() {
}

func TestAD9910Suite(t *testing.T) {
	suite.Run(t, new(AD9910TestSuite))
}
