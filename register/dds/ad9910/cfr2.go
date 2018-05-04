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

// CFR2 mirrors the CFR2 register.
type CFR2 [4]byte

// NewCFR2 returns an initialized CFR2.
func NewCFR2() CFR2 {
	return [4]byte{0x00, 0x40, 0x08, 0x20}
}

// STAmplScaleEnabled returns true if amplitude scale from singletone profile
// flag is set.
func (r *CFR2) STAmplScaleEnabled() bool {
	return r[0]&flagSTAmplScaleEnable > 0
}

// SetSTAmplScaleEnabled sets the amplitude scale from singletone profile
// flag to be set if true.
func (r *CFR2) SetSTAmplScaleEnabled(x bool) {
	r[0] &= ^byte(flagSTAmplScaleEnable)

	if x {
		r[0] |= flagSTAmplScaleEnable
	}
}

// SyncClockEnabled returns true if sync clock enabled flag is set.
func (r *CFR2) SyncClockEnabled() bool {
	return r[1]&flagSyncClockEnable > 0
}

// SetSyncClockEnabled sets sync clock enabled flag to be set if true.
func (r *CFR2) SetSyncClockEnabled(x bool) {
	r[1] &= ^byte(flagSyncClockEnable)

	if x {
		r[1] |= flagSyncClockEnable
	}
}

// RampDest specifies the parameter to be controlled by the digital ramp (sweep).
type RampDest uint8

// Available destinations for digital ramp.
const (
	RampDestFrequency RampDest = iota
	RampDestPhase
	RampDestAmplitude
)

// RampDest returns the configured RampDest.
func (r *CFR2) RampDest() RampDest {
	switch (r[1] << 2) >> 6 {
	case 0x00:
		return RampDestFrequency
	case 0x01:
		return RampDestPhase
	}

	return RampDestAmplitude
}

// SetRampDest configures RampDest as ramp destination.
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

// RampEnabled returns true if ramp enabled flag is set.
func (r *CFR2) RampEnabled() bool {
	return r[1]&flagRampEnable > 0
}

// SetRampEnabled sets the ramp enabled flag if true.
func (r *CFR2) SetRampEnabled(x bool) {
	r[1] &= ^byte(flagRampEnable)

	if x {
		r[1] |= flagRampEnable
	}
}

// RampNoDwellLow returns true if ramp no dwell low flag is set.
func (r *CFR2) RampNoDwellLow() bool {
	return r[1]&flagRampNoDwellLow > 0
}

// SetRampNoDwellLow sets the ramp no dwell low flag if true.
func (r *CFR2) SetRampNoDwellLow(x bool) {
	r[1] &= ^byte(flagRampNoDwellLow)

	if x {
		r[1] |= flagRampNoDwellLow
	}
}

// RampNoDwellHigh returns true if ramp no dwell high flag is set.
func (r *CFR2) RampNoDwellHigh() bool {
	return r[1]&flagRampNoDwellHigh > 0
}

// SetRampNoDwellHigh sets the ramp no dwell high flag if true.
func (r *CFR2) SetRampNoDwellHigh(x bool) {
	r[1] &= ^byte(flagRampNoDwellHigh)

	if x {
		r[1] |= flagRampNoDwellHigh
	}
}

// SyncTimingValidationDisabled returns true if sync timing validation disabled
// flag is set.
func (r *CFR2) SyncTimingValidationDisabled() bool {
	return r[3]&flagSyncTimingValidDisable > 0
}

// SetSyncTimingValidationDisabled sets the sync timing validation disabled
// flag if true.
func (r *CFR2) SetSyncTimingValidationDisabled(x bool) {
	r[3] &= ^byte(flagSyncTimingValidDisable)

	if x {
		r[3] |= flagSyncTimingValidDisable
	}
}
