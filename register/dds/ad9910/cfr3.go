package ad9910

const (
	flagPLLEnable = 1 << 0
	flagPFDReset  = 1 << 2

	flagRefClockDivReset  = 1 << 6
	flagRefClockDivBypass = 1 << 7
)

// CFR3 mirrors the CFR3 register.
type CFR3 [4]byte

// NewCFR3 returns an initialized CFR3.
func NewCFR3() CFR3 {
	return [4]byte{0x1f, 0x3f, 0x40, 0x00}
}

// PLLEnabled returns true if PLL enabled flag is set.
func (r *CFR3) PLLEnabled() bool {
	return r[2]&flagPLLEnable > 0
}

// SetPLLEnabled sets PLL enabled flag if true.
func (r *CFR3) SetPLLEnabled(x bool) {
	r[2] &= ^byte(flagPLLEnable)

	if x {
		r[2] |= flagPLLEnable
	}
}

// VCORange defines frequency ranges to be used for PLL.
type VCORange uint8

// Available VCO frequency ranges (see datasheet).
const (
	VCORange0 VCORange = iota
	VCORange1
	VCORange2
	VCORange3
	VCORange4
	VCORange5
	VCORangeByPassed
)

// VCORange returns the configured VCORange.
func (r *CFR3) VCORange() VCORange {
	return VCORange((r[0] << 5) >> 5)
}

// SetVCORange configures the given VCORange.
func (r *CFR3) SetVCORange(x VCORange) {
	r[0] = byte(x)
}

// Divider returns the ratio between system and reference clock.
func (r *CFR3) Divider() uint8 {
	return uint8(r[3] >> 1)
}

// SetDivider sets the ratio between system and reference clock.
func (r *CFR3) SetDivider(x uint8) {
	r[3] = byte(x << 1)
}
