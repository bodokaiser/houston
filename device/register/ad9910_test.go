package register

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAD9910Constants(t *testing.T) {
	assert.Equal(t, 0x00, AD9910CtrlFunc1Address)
	assert.Equal(t, 0x04, AD9910IOUpdateRateAddress)
	assert.Equal(t, 0x07, AD9910FreqTuningWordAddress)
	assert.Equal(t, 0x16, AD9910RAMAddress)
}

func TestAD9910LSBFirst(t *testing.T) {
	r := new(AD9910)

	r.CtrlFunc1[0] = 0x00
	assert.False(t, r.LSBFirst())

	r.CtrlFunc1[0] = 0x01
	assert.True(t, r.LSBFirst())
}

func TestAD9910SetLSBFirst(t *testing.T) {
	r := new(AD9910)
	r.CtrlFunc1 = AD9910CtrlFunc1Default

	r.SetLSBFirst(true)
	assert.Equal(t, r.CtrlFunc1, 0x01)

	r.SetLSBFirst(false)
	assert.Equal(t, r.CtrlFunc1, 0x00)
}
