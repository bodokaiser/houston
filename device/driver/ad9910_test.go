package driver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAD9910Constants(t *testing.T) {
	assert.Equal(t, 0x00, AD9910CtrlFunc1)
	assert.Equal(t, 0x04, AD9910IOUpdateRate)
	assert.Equal(t, 0x07, AD9910FreqTuningWord)
	assert.Equal(t, 0x16, AD9910RAM)
}
