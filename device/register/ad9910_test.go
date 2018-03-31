package register

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AD9910TestSuite struct {
	suite.Suite

	r *AD9910
}

func (s *AD9910TestSuite) SetupTest() {
	s.r = new(AD9910)
	s.r.CtrlFunc1 = AD9910CtrlFunc1Default
	s.r.CtrlFunc2 = AD9910CtrlFunc2Default
	s.r.CtrlFunc3 = AD9910CtrlFunc3Default
}

func (s *AD9910TestSuite) TestLSBFirst() {
	s.r.CtrlFunc1[0] = 0x00
	assert.False(s.T(), s.r.LSBFirst())

	s.r.CtrlFunc1[0] = 0x01
	assert.True(s.T(), s.r.LSBFirst())
}

func (s *AD9910TestSuite) TestSetLSBFirst() {
	s.r.CtrlFunc1 = AD9910CtrlFunc1Default

	s.r.SetLSBFirst(true)
	assert.Equal(s.T(), s.r.CtrlFunc1[0], byte(0x01), "not active")

	s.r.SetLSBFirst(false)
	assert.Equal(s.T(), s.r.CtrlFunc1[0], byte(0x00), "not inactive")
}

func (s *AD9910TestSuite) TestSDIOInputOnly() {
	s.r.CtrlFunc1[0] = 0x00
	assert.False(s.T(), s.r.SDIOInputOnly())

	s.r.CtrlFunc1[0] = 0x02
	assert.True(s.T(), s.r.SDIOInputOnly())
}

func TestAD9910Constants(t *testing.T) {
	assert.Equal(t, 0x00, AD9910CtrlFunc1Address)
	assert.Equal(t, 0x04, AD9910IOUpdateRateAddress)
	assert.Equal(t, 0x07, AD9910FreqTuningWordAddress)
	assert.Equal(t, 0x16, AD9910RAMAddress)
}

func TestAD9910TestSuite(t *testing.T) {
	suite.Run(t, new(AD9910TestSuite))
}
