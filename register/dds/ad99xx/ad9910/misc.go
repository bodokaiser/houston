package ad9910

import (
	"encoding/binary"
)

type AuxDAC [4]byte

func NewAuxDAC() AuxDAC {
	return [4]byte{0x00, 0x00, 0x00, 0x7f}
}

type IOUpdateRate [4]byte

func NewIOUpdateRate() IOUpdateRate {
	return [4]byte{0xff, 0xff, 0xff, 0xff}
}

type FTW [4]byte

func NewFTW() FTW {
	return [4]byte{}
}

func (r *FTW) FreqTuningWord() uint32 {
	return binary.BigEndian.Uint32(r[:])
}

func (r *FTW) SetFreqTuningWord(x uint32) {
	binary.BigEndian.PutUint32(r[:], x)
}

type POW [2]byte

func NewPOW() POW {
	return [2]byte{}
}

func (r *POW) PhaseOffsetWord() uint16 {
	return binary.BigEndian.Uint16(r[:])
}

func (r *POW) SetPhaseOffsetWord(x uint16) {
	binary.BigEndian.PutUint16(r[:], x)
}

type ASF [4]byte

func NewASF() ASF {
	return [4]byte{}
}

func (r *ASF) AmplRampRate() uint16 {
	return binary.BigEndian.Uint16(r[:2])
}

func (r *ASF) SetAmplRampRate(x uint16) {
	binary.BigEndian.PutUint16(r[:2], x)
}

func (r *ASF) AmplScaleFactor() uint16 {
	return binary.BigEndian.Uint16(r[2:]) >> 2
}

func (r *ASF) SetAmplScaleFactor(x uint16) {
	if x > 1<<14 {
		panic("amplitude scale factor not 14 bit")
	}

	binary.BigEndian.PutUint16(r[2:], x<<2)
}

type AmplStepSize uint8

const (
	AmplStepSize1 AmplStepSize = iota
	AmplStepSize2 AmplStepSize = iota
	AmplStepSize4 AmplStepSize = iota
	AmplStepSize8 AmplStepSize = iota
)

func (r *ASF) AmplStepSize() AmplStepSize {
	return AmplStepSize((r[3] << 6) >> 2)
}

func (r *ASF) SetAmplStepSize(x AmplStepSize) {
	r[3] |= byte(x)
}
