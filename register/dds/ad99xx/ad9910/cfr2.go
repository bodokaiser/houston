package ad9910

const (
	rampEnableFlag      = 1 << 3
	rampNoDwellLowFlag  = 1 << 1
	rampNoDwellHighFlag = 1 << 2

	pdataEnableFlag      = 1 << 4
	pdataClockInvertFlag = 1 << 2
	pdataClockEnableFlag = 1 << 3
	txEnableInvertFlag   = 1 << 1

	syncClockEnableFlag             = 1 << 6
	syncTimingValidationDisableFlag = 1 << 5

	stAmplScaleEnableFlag      = 1 << 0
	readEffectiveFTWFlag       = 1 << 0
	intIOUpdateEnableFlag      = 1 << 7
	dataAssemblerLastValueFlag = 1 << 6
	matchedLatencyEnableFlag   = 1 << 7
)

type CFR2 []byte

func NewCFR2() CFR2 {
	return []byte{0x00, 0x40, 0x08, 0x20}
}

func (r CFR2) STAmplScaleEnabled() bool {
	return r[0]&stAmplScaleEnableFlag > 0
}

func (r CFR2) SetSTAmplScaleEnabled(x bool) {
	r[0] &= ^byte(stAmplScaleEnableFlag)

	if x {
		r[0] |= stAmplScaleEnableFlag
	}
}

func (r CFR2) SyncClockEnabled() bool {
	return r[1]&syncClockEnableFlag > 0
}

func (r CFR2) SetSyncClockEnabled(x bool) {
	r[1] &= ^byte(syncClockEnableFlag)

	if x {
		r[1] |= syncClockEnableFlag
	}
}

type RampDest uint8

const (
	RampDestFrequency RampDest = iota
	RampDestPhase
	RampDestAmplitude
)

func (r CFR2) RampDest() RampDest {
	return RampDest((r[1] << 2) >> 6)
}

func (r CFR2) SetRampDest(x RampDest) {
	r[1] = byte(x) << 4
}

func (r CFR2) RampEnabled() bool {
	return r[1]&rampEnableFlag > 0
}

func (r CFR2) SetRampEnabled(x bool) {
	r[1] &= ^byte(rampEnableFlag)

	if x {
		r[1] |= rampEnableFlag
	}
}

func (r CFR2) RampNoDwellLow() bool {
	return r[1]&rampNoDwellLowFlag > 0
}

func (r CFR2) SetRampNoDwellLow(x bool) {
	r[1] &= ^byte(rampNoDwellLowFlag)

	if x {
		r[1] |= rampNoDwellLowFlag
	}
}

func (r CFR2) RampNoDwellHigh() bool {
	return r[1]&rampNoDwellHighFlag > 0
}

func (r CFR2) SetRampNoDwellHigh(x bool) {
	r[1] &= ^byte(rampNoDwellHighFlag)

	if x {
		r[1] |= rampNoDwellHighFlag
	}
}

func (r CFR2) SyncTimingValidationDisabled() bool {
	return r[3]&syncTimingValidationDisableFlag > 0
}

func (r CFR2) SetSyncTimingValidationDisabled(x bool) {
	r[3] &= ^byte(syncTimingValidationDisableFlag)

	if x {
		r[3] |= syncTimingValidationDisableFlag
	}
}
