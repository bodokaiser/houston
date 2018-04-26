package ad9910

import (
	"encoding/binary"
)

type RampLimit [8]byte

func NewRampLimit() RampLimit {
	return [8]byte{}
}

func (r *RampLimit) UpperFTW() uint32 {
	return binary.BigEndian.Uint32(r[0:4])
}

func (r *RampLimit) SetUpperFTW(x uint32) {
	binary.BigEndian.PutUint32(r[0:4], x)
}

func (r *RampLimit) UpperPOW() uint16 {
	return binary.BigEndian.Uint16(r[0:2])
}

func (r *RampLimit) SetUpperPOW(x uint16) {
	binary.BigEndian.PutUint32(r[0:4], uint32(x)<<16)
}

func (r *RampLimit) UpperASF() uint16 {
	return binary.BigEndian.Uint16(r[0:2]) >> 2
}

func (r *RampLimit) SetUpperASF(x uint16) {
	binary.BigEndian.PutUint32(r[0:4], uint32(x)<<18)
}

func (r *RampLimit) LowerFTW() uint32 {
	return binary.BigEndian.Uint32(r[4:8])
}

func (r *RampLimit) SetLowerFTW(x uint32) {
	binary.BigEndian.PutUint32(r[4:8], x)
}

func (r *RampLimit) LowerPOW() uint16 {
	return binary.BigEndian.Uint16(r[4:6])
}

func (r *RampLimit) SetLowerPOW(x uint16) {
	binary.BigEndian.PutUint32(r[4:8], uint32(x)<<16)
}

func (r *RampLimit) LowerASF() uint16 {
	return binary.BigEndian.Uint16(r[4:6]) >> 2
}

func (r *RampLimit) SetLowerASF(x uint16) {
	binary.BigEndian.PutUint32(r[4:8], uint32(x)<<18)
}

type RampStep [8]byte

func NewRampStep() RampStep {
	return [8]byte{}
}

func (r *RampStep) DecrStepSize() uint32 {
	return binary.BigEndian.Uint32(r[:4])
}

func (r *RampStep) SetDecrStepSize(x uint32) {
	binary.BigEndian.PutUint32(r[:4], x)
}

func (r *RampStep) IncrStepSize() uint32 {
	return binary.BigEndian.Uint32(r[4:])
}

func (r *RampStep) SetIncrStepSize(x uint32) {
	binary.BigEndian.PutUint32(r[4:], x)
}

type RampRate [4]byte

func NewRampRate() RampRate {
	return [4]byte{}
}

func (r *RampRate) NegSlopeRate() uint16 {
	return binary.BigEndian.Uint16(r[:2])
}

func (r *RampRate) SetNegSlopeRate(x uint16) {
	binary.BigEndian.PutUint16(r[:2], x)
}

func (r *RampRate) PosSlopeRate() uint16 {
	return binary.BigEndian.Uint16(r[2:])
}

func (r *RampRate) SetPosSlopeRate(x uint16) {
	binary.BigEndian.PutUint16(r[2:], x)
}
