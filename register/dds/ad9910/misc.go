package ad9910

import (
	"encoding/binary"
	"math"
)

// AuxDAC mirrors the auxilliary DAC register.
type AuxDAC [4]byte

// NewAuxDAC returns an initialized AuxDAC.
func NewAuxDAC() AuxDAC {
	return [4]byte{0x00, 0x00, 0x00, 0x7f}
}

// IOUpdateRate mirrors the I/O update rate register.
type IOUpdateRate [4]byte

// NewIOUpdateRate returns an initialized IOUpdateRate.
func NewIOUpdateRate() IOUpdateRate {
	return [4]byte{0xff, 0xff, 0xff, 0xff}
}

// FTW mirrors the frequency tuning word register.
type FTW [4]byte

// NewFTW returns an initialized FTW.
func NewFTW() FTW {
	return [4]byte{}
}

// FreqTuningWord returns the frequency tuning word.
func (r *FTW) FreqTuningWord() uint32 {
	return binary.BigEndian.Uint32(r[:])
}

// SetFreqTuningWord sets the frequency tuning word.
func (r *FTW) SetFreqTuningWord(x uint32) {
	binary.BigEndian.PutUint32(r[:], x)
}

// POW mirrors the phase offset word register.
type POW [2]byte

// NewPOW returns an initialized POW.
func NewPOW() POW {
	return [2]byte{}
}

// PhaseOffsetWord returns the phase offset word.
func (r *POW) PhaseOffsetWord() uint16 {
	return binary.BigEndian.Uint16(r[:])
}

// SetPhaseOffsetWord sets the phase offset word.
func (r *POW) SetPhaseOffsetWord(x uint16) {
	binary.BigEndian.PutUint16(r[:], x)
}

// ASF mirrors the amplitude scale factor register.
type ASF [4]byte

// NewASF returns an initialized ASF.
func NewASF() ASF {
	return [4]byte{}
}

// AmplRampRate returns the amplitude ramp rate used with OSK.
func (r *ASF) AmplRampRate() uint16 {
	return binary.BigEndian.Uint16(r[:2])
}

// SetAmplRampRate sets the amplitude ramp rate used with OSK.
func (r *ASF) SetAmplRampRate(x uint16) {
	binary.BigEndian.PutUint16(r[:2], x)
}

// AmplScaleFactor returns the amplitude scale factor.
func (r *ASF) AmplScaleFactor() uint16 {
	return binary.BigEndian.Uint16(r[2:]) >> 2
}

// SetAmplScaleFactor sets the amplitude scale factor.
func (r *ASF) SetAmplScaleFactor(x uint16) {
	if x > math.MaxUint16>>2 {
		panic("amplitude scale factor not 14 bit")
	}

	binary.BigEndian.PutUint16(r[2:], x<<2)
}

// AmplStepSize defines amplitude step sizes to be used with OSK.
type AmplStepSize uint8

// Available amplitude step sizes.
const (
	AmplStepSize1 AmplStepSize = iota
	AmplStepSize2 AmplStepSize = iota
	AmplStepSize4 AmplStepSize = iota
	AmplStepSize8 AmplStepSize = iota
)

// AmplStepSize returns the configured AmplStepSize.
func (r *ASF) AmplStepSize() AmplStepSize {
	return AmplStepSize((r[3] << 6) >> 2)
}

// SetAmplStepSize configures the given AmplStepSize.
func (r *ASF) SetAmplStepSize(x AmplStepSize) {
	r[3] |= byte(x)
}

// RAM mirrors the RAM word register.
//
// A single RAM word equals one playback value.
type RAM [4]byte

// NewRAM returns an initialized RAM.
func NewRAM() RAM {
	return RAM{}
}

// AmplScaleFactor returns the amplitude scale factor set in RAM.
func (r *RAM) AmplScaleFactor() uint16 {
	return binary.BigEndian.Uint16(r[0:2]) >> 2
}

// SetAmplScaleFactor sets the amplitude scale factor to RAM.
func (r *RAM) SetAmplScaleFactor(x uint16) {
	binary.BigEndian.PutUint16(r[0:2], x<<2)
}

// PhaseOffsetWord returns the phase offset word set in RAM.
func (r *RAM) PhaseOffsetWord() uint16 {
	return binary.BigEndian.Uint16(r[0:2])
}

// SetPhaseOffsetWord sets the phase offset word to RAM.
func (r *RAM) SetPhaseOffsetWord(x uint16) {
	binary.BigEndian.PutUint16(r[0:2], x)
}

// FreqTuningWord returns the frequency tuning word set in RAM.
func (r *RAM) FreqTuningWord() uint32 {
	return binary.BigEndian.Uint32(r[0:4])
}

// SetFreqTuningWord sets the frequency tuning word to RAM.
func (r *RAM) SetFreqTuningWord(x uint32) {
	binary.BigEndian.PutUint32(r[0:4], x)
}
