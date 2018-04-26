package ad9910

const (
	flagRampEnable      = 1 << 3
	flagRampNoDwellLow  = 1 << 1
	flagRampNoDwellHigh = 1 << 2

	flagPDataEnable      = 1 << 4
	flagPDataClockInvert = 1 << 2
	flagPDataClockEnable = 1 << 3
	flagTXEnableInvert   = 1 << 1

	flagSyncClockEnable        = 1 << 6
	flagSyncTimingValidDisable = 1 << 5

	flagSTAmplScaleEnable      = 1 << 0
	flagReadEffectiveFTW       = 1 << 0
	flagIntUpdateEnable        = 1 << 7
	flagDataAssemblerLastValue = 1 << 6
	flagMatchedLatencyEnable   = 1 << 7
)

type CFR2 [4]byte

func NewCFR2() CFR2 {
	return [4]byte{0x00, 0x40, 0x08, 0x20}
}

func (r *CFR2) STAmplScaleEnabled() bool {
	return r[0]&flagSTAmplScaleEnable > 0
}

func (r *CFR2) SetSTAmplScaleEnabled(x bool) {
	r[0] &= ^byte(flagSTAmplScaleEnable)

	if x {
		r[0] |= flagSTAmplScaleEnable
	}
}

func (r *CFR2) SyncClockEnabled() bool {
	return r[1]&flagSyncClockEnable > 0
}

func (r *CFR2) SetSyncClockEnabled(x bool) {
	r[1] &= ^byte(flagSyncClockEnable)

	if x {
		r[1] |= flagSyncClockEnable
	}
}

type RampDest uint8

const (
	RampDestFrequency RampDest = iota
	RampDestPhase
	RampDestAmplitude
)

func (r *CFR2) RampDest() RampDest {
	switch (r[1] << 2) >> 6 {
	case 0x00:
		return RampDestFrequency
	case 0x01:
		return RampDestPhase
	}

	return RampDestAmplitude
}

func (r *CFR2) SetRampDest(x RampDest) {
	r[1] &= ^byte(0x30)

	v := byte(x)

	if v&1 > 0 {
		r[1] |= 1 << 4
	}
	if v&(1<<1) > 0 {
		r[1] |= 1 << 5
	}
}

func (r *CFR2) RampEnabled() bool {
	return r[1]&flagRampEnable > 0
}

func (r *CFR2) SetRampEnabled(x bool) {
	r[1] &= ^byte(flagRampEnable)

	if x {
		r[1] |= flagRampEnable
	}
}

func (r *CFR2) RampNoDwellLow() bool {
	return r[1]&flagRampNoDwellLow > 0
}

func (r *CFR2) SetRampNoDwellLow(x bool) {
	r[1] &= ^byte(flagRampNoDwellLow)

	if x {
		r[1] |= flagRampNoDwellLow
	}
}

func (r *CFR2) RampNoDwellHigh() bool {
	return r[1]&flagRampNoDwellHigh > 0
}

func (r *CFR2) SetRampNoDwellHigh(x bool) {
	r[1] &= ^byte(flagRampNoDwellHigh)

	if x {
		r[1] |= flagRampNoDwellHigh
	}
}

func (r *CFR2) SyncTimingValidationDisabled() bool {
	return r[3]&flagSyncTimingValidDisable > 0
}

func (r *CFR2) SetSyncTimingValidationDisabled(x bool) {
	r[3] &= ^byte(flagSyncTimingValidDisable)

	if x {
		r[3] |= flagSyncTimingValidDisable
	}
}
