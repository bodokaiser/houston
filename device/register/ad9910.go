package register

import "github.com/bodokaiser/beagle/encoding/binary"

// Default values for the AD9910 register from the datasheet
// reference. If a register is not explicitly named then
// default value can be assumed to consist of zero bytes.
var (
	AD9910CtrlFunc1Default = [4]byte{
		0x00, 0x40, 0x08, 0x20}
	AD9910CtrlFunc2Default = [4]byte{
		0x1f, 0x3f, 0x40, 0x00}
	AD9910CtrlFunc3Default = [4]byte{
		0x00, 0x00, 0x00, 0x7f}
	AD9910AuxDACCtrlDefault = [4]byte{
		0xff, 0xff, 0xff, 0xff}
	AD9910StProfile0Default = [8]byte{
		0x08, 0xb5, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
)

// Register addresses for the AD9910 with registers
// 0x05 and 0x06 (see datasheet).
const (
	AD9910CtrlFunc1Address = iota
	AD9910CtrlFunc2Address
	AD9910CtrlFunc3Address
	AD9910AuxDACCtrlAddress
	AD9910IOUpdateRateAddress
	_
	_
	AD9910FreqTuningWordAddress
	AD9910PhaseOffsetWordAddress
	AD9910AmplScaleFactorAddress
	AD9910MultichipSyncAddress
	AD9910DigitalRampLimitAddress
	AD9910DigitalRampStepSizeAddress
	AD9910DigitalRampRateAddress
	AD9910Profile0Address
	AD9910Profile1Address
	AD9910Profile2Address
	AD9910Profile3Address
	AD9910Profile4Address
	AD9910Profile5Address
	AD9910Profile6Address
	AD9910Profile7Address
	AD9910RAMAddress
)

// AD9910 reproduces the AD9910 registers and provides
// comprehensible operations close to the datasheet reference
// to interact with the register values.
type AD9910 struct {
	CtrlFunc1           [4]byte
	CtrlFunc2           [4]byte
	CtrlFunc3           [4]byte
	AuxDACCtrl          [4]byte
	IOUpdateRate        [4]byte
	FreqTuningWord      [4]byte
	PhaseOffsetWord     [2]byte
	AmplScaleFactor     [4]byte
	MultichipSync       [4]byte
	DigitalRampLimit    [8]byte
	DigitalRampStepSize [8]byte
	DigitalRampRate     [4]byte
	STProfile0          [8]byte
	STProfile1          [8]byte
	RAMProfile0         [8]byte
	RAMProfile1         [8]byte
	RAMMemory           [4]byte
}

// LSBFirst returns true if SPI byte order is configured to be
// LSB and false if SPI byte order is MSB.
func (r *AD9910) LSBFirst() bool {
	return binary.HasBit(r.CtrlFunc1[0], 0)
}

// SetLSBFirst configures the SPI byte order to be LSB on
// true and MSB on false.
func (r *AD9910) SetLSBFirst(active bool) {
	r.CtrlFunc1[0] = binary.UnsetBit(r.CtrlFunc1[0], 0)

	if active {
		r.CtrlFunc1[0] = binary.SetBit(r.CtrlFunc1[0], 0)
	}
}

// SDIOInputOnly returns true if SPI uses 3-wire mode and
// false if SDIO pin is used bidirectional.
func (r *AD9910) SDIOInputOnly() bool {
	return binary.HasBit(r.CtrlFunc1[0], 1)
}

// SetSDIOInputOnly configures SPI to use 3-wire on true and
// 2-wire on false.
func (r *AD9910) SetSDIOInputOnly(active bool) {
	r.CtrlFunc1[0] = binary.UnsetBit(r.CtrlFunc1[0], 1)

	if active {
		r.CtrlFunc1[0] = binary.SetBit(r.CtrlFunc1[0], 1)
	}
}

// AD9910PowerDownMode defines how the AD9910 responds to
// an external power down request.
type AD9910PowerDownMode int

// AD9910 can be configured to go in full or fast recovery
// power down.
const (
	PowerDownFull AD9910PowerDownMode = iota
	PowerDownFastRecovery
)

// ExtPowerDownCtrl returns the configured power down
// behaviour (see AD9910PowerDownModee).
func (r *AD9910) ExtPowerDownCtrl() AD9910PowerDownMode {
	if binary.HasBit(r.CtrlFunc1[0], 3) {
		return PowerDownFastRecovery
	}
	return PowerDownFull
}

// SetExtPowerDownCtrl configures the power down
// behaviour (see AD9910PowerDownModee).
func (r *AD9910) SetExtPowerDownCtrl(m AD9910PowerDownMode) {
	r.CtrlFunc1[0] = binary.UnsetBit(r.CtrlFunc1[0], 3)

	if PowerDownFastRecovery == m {
		r.CtrlFunc1[0] = binary.SetBit(r.CtrlFunc1[0], 3)
	}
}

// AuxiliaryDACPowerDown returns true if auxiliary DAC is
// configured to be powered down.
func (r *AD9910) AuxiliaryDACPowerDown() bool {
	return binary.HasBit(r.CtrlFunc1[0], 4)
}

// SetAuxiliaryDACPowerDown configures the auxiliary DAC to
// be powered down if active is true.
func (r *AD9910) SetAuxiliaryDACPowerDown(active bool) {
	r.CtrlFunc1[0] = binary.UnsetBit(r.CtrlFunc1[0], 4)

	if active {
		r.CtrlFunc1[0] = binary.SetBit(r.CtrlFunc1[0], 4)
	}
}

// RefClockInputPowerDown returns true if external reference
// clock input pin is configured to be powered down.
func (r *AD9910) RefClockInputPowerDown() bool {
	return binary.HasBit(r.CtrlFunc1[0], 5)
}

// SetRefClockInputPowerDown configures the external reference
// clock input pin to be configured powered down if
// active is true.
func (r *AD9910) SetRefClockInputPowerDown(active bool) {
	r.CtrlFunc1[0] = binary.UnsetBit(r.CtrlFunc1[0], 5)

	if active {
		r.CtrlFunc1[0] = binary.SetBit(r.CtrlFunc1[0], 5)
	}
}

// DACPowerDown returns true if main DAC is configured to be
// powered down.
func (r *AD9910) DACPowerDown() bool {
	return binary.HasBit(r.CtrlFunc1[0], 6)
}

// SetDACPowerDown configures the main DAC to be configured
// powered down if active is true.
func (r *AD9910) SetDACPowerDown(active bool) {
	r.CtrlFunc1[0] = binary.UnsetBit(r.CtrlFunc1[0], 6)

	if active {
		r.CtrlFunc1[0] = binary.SetBit(r.CtrlFunc1[0], 6)
	}
}

// DigitalPowerDown returns true if core is configured to be
// powered down.
func (r *AD9910) DigitalPowerDown() bool {
	return binary.HasBit(r.CtrlFunc1[0], 7)
}

// SetDigitalPowerDown configures the core to be powered down
// if active is true.
func (r *AD9910) SetDigitalPowerDown(active bool) {
	r.CtrlFunc1[0] = binary.UnsetBit(r.CtrlFunc1[0], 7)

	if active {
		r.CtrlFunc1[0] = binary.SetBit(r.CtrlFunc1[0], 7)
	}
}

// SelectAutoOSK returns true if automatic OSK is enabled.
func (r *AD9910) SelectAutoOSK() bool {
	return binary.HasBit(r.CtrlFunc1[1], 0)
}

// SetSelectAutoOSK configures automatic OSK to be enabled
// if active is true.
func (r *AD9910) SetSelectAutoOSK(active bool) {
	r.CtrlFunc1[1] = binary.UnsetBit(r.CtrlFunc1[1], 0)

	if active {
		r.CtrlFunc1[1] = binary.SetBit(r.CtrlFunc1[1], 0)
	}
}

// OSKEnable returns true if output shift keying (OSK) is
// enabled.
func (r *AD9910) OSKEnable() bool {
	return binary.HasBit(r.CtrlFunc1[1], 1)
}

// SetOSKEnable configures output shift keying (OSK) to be
// enabled if active is true.
func (r *AD9910) SetOSKEnable(active bool) {
	r.CtrlFunc1[1] = binary.UnsetBit(r.CtrlFunc1[1], 1)

	if active {
		r.CtrlFunc1[1] = binary.SetBit(r.CtrlFunc1[1], 1)
	}
}

// LoadARR is over my paygrade.
func (r *AD9910) LoadARR() bool {
	return binary.HasBit(r.CtrlFunc1[1], 2)
}

// SetLoadARR is over my paygrade.
func (r *AD9910) SetLoadARR(active bool) {
	r.CtrlFunc1[1] = binary.UnsetBit(r.CtrlFunc1[1], 2)

	if active {
		r.CtrlFunc1[1] = binary.SetBit(r.CtrlFunc1[1], 2)
	}
}

// ClearPhaseAccumulator returns true if async, static reset
// of the phase accumulator is configured.
func (r *AD9910) ClearPhaseAccumulator() bool {
	return binary.HasBit(r.CtrlFunc1[1], 3)
}

// SetClearPhaseAccumulator configures a static reset of the
// phase accumulator if active is true.
func (r *AD9910) SetClearPhaseAccumulator(active bool) {
	r.CtrlFunc1[1] = binary.UnsetBit(r.CtrlFunc1[1], 3)

	if active {
		r.CtrlFunc1[1] = binary.SetBit(r.CtrlFunc1[1], 3)
	}
}

// ClearDigitalRampAccumulator returns true if async, static
// reset of the digital ramp accumulator is configured.
func (r *AD9910) ClearDigitalRampAccumulator() bool {
	return binary.HasBit(r.CtrlFunc1[1], 4)
}

// SetClearDigitalRampAccumulator configures a static reset of the
// digital ramp accumulator if active is true.
func (r *AD9910) SetClearDigitalRampAccumulator(active bool) {
	r.CtrlFunc1[1] = binary.UnsetBit(r.CtrlFunc1[1], 4)

	if active {
		r.CtrlFunc1[1] = binary.SetBit(r.CtrlFunc1[1], 4)
	}
}

// AutoClearPhaseAccumulator returns true if phase accumulator
// is configured to be reset on IOUpdate or profile change.
func (r *AD9910) AutoClearPhaseAccumulator() bool {
	return binary.HasBit(r.CtrlFunc1[1], 5)
}

// SetAutoClearPhaseAccumulator configures a sync reset on
// IOUpdate or profile change of the phase accumulator if
// active is true.
func (r *AD9910) SetAutoClearPhaseAccumulator(active bool) {
	r.CtrlFunc1[1] = binary.UnsetBit(r.CtrlFunc1[1], 5)

	if active {
		r.CtrlFunc1[1] = binary.SetBit(r.CtrlFunc1[1], 5)
	}
}

// AutoClearDigitalRampAccumulator returns true if
// digital ramp accumulator is configured to be reset on
// IOUpdate or profile change.
func (r *AD9910) AutoClearDigitalRampAccumulator() bool {
	return binary.HasBit(r.CtrlFunc1[1], 5)
}

// SetAutoClearDigitalRampAccumulator configures a sync reset on
// IOUpdate or profile change of the digital ramp accumulator if
// active is true.
func (r *AD9910) SetAutoClearDigitalRampAccumulator(active bool) {
	r.CtrlFunc1[1] = binary.UnsetBit(r.CtrlFunc1[1], 5)

	if active {
		r.CtrlFunc1[1] = binary.SetBit(r.CtrlFunc1[1], 5)
	}
}

// LoadLRR is over my paygrade.
func (r *AD9910) LoadLRR() bool {
	return binary.HasBit(r.CtrlFunc1[1], 6)
}

// SetLoadLRR is over my paygrade.
func (r *AD9910) SetLoadLRR(active bool) {
	r.CtrlFunc1[1] = binary.UnsetBit(r.CtrlFunc1[1], 6)

	if active {
		r.CtrlFunc1[1] = binary.SetBit(r.CtrlFunc1[1], 6)
	}
}

// SelectSineOutput returns true if sine output is selected
// as output and false if cosine output is configured.
func (r *AD9910) SelectSineOutput() bool {
	return binary.HasBit(r.CtrlFunc1[1], 7)
}

// SetSelectSineOutput configures output to be sine if
// active is true.
func (r *AD9910) SetSelectSineOutput(active bool) {
	r.CtrlFunc1[1] = binary.UnsetBit(r.CtrlFunc1[1], 7)

	if active {
		r.CtrlFunc1[1] = binary.SetBit(r.CtrlFunc1[1], 7)
	}
}

// AD9910RAMProfileMode defines the composition of the
// different RAM profiles.
type AD9910RAMProfileMode int

// Supported RAM profile compositions. BurstN can be read that
// profile 0 to N will be executed once while ContN
// continuously executes profiles 0 to N.
const (
	AD9910ProfileModeDisabled AD9910RAMProfileMode = iota
	AD9910ProfileModeBurst1
	AD9910ProfileModeBurst2
	AD9910ProfileModeBurst3
	AD9910ProfileModeBurst4
	AD9910ProfileModeBurst5
	AD9910ProfileModeBurst6
	AD9910ProfileModeBurst7
	AD9910ProfileModeCont1
	AD9910ProfileModeCont2
	AD9910ProfileModeCont3
	AD9910ProfileModeCont4
	AD9910ProfileModeCont5
	AD9910ProfileModeCont6
	AD9910ProfileModeCont7
)

// IntProfileCtrl returns the configured AD9910RAMProfile mode.
func (r *AD9910) IntProfileCtrl() AD9910RAMProfileMode {
	b := binary.ReadBits(r.CtrlFunc1[2], 1, 4)

	return AD9910RAMProfileMode(b)
}

// SetIntProfileCtrl configures the AD9910 to use the given
// AD9910RAMProfileMode m.
func (r *AD9910) SetIntProfileCtrl(m AD9910RAMProfileMode) {
	binary.WriteBits(r.CtrlFunc1[2], 1, 4, byte(m))
}

// InverseSincFilterEnable returns true if the inverse sinc
// filter is configured to be enabled.
func (r *AD9910) InverseSincFilterEnable() bool {
	return binary.HasBit(r.CtrlFunc1[2], 6)
}

// SetInverseSincFilterEnable configures the inverse sinc filter
// to be enabled if active is true.
func (r *AD9910) SetInverseSincFilterEnable(active bool) {
	r.CtrlFunc1[2] = binary.UnsetBit(r.CtrlFunc1[2], 6)

	if active {
		r.CtrlFunc1[2] = binary.SetBit(r.CtrlFunc1[2], 6)
	}
}

// ManualOSKExtCtrl return true if external OSK pin is
// configured to be enabled.
func (r *AD9910) ManualOSKExtCtrl() bool {
	return binary.HasBit(r.CtrlFunc1[2], 7)
}

// SetManualOSKExtCtrl configures the external OSK pin to be
// configured enabled if active is true.
func (r *AD9910) SetManualOSKExtCtrl(active bool) {
	r.CtrlFunc1[2] = binary.UnsetBit(r.CtrlFunc1[2], 7)

	if active {
		r.CtrlFunc1[2] = binary.SetBit(r.CtrlFunc1[2], 7)
	}
}

// AD9910RAMPlaybackDest defines what parameters to control
// by RAM playback.
type AD9910RAMPlaybackDest int

// See Table 12 in the datasheet for details.
const (
	AD9910RAMPlaybackDestFreq AD9910RAMPlaybackDest = iota
	AD9910RAMPlaybackDestPhase
	AD9910RAMPlaybackDestAmpl
	AD9910RAMPlaybackDestPolar
)

// RAMPlaybackDest returns the configured playback destination
// in RAM mode. See AD9910RAMPPlaybackDest.
func (r *AD9910) RAMPlaybackDest() AD9910RAMPlaybackDest {
	b := binary.ReadBits(r.CtrlFunc1[3], 5, 2)

	return AD9910RAMPlaybackDest(b)
}

// SetRAMPlaybackDest configures the playback destination
// in RAM mode. See AD9910RAMPPlaybackDest.
func (r *AD9910) SetRAMPlaybackDest(d AD9910RAMPlaybackDest) {
	binary.WriteBits(r.CtrlFunc1[3], 5, 2, byte(d))
}

// RAMEnable returns true if RAM functionality (for playback
// operation) is configured to be enabled.
func (r *AD9910) RAMEnable() bool {
	return binary.HasBit(r.CtrlFunc1[3], 7)
}

// SetRAMEnable configures RAM functionality to be enabled
// if active is true.
func (r *AD9910) SetRAMEnable(active bool) {
	r.CtrlFunc1[3] = binary.UnsetBit(r.CtrlFunc1[3], 7)

	if active {
		r.CtrlFunc1[3] = binary.SetBit(r.CtrlFunc1[3], 7)
	}
}

// ParallelPortEnable returns true if parallel port functionality is configured
// to be enabled.
func (r *AD9910) ParallelPortEnable() bool {
	return binary.HasBit(r.CtrlFunc2[0], 4)
}

// SetParallelPortEnable configures the parallel port functionality to be
// enabled if true.
func (r *AD9910) SetParallelPortEnable(active bool) {
	r.CtrlFunc2[0] = binary.UnsetBit(r.CtrlFunc2[0], 4)

	if active {
		r.CtrlFunc2[0] = binary.SetBit(r.CtrlFunc2[0], 4)
	}
}

// SyncTimingValidationDisable returns false if SYNC_SMP_ERR pin is configured
// to detect syncronization pulse sampling errors.
func (r *AD9910) SyncTimingValidationDisable() bool {
	return binary.HasBit(r.CtrlFunc2[0], 5)
}

// SetSyncTimingValidationDisable configures SYNC_SMP_ERR pin to detect
// syncronization pulse sampling errors.
func (r *AD9910) SetSyncTimingValidationDisable(active bool) {
	r.CtrlFunc2[0] = binary.UnsetBit(r.CtrlFunc2[0], 5)

	if active {
		r.CtrlFunc2[0] = binary.SetBit(r.CtrlFunc2[0], 5)
	}
}

// DataAssemblerHoldLastValue relates to some parallel port communication.
func (r *AD9910) DataAssemblerHoldLastValue() bool {
	return binary.HasBit(r.CtrlFunc2[0], 6)
}

// SetDataAssemblerHoldLastValue relates to some parallel port communication.
func (r *AD9910) SetDataAssemblerHoldLastValue(active bool) {
	r.CtrlFunc2[0] = binary.UnsetBit(r.CtrlFunc2[0], 6)

	if active {
		r.CtrlFunc2[0] = binary.SetBit(r.CtrlFunc2[0], 6)
	}
}

// MatchedLatencyEnable returns true if parameter changes affect the output
// signal in the order made.
func (r *AD9910) MatchedLatencyEnable() bool {
	return binary.HasBit(r.CtrlFunc2[0], 7)
}

// SetMatchedLatencyEnable configures the output signal to be affected by
// multiple parameter changes at once if true.
func (r *AD9910) SetMatchedLatencyEnable(active bool) {
	r.CtrlFunc2[0] = binary.UnsetBit(r.CtrlFunc2[0], 7)

	if active {
		r.CtrlFunc2[0] = binary.SetBit(r.CtrlFunc2[0], 7)
	}
}

// TxEnableInvert relates to the parallel port communication.
func (r *AD9910) TxEnableInvert() bool {
	return binary.HasBit(r.CtrlFunc2[1], 0)
}

// SetTxEnableInvert relates to the parallel port communication.
func (r *AD9910) SetTxEnableInvert(active bool) {
	r.CtrlFunc2[1] = binary.UnsetBit(r.CtrlFunc2[1], 0)

	if active {
		r.CtrlFunc2[1] = binary.SetBit(r.CtrlFunc2[1], 0)
	}
}

// ParallelDataClockInvert relates to the parallel port communication.
func (r *AD9910) ParallelDataClockInvert() bool {
	return binary.HasBit(r.CtrlFunc2[1], 2)
}

// SetParallelDataClockInvert relates to the parallel port communication.
func (r *AD9910) SetParallelDataClockInvert(active bool) {
	r.CtrlFunc2[1] = binary.UnsetBit(r.CtrlFunc2[1], 2)

	if active {
		r.CtrlFunc2[1] = binary.SetBit(r.CtrlFunc2[1], 2)
	}
}

// ParallelDataClockEnable relates to the parallel port communication.
func (r *AD9910) ParallelDataClockEnable() bool {
	return binary.HasBit(r.CtrlFunc2[1], 3)
}

// SetParallelDataClockEnable relates to the parallel port communication.
func (r *AD9910) SetParallelDataClockEnable(active bool) {
	r.CtrlFunc2[1] = binary.UnsetBit(r.CtrlFunc2[1], 3)

	if active {
		r.CtrlFunc2[1] = binary.SetBit(r.CtrlFunc2[1], 3)
	}
}

// ReadEffectiveFreqTuningWord returns true if the AD9910 is configured to
// reply with the actual measured frequency if FTW is requested.
func (r *AD9910) ReadEffectiveFreqTuningWord() bool {
	return binary.HasBit(r.CtrlFunc2[2], 0)
}

// SetReadEffectiveFreqTuningWord configures the AD9910 to reply with the
// actual frequency instead of the register value if true.
func (r *AD9910) SetReadEffectiveFreqTuningWord(active bool) {
	r.CtrlFunc2[2] = binary.UnsetBit(r.CtrlFunc2[2], 0)

	if active {
		r.CtrlFunc2[2] = binary.SetBit(r.CtrlFunc2[2], 0)
	}
}

// DigitalRampNoDwellLow returns true if the digital ramp is configured to
// directly skip back to the configured lower limit instead of ramping back.
func (r *AD9910) DigitalRampNoDwellLow() bool {
	return binary.HasBit(r.CtrlFunc2[2], 1)
}

// SetDigitalRampNoDwellLow configures the digital ramp to skip instead of
// ramping back to the opposing limit if true.
func (r *AD9910) SetDigitalRampNoDwellLow(active bool) {
	r.CtrlFunc2[2] = binary.UnsetBit(r.CtrlFunc2[2], 1)

	if active {
		r.CtrlFunc2[2] = binary.SetBit(r.CtrlFunc2[2], 1)
	}
}

// DigitalRampEnable returns true if the digital ramp is configured to be
// enabled.
func (r *AD9910) DigitalRampEnable() bool {
	return binary.HasBit(r.CtrlFunc2[2], 2)
}

// SetDigitalRampEnable configures the digital ramp to be enabled if
// active is true.
func (r *AD9910) SetDigitalRampEnable(active bool) {
	r.CtrlFunc2[2] = binary.UnsetBit(r.CtrlFunc2[2], 2)

	if active {
		r.CtrlFunc2[2] = binary.SetBit(r.CtrlFunc2[2], 2)
	}
}

// SyncClockEnable returns true if the AD9910 is configured to generate
// a clock signal to syncronize SPI communication.
func (r *AD9910) SyncClockEnable() bool {
	return binary.HasBit(r.CtrlFunc2[2], 6)
}

// SetSyncClockEnable configures the AD9910 to generate a clock signal
// for SPI communication if active is true.
func (r *AD9910) SetSyncClockEnable(active bool) {
	r.CtrlFunc2[2] = binary.UnsetBit(r.CtrlFunc2[2], 6)

	if active {
		r.CtrlFunc2[2] = binary.SetBit(r.CtrlFunc2[2], 6)
	}
}

// IntIOUpdateActive returns true if the SPI communication is configured
// to be syncronized with an internally generated I/O update signal.
func (r *AD9910) IntIOUpdateActive() bool {
	return binary.HasBit(r.CtrlFunc2[2], 7)
}

// SetIntIOUpdateActive configures the SPI communication to be syncronized
// with an internally generated I/O update signal if active is true.
func (r *AD9910) SetIntIOUpdateActive(active bool) {
	r.CtrlFunc2[2] = binary.UnsetBit(r.CtrlFunc2[2], 7)

	if active {
		r.CtrlFunc2[2] = binary.SetBit(r.CtrlFunc2[2], 7)
	}
}

// AmplScaleFromSTProfileEnable returns true if amplitude from single tone
// profile is configured to be scaled by amplitude scalar factor.
func (r *AD9910) AmplScaleFromSTProfileEnable() bool {
	return binary.HasBit(r.CtrlFunc2[3], 0)
}

// SetAmplScaleFromSTProfileEnable configures the amplitude from single tone
// profile to be configured to be scaled by amplitude scalar factor if
// active is true.
func (r *AD9910) SetAmplScaleFromSTProfileEnable(active bool) {
	r.CtrlFunc2[3] = binary.UnsetBit(r.CtrlFunc2[3], 0)

	if active {
		r.CtrlFunc2[3] = binary.SetBit(r.CtrlFunc2[3], 0)
	}
}

// PhaseLockedLoopEnable returns true if reference clock phase locked loop
// is configured to be enabled.
func (r *AD9910) PhaseLockedLoopEnable() bool {
	return binary.HasBit(r.CtrlFunc3[0], 0)
}

// SetPhaseLockedLoopEnable configures the reference clock phase locked loop
// (PLL) to be configured enabled if active is true.
func (r *AD9910) SetPhaseLockedLoopEnable(active bool) {
	r.CtrlFunc3[0] = binary.UnsetBit(r.CtrlFunc3[0], 0)

	if active {
		r.CtrlFunc3[0] = binary.SetBit(r.CtrlFunc3[0], 0)
	}
}

// PhaseFreqDetectorReset returns true if phase frequency detector (PFD) is
// configured to be disabled.
func (r *AD9910) PhaseFreqDetectorReset() bool {
	return binary.HasBit(r.CtrlFunc3[0], 2)
}

// SetPhaseFreqDetectorReset configures the phase frequency detector (PFD)
// to be enabled if active is true.
func (r *AD9910) SetPhaseFreqDetectorReset(active bool) {
	r.CtrlFunc3[0] = binary.UnsetBit(r.CtrlFunc3[0], 2)

	if active {
		r.CtrlFunc3[0] = binary.SetBit(r.CtrlFunc3[0], 2)
	}
}

// RefClockInputDividerReset returns true if reference clock input divider
// operates normally.
func (r *AD9910) RefClockInputDividerReset() bool {
	return binary.HasBit(r.CtrlFunc3[0], 6)
}

// SetRefClockInputDividerReset configures the reference clock input divider
// to operate normally if active is true and be reset else.
func (r *AD9910) SetRefClockInputDividerReset(active bool) {
	r.CtrlFunc3[0] = binary.UnsetBit(r.CtrlFunc3[0], 6)

	if active {
		r.CtrlFunc3[0] = binary.SetBit(r.CtrlFunc3[0], 6)
	}
}

// RefClockInputDividerBypass returns true if reference clock input divider
// is configured to be bypassed.
func (r *AD9910) RefClockInputDividerBypass() bool {
	return binary.HasBit(r.CtrlFunc3[0], 7)
}

// SetRefClockInputDividerBypass configures the reference clock input divider
// to be configured bypassed if active is true.
func (r *AD9910) SetRefClockInputDividerBypass(active bool) {
	r.CtrlFunc3[0] = binary.UnsetBit(r.CtrlFunc3[0], 7)

	if active {
		r.CtrlFunc3[0] = binary.SetBit(r.CtrlFunc3[0], 7)
	}
}
