package ad9910

import "encoding/binary"

const (
	AddrCFR1         = 0x00
	AddrCFR2         = 0x01
	AddrCFR3         = 0x02
	AddrAuxDAC       = 0x03
	AddrIOUpdateRate = 0x04
	AddrFTW          = 0x07
	AddrPOW          = 0x08
	AddrASF          = 0x09
	AddrMultiChip    = 0x0a
	AddrRampLimit    = 0x0b
	AddrRampStep     = 0x0c
	AddrRampRate     = 0x0d
	AddrProfile0     = 0x0e
	AddrProfile1     = 0x0f
	AddrProfile2     = 0x10
	AddrProfile3     = 0x11
	AddrProfile4     = 0x12
	AddrProfile5     = 0x13
	AddrProfile6     = 0x14
	AddrProfile7     = 0x15
	AddrRAM          = 0x16
)

type AuxDAC []byte

func NewAuxDAC() AuxDAC {
	return []byte{0x00, 0x00, 0x00, 0x7f}
}

type IOUpdateRate []byte

func NewIOUpdateRate() IOUpdateRate {
	return []byte{0xff, 0xff, 0xff, 0xff}
}

func (r IOUpdateRate) Marshal() []byte {
	return append([]byte{0x04}, []byte(r)...)
}

type FTW []byte

func NewFTW() FTW {
	return make([]byte, 8)
}

func (r FTW) FreqTuningWord() uint32 {
	return binary.BigEndian.Uint32(r)
}

func (r FTW) SetFreqTuningWord(x uint32) {
	binary.BigEndian.PutUint32(r, x)
}

type POW []byte

func NewPOW() POW {
	return make([]byte, 2)
}

func (r POW) PhaseOffsetWord() uint16 {
	return binary.BigEndian.Uint16(r)
}

func (r POW) SetPhaseOffsetWord(x uint16) {
	binary.BigEndian.PutUint16(r, x)
}

type ASF []byte

func NewASF() ASF {
	return make([]byte, 4)
}

func (r ASF) AmplRampRate() uint16 {
	return binary.BigEndian.Uint16(r[:2])
}

func (r ASF) SetAmplRampRate(x uint16) {
	binary.BigEndian.PutUint16(r[:2], x)
}

func (r ASF) AmplScaleFactor() uint16 {
	return binary.BigEndian.Uint16(r[2:]) >> 2
}

func (r ASF) SetAmplScaleFactor(x uint16) {
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

func (r ASF) AmplStepSize() AmplStepSize {
	return AmplStepSize((r[3] << 6) >> 2)
}

func (r ASF) SetAmplStepSize(x AmplStepSize) {
	r[3] |= byte(x)
}
