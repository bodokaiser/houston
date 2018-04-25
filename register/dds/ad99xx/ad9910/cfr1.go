package ad9910

const (
	ramEnableFlag = 1 << 7
	oskEnableFlag = 1 << 1
	oskManualFlag = 1 << 7
	oskAutoFlag   = 1 << 0

	clearPhaseAccFlag     = 1 << 3
	clearDRampAccFlag     = 1 << 4
	autoclearPhaseAccFlag = 1 << 5
	autoclearDRampAccFlag = 1 << 6

	loadARRFlag = 1 << 2
	loadLRRFlag = 1 << 7

	inverseSincFlag = 1 << 6
	sineOutputFlag  = 1 << 0
	lsbFirstFlag    = 1 << 0
	sdioInputFlag   = 1 << 1

	powerDownExtCtrlFlag  = 1 << 2
	powerDownAuxDACFlag   = 1 << 3
	powerDownRefInputFlag = 1 << 4
	powerDownDACFlag      = 1 << 5
	powerDownCoreFlag     = 1 << 6
)

type CFR1 [4]byte

func NewCFR1() CFR1 {
	return [4]byte{}
}

func (r *CFR1) RAMEnabled() bool {
	return r[0]&ramEnableFlag > 0
}

func (r *CFR1) SetRAMEnabled(x bool) {
	r[0] &= ^byte(ramEnableFlag)

	if x {
		r[0] |= ramEnableFlag
	}
}

type RAMDest uint8

const (
	RAMDestFrequency RAMDest = iota
	RAMDestPhase
	RAMDestAmplitude
	RAMDestPolar
)

func (r *CFR1) RAMDest() RAMDest {
	return RAMDest((r[0] << 1) >> 6)
}

func (r *CFR1) SetRAMDest(x RAMDest) {
	r[0] &= ^byte(0x30)

	v := byte(x)

	if v&1 > 0 {
		r[0] |= 1 << 5
	}
	if v&(1<<1) > 0 {
		r[0] |= 1 << 6
	}
}

func (r *CFR1) OSKEnabled() bool {
	return r[2]&oskEnableFlag > 0
}

func (r *CFR1) SetOSKEnabled(x bool) {
	r[2] &= ^byte(oskEnableFlag)

	if x {
		r[2] |= oskEnableFlag
	}
}

func (r *CFR1) OSKManual() bool {
	return r[1]&oskManualFlag > 0
}

func (r *CFR1) SetOSKManual(x bool) {
	r[1] &= ^byte(oskManualFlag)

	if x {
		r[1] |= oskManualFlag
	}
}

func (r *CFR1) OSKAuto() bool {
	return r[2]&oskAutoFlag > 0
}

func (r *CFR1) SetOSKAuto(x bool) {
	r[2] &= ^byte(oskAutoFlag)

	if x {
		r[2] |= oskAutoFlag
	}
}

func (r *CFR1) SDIOInputOnly() bool {
	return r[3]&sdioInputFlag > 0
}

func (r *CFR1) SetSDIOInputOnly(x bool) {
	r[3] &= ^byte(sdioInputFlag)

	if x {
		r[3] |= sdioInputFlag
	}
}
