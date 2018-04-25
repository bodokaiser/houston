package ad9910

import (
	"encoding/binary"
)

const (
	noDwellHighFlag  = 1 << 5
	zeroCrossingFlag = 1 << 3

	profileAddrOffset = 0x0e
)

type STProfile [8]byte

func NewSTProfile() STProfile {
	return [8]byte{0x08, 0xb5}
}

func (r *STProfile) AmplScaleFactor() uint16 {
	return (binary.BigEndian.Uint16(r[:2]) << 2) >> 2
}

func (r *STProfile) SetAmplScaleFactor(x uint16) {
	if x > 1<<14 {
		panic("amplitude scale factor not 14 bit")
	}

	binary.BigEndian.PutUint16(r[:2], x)
}

func (r *STProfile) PhaseOffsetWord() uint16 {
	return binary.BigEndian.Uint16(r[2:4])
}

func (r *STProfile) SetPhaseOffsetWord(x uint16) {
	binary.BigEndian.PutUint16(r[2:4], x)
}

func (r *STProfile) FreqTuningWord() uint32 {
	return binary.BigEndian.Uint32(r[4:])
}

func (r *STProfile) SetFreqTuningWord(x uint32) {
	binary.BigEndian.PutUint32(r[4:], x)
}

type RAMProfile [8]byte

func NewRAMProfile() RAMProfile {
	return [8]byte{}
}

func (r *RAMProfile) AddrStepRate() uint16 {
	return binary.BigEndian.Uint16(r[1:3])
}

func (r *RAMProfile) SetAddrStepRate(x uint16) {
	binary.BigEndian.PutUint16(r[1:3], x)
}

func (r *RAMProfile) WaveformStartAddr() uint16 {
	return binary.BigEndian.Uint16(r[5:7]) >> 6
}

func (r *RAMProfile) SetWaveformStartAddr(x uint16) {
	binary.BigEndian.PutUint16(r[5:7], x<<6)
}

func (r *RAMProfile) WaveformEndAddr() uint16 {
	return binary.BigEndian.Uint16(r[3:5]) >> 6
}

func (r *RAMProfile) SetWaveformEndAddr(x uint16) {
	binary.BigEndian.PutUint16(r[3:5], x<<6)
}

type RAMControlMode uint8

const (
	RAMControlModeDirectSwitch RAMControlMode = iota
	RAMControlModeRampUp
	RAMControlModeBiRampUp
	RAMControlModeContBiRampUp
	RAMControlModeContRecirculate
)

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

func (r *RAMProfile) NoDwellHigh() bool {
	return r[7]&noDwellHighFlag > 0
}

func (r *RAMProfile) SetNoDwellHigh(x bool) {
	r[7] &= ^byte(noDwellHighFlag)

	if x {
		r[7] |= noDwellHighFlag
	}
}

func (r *RAMProfile) ZeroCrossing() bool {
	return r[7]&zeroCrossingFlag > 0
}

func (r *RAMProfile) SetZeroCrossing(x bool) {
	r[7] &= ^byte(zeroCrossingFlag)

	if x {
		r[7] |= zeroCrossingFlag
	}
}
