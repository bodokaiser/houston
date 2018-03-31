package binary

// HasBit returns true if bit is set at position p else false.
func HasBit(b byte, p uint) bool {
	return b&(1<<p) > 0
}

// SetBit returns a copy of byte with a bit set at position p.
func SetBit(b byte, p uint) byte {
	return b | (1 << p)
}

// UnsetBit returns a copy of byte with no bit set at position p.
func UnsetBit(b byte, p uint) byte {
	return b &^ (1 << p)
}
