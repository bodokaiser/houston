package device

// Register represents a device register.
type Register struct {
	address  byte
	defaults []byte
}

// Address returns the register address byte.
func (r *Register) Address() byte {
	return r.address
}

// Length returns the byte length of the register.
func (r *Register) Length() int {
	return len(r.defaults)
}

// Defaults returns the default bytes for the register.
func (r *Register) Defaults() []byte {
	return r.defaults
}
