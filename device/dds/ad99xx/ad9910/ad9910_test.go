package ad9910

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/bodokaiser/houston/device/dds"
)

type AD9910TestSuite struct {
	suite.Suite

	d AD9910
}

func (s *AD9910TestSuite) SetupTest() {
	s.d = NewAD9910(dds.Config{
		SysClock: 1e9,
		RefClock: 1e7,
	})
}

func (s *AD9910TestSuite) TestFreqToFTW() {
	fout := 41e6     // 41 MHz
	fsys := 122880e3 // 122.88 MHz

	assert.Equal(s.T(), uint32(0x556AAAAB), freqToFTW(fout, fsys))
}

func TestAD9910Suite(t *testing.T) {
	suite.Run(t, new(AD9910TestSuite))
}
