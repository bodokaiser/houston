package ad99xx

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrequencyToFTW(t *testing.T) {
	assert.Equal(t, uint32(0x556aaaab), FrequencyToFTW(122.88e6, 41e6))
}

func TestAmplitudeToASF(t *testing.T) {
	assert.Equal(t, uint16(1<<14)-1, AmplitudeToASF(1.0))
	assert.Equal(t, uint16(1<<13), AmplitudeToASF(0.5))
	assert.Equal(t, uint16(0), AmplitudeToASF(0.0))
}

func TestPhaseToPOW(t *testing.T) {
	assert.Equal(t, uint16(0), PhaseToPOW(0*2*math.Pi))
	assert.Equal(t, uint16(0), PhaseToPOW(1*2*math.Pi))
	assert.Equal(t, uint16(1<<15), PhaseToPOW(1*math.Pi))
}
