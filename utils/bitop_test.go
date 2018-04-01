package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasBit(t *testing.T) {
	b := []byte{0x00, 0xff, 0x82}

	for i := uint(0); i < 8; i++ {
		assert.False(t, HasBit(b[0], i),
			fmt.Sprintf("bit %d not detected unset in 0x00", i))
	}
	for i := uint(0); i < 8; i++ {
		assert.True(t, HasBit(b[1], i),
			fmt.Sprintf("bit %d not detected set 0xff", i))
	}

	assert.False(t, HasBit(b[2], 0), "bit 0 not detected unset in 0x82")
	assert.True(t, HasBit(b[2], 1), "bit 1 not detected set in 0x82")
	assert.False(t, HasBit(b[2], 2), "bit 2 not detected unset in 0x82")
	assert.False(t, HasBit(b[2], 3), "bit 3 not detected unset in 0x82")
	assert.False(t, HasBit(b[2], 4), "bit 4 not detected unset in 0x82")
	assert.False(t, HasBit(b[2], 5), "bit 5 not detected unset in 0x82")
	assert.False(t, HasBit(b[2], 6), "bit 6 not detected unset in 0x82")
	assert.True(t, HasBit(b[2], 7), "bit 7 not detected set in 0x82")
}

func TestSetBit(t *testing.T) {
	for i := uint(0); i < 8; i++ {
		assert.Equal(t, byte(1<<i), SetBit(byte(0x00), i),
			fmt.Sprintf("bit %d has not been set in 0x00", i))
	}

	b := byte(0x82)
	b = SetBit(b, 0)
	b = SetBit(b, 6)
	assert.Equal(t, byte(0xc3), b,
		fmt.Sprintf("0x82 should be 0xc3 but was %#x", b))
}

func TestUnsetBit(t *testing.T) {
	b := byte(0xff)

	for i := uint(0); i < 8; i++ {
		b = UnsetBit(b, i)
	}

	assert.Equal(t, byte(0x00), b)
}

func TestReadBits(t *testing.T) {
	b := byte(0x2c)

	assert.Equal(t, byte(0x0b), ReadBits(b, 2, 4))
}

func TestWriteBits(t *testing.T) {
	b := byte(0x00)

	assert.Equal(t, byte(0x2c), WriteBits(b, 2, 4, 0x0b))
}
