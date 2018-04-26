package ad9910

const (
	flagRAMEnable = 1 << 7
	flagOSKEnable = 1 << 1
	flagOSKManual = 1 << 7
	flagOSKAuto   = 1 << 0

	flagClearPhaseAcc     = 1 << 3
	flagClearRampAcc      = 1 << 4
	flagAutoClearPhaseAcc = 1 << 5
	flagAutoClearRampAcc  = 1 << 6

	flagLoadARR = 1 << 2
	flagLoadLRR = 1 << 7

	flagInverseSinc = 1 << 6
	flagSineOutput  = 1 << 0
	flagLSBFirst    = 1 << 0
	flagSPI3Wire    = 1 << 1

	flagPowerDownExtCtrl  = 1 << 2
	flagPowerDownAuxDAC   = 1 << 3
	flagPowerDownRefInput = 1 << 4
	flagPowerDownDAC      = 1 << 5
	flagPowerDownCore     = 1 << 6
)

type CFR1 [4]byte

func NewCFR1() CFR1 {
	return [4]byte{}
}

func (r *CFR1) RAMEnabled() bool {
	return r[0]&flagRAMEnable > 0
}

func (r *CFR1) SetRAMEnabled(x bool) {
	r[0] &= ^byte(flagRAMEnable)

	if x {
		r[0] |= flagRAMEnable
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
	return r[2]&flagOSKEnable > 0
}

func (r *CFR1) SetOSKEnabled(x bool) {
	r[2] &= ^byte(flagOSKEnable)

	if x {
		r[2] |= flagOSKEnable
	}
}

func (r *CFR1) OSKManual() bool {
	return r[1]&flagOSKManual > 0
}

func (r *CFR1) SetOSKManual(x bool) {
	r[1] &= ^byte(flagOSKManual)

	if x {
		r[1] |= flagOSKManual
	}
}

func (r *CFR1) OSKAuto() bool {
	return r[2]&flagOSKAuto > 0
}

func (r *CFR1) SetOSKAuto(x bool) {
	r[2] &= ^byte(flagOSKAuto)

	if x {
		r[2] |= flagOSKAuto
	}
}

func (r *CFR1) SDIOInputOnly() bool {
	return r[3]&flagSPI3Wire > 0
}

func (r *CFR1) SetSDIOInputOnly(x bool) {
	r[3] &= ^byte(flagSPI3Wire)

	if x {
		r[3] |= flagSPI3Wire
	}
}
