package binary

// HasBit returns true if bit is set at position p else false.
func HasBit(b byte, p uint) bool {
	return b&(1<<p) > 0
}

// SetBit returns a copy of byte b with a bit set at position p.
func SetBit(b byte, p uint) byte {
	return b | (1 << p)
}

// UnsetBit returns a copy of byte b with no bit set at position p.
func UnsetBit(b byte, p uint) byte {
	return b &^ (1 << p)
}

// ReadBits returns the bits between position and position+length.
func ReadBits(b byte, position uint, length uint) byte {
	return (b << (8 - position - length)) >> length
}

// WriteBits returns a copy of byte b with bits d set between position
// and position+length.
func WriteBits(b byte, position uint, length uint, d byte) byte {
	m := byte((0xff << (8 - position - length)) >> length)

	return (b &^ m) | (d << position)
}
