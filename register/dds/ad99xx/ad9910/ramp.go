package ad9910

import (
	"encoding/binary"
)

type RampLimit [8]byte

func NewRampLimit() RampLimit {
	return [8]byte{}
}

func (r *RampLimit) UpperLimit() uint32 {
	return binary.BigEndian.Uint32(r[:4])
}

func (r *RampLimit) SetUpperLimit(x uint32) {
	binary.BigEndian.PutUint32(r[:4], x)
}

func (r *RampLimit) LowerLimit() uint32 {
	return binary.BigEndian.Uint32(r[4:])
}

func (r *RampLimit) SetLowerLimit(x uint32) {
	binary.BigEndian.PutUint32(r[4:], x)
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
