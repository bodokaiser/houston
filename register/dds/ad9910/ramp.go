package ad9910

import (
	"encoding/binary"
)

// RampLimit mirrors the digital ramp limit register.
type RampLimit [8]byte

// NewRampLimit returns an initialized RampLimit.
func NewRampLimit() RampLimit {
	return [8]byte{}
}

// UpperFTW returns the upper frequency tuning word.
func (r *RampLimit) UpperFTW() uint32 {
	return binary.BigEndian.Uint32(r[0:4])
}

// SetUpperFTW sets the upper frequency tuning word.
func (r *RampLimit) SetUpperFTW(x uint32) {
	binary.BigEndian.PutUint32(r[0:4], x)
}

// UpperPOW returns the upper phase offset word.
func (r *RampLimit) UpperPOW() uint16 {
	return binary.BigEndian.Uint16(r[0:2])
}

// SetUpperPOW sets the upper phase offset word.
func (r *RampLimit) SetUpperPOW(x uint16) {
	binary.BigEndian.PutUint32(r[0:4], uint32(x)<<16)
}

// UpperASF returns the upper amplitude scale factor.
func (r *RampLimit) UpperASF() uint16 {
	return binary.BigEndian.Uint16(r[0:2]) >> 2
}

// SetUpperASF sets the upper amplitude scale factor.
func (r *RampLimit) SetUpperASF(x uint16) {
	binary.BigEndian.PutUint32(r[0:4], uint32(x)<<18)
}

// LowerFTW returns the lower frequency tuning word.
func (r *RampLimit) LowerFTW() uint32 {
	return binary.BigEndian.Uint32(r[4:8])
}

// SetLowerFTW sets the lower frequency tuning word.
func (r *RampLimit) SetLowerFTW(x uint32) {
	binary.BigEndian.PutUint32(r[4:8], x)
}

// LowerPOW returns the lower phase offset word.
func (r *RampLimit) LowerPOW() uint16 {
	return binary.BigEndian.Uint16(r[4:6])
}

// SetLowerPOW sets the lower phase offset word.
func (r *RampLimit) SetLowerPOW(x uint16) {
	binary.BigEndian.PutUint32(r[4:8], uint32(x)<<16)
}

// LowerASF returns the lower amplitude scale factor.
func (r *RampLimit) LowerASF() uint16 {
	return binary.BigEndian.Uint16(r[4:6]) >> 2
}

// SetLowerASF sets the lower amplitude scale factor.
func (r *RampLimit) SetLowerASF(x uint16) {
	binary.BigEndian.PutUint32(r[4:8], uint32(x)<<18)
}

// RampStep defines the step size of the digital ramp.
type RampStep [8]byte

// NewRampStep returns an initialized RampStep.
func NewRampStep() RampStep {
	return [8]byte{}
}

// DecrStepSize returns the decrement step size.
func (r *RampStep) DecrStepSize() uint32 {
	return binary.BigEndian.Uint32(r[:4])
}

// SetDecrStepSize sets the decrement step size.
func (r *RampStep) SetDecrStepSize(x uint32) {
	binary.BigEndian.PutUint32(r[:4], x)
}

// IncrStepSize returns the increment step size.
func (r *RampStep) IncrStepSize() uint32 {
	return binary.BigEndian.Uint32(r[4:])
}

// SetIncrStepSize sets the increment step size.
func (r *RampStep) SetIncrStepSize(x uint32) {
	binary.BigEndian.PutUint32(r[4:], x)
}

// RampRate mirrors the digital ramp rate register.
type RampRate [4]byte

// NewRampRate returns an initialized RampRate.
func NewRampRate() RampRate {
	return [4]byte{}
}

// NegSlopeRate returns the negative slope rate.
func (r *RampRate) NegSlopeRate() uint16 {
	return binary.BigEndian.Uint16(r[:2])
}

// SetNegSlopeRate sets the negative slope rate.
func (r *RampRate) SetNegSlopeRate(x uint16) {
	binary.BigEndian.PutUint16(r[:2], x)
}

// PosSlopeRate returns the positive slope rate.
func (r *RampRate) PosSlopeRate() uint16 {
	return binary.BigEndian.Uint16(r[2:])
}

// SetPosSlopeRate sets the positive slope rate.
func (r *RampRate) SetPosSlopeRate(x uint16) {
	binary.BigEndian.PutUint16(r[2:], x)
}
