package ad9910

const (
	flagPLLEnable = 1 << 0
	flagPFDReset  = 1 << 2

	flagRefClockDivReset  = 1 << 6
	flagRefClockDivBypass = 1 << 7
)

type CFR3 [4]byte

func NewCFR3() CFR3 {
	return [4]byte{0x1f, 0x3f, 0x40, 0x00}
}

func (r *CFR3) PLLEnabled() bool {
	return r[2]&flagPLLEnable > 0
}

func (r *CFR3) SetPLLEnabled(x bool) {
	r[2] &= ^byte(flagPLLEnable)

	if x {
		r[2] |= flagPLLEnable
	}
}

type VCORange uint8

const (
	VCORange0 VCORange = iota
	VCORange1
	VCORange2
	VCORange3
	VCORange4
	VCORange5
	VCORangeByPassed
)

func (r *CFR3) VCORange() VCORange {
	return VCORange((r[0] << 5) >> 5)
}

func (r *CFR3) SetVCORange(x VCORange) {
	r[0] = byte(x)
}

func (r *CFR3) Divider() uint8 {
	return uint8(r[3] >> 1)
}

func (r *CFR3) SetDivider(x uint8) {
	r[3] = byte(x << 1)
}
