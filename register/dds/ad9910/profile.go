package ad9910

import (
	"encoding/binary"
)

const (
	flagNoDwellHigh  = 1 << 5
	flagZeroCrossing = 1 << 3
)

// STProfile mirrors a singletone profile.
type STProfile [8]byte

// NewSTProfile returns an initialized STProfile.
//
// Default values are actually only for STProfile0.
func NewSTProfile() STProfile {
	return [8]byte{0x08, 0xb5}
}

// AmplScaleFactor returns the amplitude scale factor set in profile.
func (r *STProfile) AmplScaleFactor() uint16 {
	return (binary.BigEndian.Uint16(r[:2]) << 2) >> 2
}

// SetAmplScaleFactor sets the amplitude scale factor to profile.
func (r *STProfile) SetAmplScaleFactor(x uint16) {
	if x > 1<<14 {
		panic("amplitude scale factor not 14 bit")
	}

	binary.BigEndian.PutUint16(r[:2], x)
}

// PhaseOffsetWord returns the phase offset word set in profile.
func (r *STProfile) PhaseOffsetWord() uint16 {
	return binary.BigEndian.Uint16(r[2:4])
}

// SetPhaseOffsetWord sets the phase offset word to profile.
func (r *STProfile) SetPhaseOffsetWord(x uint16) {
	binary.BigEndian.PutUint16(r[2:4], x)
}

// FreqTuningWord returns the frequency tuning word in profile.
func (r *STProfile) FreqTuningWord() uint32 {
	return binary.BigEndian.Uint32(r[4:])
}

// SetFreqTuningWord sets the frequency tuning word to profile.
func (r *STProfile) SetFreqTuningWord(x uint32) {
	binary.BigEndian.PutUint32(r[4:], x)
}

// RAMProfile mirrors a RAM profile.
type RAMProfile [8]byte

// NewRAMProfile returns an initialized RAMProfile.
func NewRAMProfile() RAMProfile {
	return [8]byte{}
}

// AddrStepRate returns the address step rate defined in profile.
func (r *RAMProfile) AddrStepRate() uint16 {
	return binary.BigEndian.Uint16(r[1:3])
}

// SetAddrStepRate sets the address step rate to profile.
func (r *RAMProfile) SetAddrStepRate(x uint16) {
	binary.BigEndian.PutUint16(r[1:3], x)
}

// WaveformStartAddr returns the start address to read from RAM.
func (r *RAMProfile) WaveformStartAddr() uint16 {
	return binary.BigEndian.Uint16(r[5:7]) >> 6
}

// SetWaveformStartAddr sets the start address to RAM.
func (r *RAMProfile) SetWaveformStartAddr(x uint16) {
	binary.BigEndian.PutUint16(r[5:7], x<<6)
}

// WaveformEndAddr returns the end address to read from RAM.
func (r *RAMProfile) WaveformEndAddr() uint16 {
	return binary.BigEndian.Uint16(r[3:5]) >> 6
}

// SetWaveformEndAddr sets the end address to RAM.
func (r *RAMProfile) SetWaveformEndAddr(x uint16) {
	binary.BigEndian.PutUint16(r[3:5], x<<6)
}

// RAMControlMode defines the possible RAM modes.
type RAMControlMode uint8

// Available RAM modes (see datasheet).
const (
	RAMControlModeDirectSwitch RAMControlMode = iota
	RAMControlModeRampUp
	RAMControlModeBiRampUp
	RAMControlModeContBiRampUp
	RAMControlModeContRecirculate
)

// RAMControlMode returns the configured RAMControlMode.
func (r *RAMProfile) RAMControlMode() RAMControlMode {
	switch (r[7] << 5) >> 5 {
	case 0x00:
		fallthrough
	case 0x05:
		fallthrough
	case 0x06:
		fallthrough
	case 0x07:
		return RAMControlModeDirectSwitch
	case 0x01:
		return RAMControlModeRampUp
	case 0x02:
		return RAMControlModeBiRampUp
	case 0x03:
		return RAMControlModeContBiRampUp
	case 0x04:
		return RAMControlModeContRecirculate
	}

	panic("invalid RAM control mode")
}

// SetRAMControlMode configures RAM mode to RAMControlMode.
func (r *RAMProfile) SetRAMControlMode(x RAMControlMode) {
	r[7] &= ^byte(0x7)

	v := byte(x)

	if v&1 > 0 {
		r[7] |= 1
	}
	if v&(1<<1) > 0 {
		r[7] |= 1 << 1
	}
	if v&(1<<2) > 0 {
		r[7] |= 1 << 2
	}
}

// NoDwellHigh returns true if no dwell high flag is set in profile.
func (r *RAMProfile) NoDwellHigh() bool {
	return r[7]&flagNoDwellHigh > 0
}

// SetNoDwellHigh sets no dwell high flag to profile if true.
func (r *RAMProfile) SetNoDwellHigh(x bool) {
	r[7] &= ^byte(flagNoDwellHigh)

	if x {
		r[7] |= flagNoDwellHigh
	}
}

// ZeroCrossing returns true if zero crossing flag is set in profile.
func (r *RAMProfile) ZeroCrossing() bool {
	return r[7]&flagZeroCrossing > 0
}

// SetZeroCrossing sets zero crossing flag to profile if true.
func (r *RAMProfile) SetZeroCrossing(x bool) {
	r[7] &= ^byte(flagZeroCrossing)

	if x {
		r[7] |= flagZeroCrossing
	}
}
