package register

import (
	"encoding/binary"

	"github.com/bodokaiser/beagle/utils"
)

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
	CtrlFunc1Data           [4]byte
	CtrlFunc2Data           [4]byte
	CtrlFunc3Data           [4]byte
	AuxDACCtrlData          [4]byte
	IOUpdateRateData        [4]byte
	FreqTuningWordData      [4]byte
	PhaseOffsetWordData     [2]byte
	AmplScaleFactorData     [4]byte
	MultichipSyncData       [4]byte
	DigitalRampLimitData    [8]byte
	DigitalRampStepSizeData [8]byte
	DigitalRampRateData     [4]byte
	STProfile0Data          [8]byte
	STProfile1Data          [8]byte
	RAMProfile0Data         [8]byte
	RAMProfile1Data         [8]byte
	RAMMemoryData           [4]byte
}

// LSBFirst returns true if SPI byte order is configured to be
// LSB and false if SPI byte order is MSB.
func (r *AD9910) LSBFirst() bool {
	return utils.HasBit(r.CtrlFunc1Data[0], 0)
}

// SetLSBFirst configures the SPI byte order to be LSB on
// true and MSB on false.
func (r *AD9910) SetLSBFirst(active bool) {
	r.CtrlFunc1Data[0] = utils.UnsetBit(r.CtrlFunc1Data[0], 0)

	if active {
		r.CtrlFunc1Data[0] = utils.SetBit(r.CtrlFunc1Data[0], 0)
	}
}

// SDIOInputOnly returns true if SPI uses 3-wire mode and
// false if SDIO pin is used bidirectional.
func (r *AD9910) SDIOInputOnly() bool {
	return utils.HasBit(r.CtrlFunc1Data[0], 1)
}

// SetSDIOInputOnly configures SPI to use 3-wire on true and
// 2-wire on false.
func (r *AD9910) SetSDIOInputOnly(active bool) {
	r.CtrlFunc1Data[0] = utils.UnsetBit(r.CtrlFunc1Data[0], 1)

	if active {
		r.CtrlFunc1Data[0] = utils.SetBit(r.CtrlFunc1Data[0], 1)
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
	if utils.HasBit(r.CtrlFunc1Data[0], 3) {
		return PowerDownFastRecovery
	}
	return PowerDownFull
}

// SetExtPowerDownCtrl configures the power down
// behaviour (see AD9910PowerDownModee).
func (r *AD9910) SetExtPowerDownCtrl(m AD9910PowerDownMode) {
	r.CtrlFunc1Data[0] = utils.UnsetBit(r.CtrlFunc1Data[0], 3)

	if PowerDownFastRecovery == m {
		r.CtrlFunc1Data[0] = utils.SetBit(r.CtrlFunc1Data[0], 3)
	}
}

// AuxiliaryDACPowerDown returns true if auxiliary DAC is
// configured to be powered down.
func (r *AD9910) AuxiliaryDACPowerDown() bool {
	return utils.HasBit(r.CtrlFunc1Data[0], 4)
}

// SetAuxiliaryDACPowerDown configures the auxiliary DAC to
// be powered down if active is true.
func (r *AD9910) SetAuxiliaryDACPowerDown(active bool) {
	r.CtrlFunc1Data[0] = utils.UnsetBit(r.CtrlFunc1Data[0], 4)

	if active {
		r.CtrlFunc1Data[0] = utils.SetBit(r.CtrlFunc1Data[0], 4)
	}
}

// RefClockInputPowerDown returns true if external reference
// clock input pin is configured to be powered down.
func (r *AD9910) RefClockInputPowerDown() bool {
	return utils.HasBit(r.CtrlFunc1Data[0], 5)
}

// SetRefClockInputPowerDown configures the external reference
// clock input pin to be configured powered down if
// active is true.
func (r *AD9910) SetRefClockInputPowerDown(active bool) {
	r.CtrlFunc1Data[0] = utils.UnsetBit(r.CtrlFunc1Data[0], 5)

	if active {
		r.CtrlFunc1Data[0] = utils.SetBit(r.CtrlFunc1Data[0], 5)
	}
}

// DACPowerDown returns true if main DAC is configured to be
// powered down.
func (r *AD9910) DACPowerDown() bool {
	return utils.HasBit(r.CtrlFunc1Data[0], 6)
}

// SetDACPowerDown configures the main DAC to be configured
// powered down if active is true.
func (r *AD9910) SetDACPowerDown(active bool) {
	r.CtrlFunc1Data[0] = utils.UnsetBit(r.CtrlFunc1Data[0], 6)

	if active {
		r.CtrlFunc1Data[0] = utils.SetBit(r.CtrlFunc1Data[0], 6)
	}
}

// DigitalPowerDown returns true if core is configured to be
// powered down.
func (r *AD9910) DigitalPowerDown() bool {
	return utils.HasBit(r.CtrlFunc1Data[0], 7)
}

// SetDigitalPowerDown configures the core to be powered down
// if active is true.
func (r *AD9910) SetDigitalPowerDown(active bool) {
	r.CtrlFunc1Data[0] = utils.UnsetBit(r.CtrlFunc1Data[0], 7)

	if active {
		r.CtrlFunc1Data[0] = utils.SetBit(r.CtrlFunc1Data[0], 7)
	}
}

// SelectAutoOSK returns true if automatic OSK is enabled.
func (r *AD9910) SelectAutoOSK() bool {
	return utils.HasBit(r.CtrlFunc1Data[1], 0)
}

// SetSelectAutoOSK configures automatic OSK to be enabled
// if active is true.
func (r *AD9910) SetSelectAutoOSK(active bool) {
	r.CtrlFunc1Data[1] = utils.UnsetBit(r.CtrlFunc1Data[1], 0)

	if active {
		r.CtrlFunc1Data[1] = utils.SetBit(r.CtrlFunc1Data[1], 0)
	}
}

// OSKEnable returns true if output shift keying (OSK) is
// enabled.
func (r *AD9910) OSKEnable() bool {
	return utils.HasBit(r.CtrlFunc1Data[1], 1)
}

// SetOSKEnable configures output shift keying (OSK) to be
// enabled if active is true.
func (r *AD9910) SetOSKEnable(active bool) {
	r.CtrlFunc1Data[1] = utils.UnsetBit(r.CtrlFunc1Data[1], 1)

	if active {
		r.CtrlFunc1Data[1] = utils.SetBit(r.CtrlFunc1Data[1], 1)
	}
}

// LoadARR is over my paygrade.
func (r *AD9910) LoadARR() bool {
	return utils.HasBit(r.CtrlFunc1Data[1], 2)
}

// SetLoadARR is over my paygrade.
func (r *AD9910) SetLoadARR(active bool) {
	r.CtrlFunc1Data[1] = utils.UnsetBit(r.CtrlFunc1Data[1], 2)

	if active {
		r.CtrlFunc1Data[1] = utils.SetBit(r.CtrlFunc1Data[1], 2)
	}
}

// ClearPhaseAccumulator returns true if async, static reset
// of the phase accumulator is configured.
func (r *AD9910) ClearPhaseAccumulator() bool {
	return utils.HasBit(r.CtrlFunc1Data[1], 3)
}

// SetClearPhaseAccumulator configures a static reset of the
// phase accumulator if active is true.
func (r *AD9910) SetClearPhaseAccumulator(active bool) {
	r.CtrlFunc1Data[1] = utils.UnsetBit(r.CtrlFunc1Data[1], 3)

	if active {
		r.CtrlFunc1Data[1] = utils.SetBit(r.CtrlFunc1Data[1], 3)
	}
}

// ClearDigitalRampAccumulator returns true if async, static
// reset of the digital ramp accumulator is configured.
func (r *AD9910) ClearDigitalRampAccumulator() bool {
	return utils.HasBit(r.CtrlFunc1Data[1], 4)
}

// SetClearDigitalRampAccumulator configures a static reset of the
// digital ramp accumulator if active is true.
func (r *AD9910) SetClearDigitalRampAccumulator(active bool) {
	r.CtrlFunc1Data[1] = utils.UnsetBit(r.CtrlFunc1Data[1], 4)

	if active {
		r.CtrlFunc1Data[1] = utils.SetBit(r.CtrlFunc1Data[1], 4)
	}
}

// AutoClearPhaseAccumulator returns true if phase accumulator
// is configured to be reset on IOUpdate or profile change.
func (r *AD9910) AutoClearPhaseAccumulator() bool {
	return utils.HasBit(r.CtrlFunc1Data[1], 5)
}

// SetAutoClearPhaseAccumulator configures a sync reset on
// IOUpdate or profile change of the phase accumulator if
// active is true.
func (r *AD9910) SetAutoClearPhaseAccumulator(active bool) {
	r.CtrlFunc1Data[1] = utils.UnsetBit(r.CtrlFunc1Data[1], 5)

	if active {
		r.CtrlFunc1Data[1] = utils.SetBit(r.CtrlFunc1Data[1], 5)
	}
}

// AutoClearDigitalRampAccumulator returns true if
// digital ramp accumulator is configured to be reset on
// IOUpdate or profile change.
func (r *AD9910) AutoClearDigitalRampAccumulator() bool {
	return utils.HasBit(r.CtrlFunc1Data[1], 5)
}

// SetAutoClearDigitalRampAccumulator configures a sync reset on
// IOUpdate or profile change of the digital ramp accumulator if
// active is true.
func (r *AD9910) SetAutoClearDigitalRampAccumulator(active bool) {
	r.CtrlFunc1Data[1] = utils.UnsetBit(r.CtrlFunc1Data[1], 5)

	if active {
		r.CtrlFunc1Data[1] = utils.SetBit(r.CtrlFunc1Data[1], 5)
	}
}

// LoadLRR is over my paygrade.
func (r *AD9910) LoadLRR() bool {
	return utils.HasBit(r.CtrlFunc1Data[1], 6)
}

// SetLoadLRR is over my paygrade.
func (r *AD9910) SetLoadLRR(active bool) {
	r.CtrlFunc1Data[1] = utils.UnsetBit(r.CtrlFunc1Data[1], 6)

	if active {
		r.CtrlFunc1Data[1] = utils.SetBit(r.CtrlFunc1Data[1], 6)
	}
}

// SelectSineOutput returns true if sine output is selected
// as output and false if cosine output is configured.
func (r *AD9910) SelectSineOutput() bool {
	return utils.HasBit(r.CtrlFunc1Data[1], 7)
}

// SetSelectSineOutput configures output to be sine if
// active is true.
func (r *AD9910) SetSelectSineOutput(active bool) {
	r.CtrlFunc1Data[1] = utils.UnsetBit(r.CtrlFunc1Data[1], 7)

	if active {
		r.CtrlFunc1Data[1] = utils.SetBit(r.CtrlFunc1Data[1], 7)
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
	b := utils.ReadBits(r.CtrlFunc1Data[2], 1, 4)

	return AD9910RAMProfileMode(b)
}

// SetIntProfileCtrl configures the AD9910 to use the given
// AD9910RAMProfileMode m.
func (r *AD9910) SetIntProfileCtrl(m AD9910RAMProfileMode) {
	utils.WriteBits(r.CtrlFunc1Data[2], 1, 4, byte(m))
}

// InverseSincFilterEnable returns true if the inverse sinc
// filter is configured to be enabled.
func (r *AD9910) InverseSincFilterEnable() bool {
	return utils.HasBit(r.CtrlFunc1Data[2], 6)
}

// SetInverseSincFilterEnable configures the inverse sinc filter
// to be enabled if active is true.
func (r *AD9910) SetInverseSincFilterEnable(active bool) {
	r.CtrlFunc1Data[2] = utils.UnsetBit(r.CtrlFunc1Data[2], 6)

	if active {
		r.CtrlFunc1Data[2] = utils.SetBit(r.CtrlFunc1Data[2], 6)
	}
}

// ManualOSKExtCtrl return true if external OSK pin is
// configured to be enabled.
func (r *AD9910) ManualOSKExtCtrl() bool {
	return utils.HasBit(r.CtrlFunc1Data[2], 7)
}

// SetManualOSKExtCtrl configures the external OSK pin to be
// configured enabled if active is true.
func (r *AD9910) SetManualOSKExtCtrl(active bool) {
	r.CtrlFunc1Data[2] = utils.UnsetBit(r.CtrlFunc1Data[2], 7)

	if active {
		r.CtrlFunc1Data[2] = utils.SetBit(r.CtrlFunc1Data[2], 7)
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
	b := utils.ReadBits(r.CtrlFunc1Data[3], 5, 2)

	return AD9910RAMPlaybackDest(b)
}

// SetRAMPlaybackDest configures the playback destination
// in RAM mode. See AD9910RAMPPlaybackDest.
func (r *AD9910) SetRAMPlaybackDest(d AD9910RAMPlaybackDest) {
	utils.WriteBits(r.CtrlFunc1Data[3], 5, 2, byte(d))
}

// RAMEnable returns true if RAM functionality (for playback
// operation) is configured to be enabled.
func (r *AD9910) RAMEnable() bool {
	return utils.HasBit(r.CtrlFunc1Data[3], 7)
}

// SetRAMEnable configures RAM functionality to be enabled
// if active is true.
func (r *AD9910) SetRAMEnable(active bool) {
	r.CtrlFunc1Data[3] = utils.UnsetBit(r.CtrlFunc1Data[3], 7)

	if active {
		r.CtrlFunc1Data[3] = utils.SetBit(r.CtrlFunc1Data[3], 7)
	}
}

// FreqModGain returns a 4 bit FM gain word which can be used in parallel port
// modulation mode, however the datasheet is a bit unspecific on possible
// values.
func (r *AD9910) FreqModGain() byte {
	return utils.ReadBits(r.CtrlFunc2Data[0], 0, 4)
}

// ParallelPortEnable returns true if parallel port functionality is configured
// to be enabled.
func (r *AD9910) ParallelPortEnable() bool {
	return utils.HasBit(r.CtrlFunc2Data[0], 4)
}

// SetParallelPortEnable configures the parallel port functionality to be
// enabled if true.
func (r *AD9910) SetParallelPortEnable(active bool) {
	r.CtrlFunc2Data[0] = utils.UnsetBit(r.CtrlFunc2Data[0], 4)

	if active {
		r.CtrlFunc2Data[0] = utils.SetBit(r.CtrlFunc2Data[0], 4)
	}
}

// SyncTimingValidationDisable returns false if SYNC_SMP_ERR pin is configured
// to detect syncronization pulse sampling errors.
func (r *AD9910) SyncTimingValidationDisable() bool {
	return utils.HasBit(r.CtrlFunc2Data[0], 5)
}

// SetSyncTimingValidationDisable configures SYNC_SMP_ERR pin to detect
// syncronization pulse sampling errors.
func (r *AD9910) SetSyncTimingValidationDisable(active bool) {
	r.CtrlFunc2Data[0] = utils.UnsetBit(r.CtrlFunc2Data[0], 5)

	if active {
		r.CtrlFunc2Data[0] = utils.SetBit(r.CtrlFunc2Data[0], 5)
	}
}

// DataAssemblerHoldLastValue relates to some parallel port communication.
func (r *AD9910) DataAssemblerHoldLastValue() bool {
	return utils.HasBit(r.CtrlFunc2Data[0], 6)
}

// SetDataAssemblerHoldLastValue relates to some parallel port communication.
func (r *AD9910) SetDataAssemblerHoldLastValue(active bool) {
	r.CtrlFunc2Data[0] = utils.UnsetBit(r.CtrlFunc2Data[0], 6)

	if active {
		r.CtrlFunc2Data[0] = utils.SetBit(r.CtrlFunc2Data[0], 6)
	}
}

// MatchedLatencyEnable returns true if parameter changes affect the output
// signal in the order made.
func (r *AD9910) MatchedLatencyEnable() bool {
	return utils.HasBit(r.CtrlFunc2Data[0], 7)
}

// SetMatchedLatencyEnable configures the output signal to be affected by
// multiple parameter changes at once if true.
func (r *AD9910) SetMatchedLatencyEnable(active bool) {
	r.CtrlFunc2Data[0] = utils.UnsetBit(r.CtrlFunc2Data[0], 7)

	if active {
		r.CtrlFunc2Data[0] = utils.SetBit(r.CtrlFunc2Data[0], 7)
	}
}

// TxEnableInvert relates to the parallel port communication.
func (r *AD9910) TxEnableInvert() bool {
	return utils.HasBit(r.CtrlFunc2Data[1], 0)
}

// SetTxEnableInvert relates to the parallel port communication.
func (r *AD9910) SetTxEnableInvert(active bool) {
	r.CtrlFunc2Data[1] = utils.UnsetBit(r.CtrlFunc2Data[1], 0)

	if active {
		r.CtrlFunc2Data[1] = utils.SetBit(r.CtrlFunc2Data[1], 0)
	}
}

// ParallelDataClockInvert relates to the parallel port communication.
func (r *AD9910) ParallelDataClockInvert() bool {
	return utils.HasBit(r.CtrlFunc2Data[1], 2)
}

// SetParallelDataClockInvert relates to the parallel port communication.
func (r *AD9910) SetParallelDataClockInvert(active bool) {
	r.CtrlFunc2Data[1] = utils.UnsetBit(r.CtrlFunc2Data[1], 2)

	if active {
		r.CtrlFunc2Data[1] = utils.SetBit(r.CtrlFunc2Data[1], 2)
	}
}

// ParallelDataClockEnable relates to the parallel port communication.
func (r *AD9910) ParallelDataClockEnable() bool {
	return utils.HasBit(r.CtrlFunc2Data[1], 3)
}

// SetParallelDataClockEnable relates to the parallel port communication.
func (r *AD9910) SetParallelDataClockEnable(active bool) {
	r.CtrlFunc2Data[1] = utils.UnsetBit(r.CtrlFunc2Data[1], 3)

	if active {
		r.CtrlFunc2Data[1] = utils.SetBit(r.CtrlFunc2Data[1], 3)
	}
}

// AD9910IOUpdateRateDivider defines the allowed prescale ratios of the
// auto I/O update clocks.
type AD9910IOUpdateRateDivider int

// Allowed dividers for the I/O update rate.
const (
	AD9910IOUpdateRateDividerBy1 AD9910IOUpdateRateDivider = iota
	AD9910IOUpdateRateDividerBy2
	AD9910IOUpdateRateDividerBy4
	AD9910IOUpdateRateDividerBy8
)

// IOUpdateRateCtrl returns the configured I/O update rate divider.
func (r *AD9910) IOUpdateRateCtrl() AD9910IOUpdateRateDivider {
	b := utils.ReadBits(r.CtrlFunc2Data[1], 6, 2)

	return AD9910IOUpdateRateDivider(b)
}

// SetIOUpdateRateCtrl configures the I/O update rate divider to be one of the
// allowed values of AD9910UpdateRateDivider.
func (r *AD9910) SetIOUpdateRateCtrl(d AD9910IOUpdateRateDivider) {
	utils.WriteBits(r.CtrlFunc2Data[1], 6, 2, byte(d))
}

// ReadEffectiveFreqTuningWord returns true if the AD9910 is configured to
// reply with the actual measured frequency if FTW is requested.
func (r *AD9910) ReadEffectiveFreqTuningWord() bool {
	return utils.HasBit(r.CtrlFunc2Data[2], 0)
}

// SetReadEffectiveFreqTuningWord configures the AD9910 to reply with the
// actual frequency instead of the register value if true.
func (r *AD9910) SetReadEffectiveFreqTuningWord(active bool) {
	r.CtrlFunc2Data[2] = utils.UnsetBit(r.CtrlFunc2Data[2], 0)

	if active {
		r.CtrlFunc2Data[2] = utils.SetBit(r.CtrlFunc2Data[2], 0)
	}
}

// DigitalRampNoDwellLow returns true if the digital ramp is configured to
// directly skip back to the configured lower limit instead of ramping back.
func (r *AD9910) DigitalRampNoDwellLow() bool {
	return utils.HasBit(r.CtrlFunc2Data[2], 1)
}

// SetDigitalRampNoDwellLow configures the digital ramp to skip instead of
// ramping back to the opposing limit if true.
func (r *AD9910) SetDigitalRampNoDwellLow(active bool) {
	r.CtrlFunc2Data[2] = utils.UnsetBit(r.CtrlFunc2Data[2], 1)

	if active {
		r.CtrlFunc2Data[2] = utils.SetBit(r.CtrlFunc2Data[2], 1)
	}
}

// DigitalRampEnable returns true if the digital ramp is configured to be
// enabled.
func (r *AD9910) DigitalRampEnable() bool {
	return utils.HasBit(r.CtrlFunc2Data[2], 2)
}

// SetDigitalRampEnable configures the digital ramp to be enabled if
// active is true.
func (r *AD9910) SetDigitalRampEnable(active bool) {
	r.CtrlFunc2Data[2] = utils.UnsetBit(r.CtrlFunc2Data[2], 2)

	if active {
		r.CtrlFunc2Data[2] = utils.SetBit(r.CtrlFunc2Data[2], 2)
	}
}

// AD9910DigitalRampDest defines the allowed parameters to be controlled by
// the digital ramp.
type AD9910DigitalRampDest int

// AD9910DigitalRampDest can be set to control frequency, phase or amplitude
// by the digital ramp.
const (
	AD9910DigitalRampDestFreq AD9910DigitalRampDest = iota
	AD9910DigitalRampDestPhase
	AD9910DigitalRampDestAmpl
)

// DigitalRampDest returns the configured digital ramp destination.
// See AD9910DigitalRampDest for more.
func (r *AD9910) DigitalRampDest() AD9910DigitalRampDest {
	b := utils.ReadBits(r.CtrlFunc2Data[2], 4, 2)

	return AD9910DigitalRampDest(b)
}

// SetDigitalRampDest configures the digital ramp destination.
// See AD9910DigitalRampDest for more.
func (r *AD9910) SetDigitalRampDest(d AD9910DigitalRampDest) {
	utils.WriteBits(r.CtrlFunc2Data[2], 4, 2, byte(d))
}

// SyncClockEnable returns true if the AD9910 is configured to generate
// a clock signal to syncronize SPI communication.
func (r *AD9910) SyncClockEnable() bool {
	return utils.HasBit(r.CtrlFunc2Data[2], 6)
}

// SetSyncClockEnable configures the AD9910 to generate a clock signal
// for SPI communication if active is true.
func (r *AD9910) SetSyncClockEnable(active bool) {
	r.CtrlFunc2Data[2] = utils.UnsetBit(r.CtrlFunc2Data[2], 6)

	if active {
		r.CtrlFunc2Data[2] = utils.SetBit(r.CtrlFunc2Data[2], 6)
	}
}

// IntIOUpdateActive returns true if the SPI communication is configured
// to be syncronized with an internally generated I/O update signal.
func (r *AD9910) IntIOUpdateActive() bool {
	return utils.HasBit(r.CtrlFunc2Data[2], 7)
}

// SetIntIOUpdateActive configures the SPI communication to be syncronized
// with an internally generated I/O update signal if active is true.
func (r *AD9910) SetIntIOUpdateActive(active bool) {
	r.CtrlFunc2Data[2] = utils.UnsetBit(r.CtrlFunc2Data[2], 7)

	if active {
		r.CtrlFunc2Data[2] = utils.SetBit(r.CtrlFunc2Data[2], 7)
	}
}

// AmplScaleFromSTProfileEnable returns true if amplitude from single tone
// profile is configured to be scaled by amplitude scalar factor.
func (r *AD9910) AmplScaleFromSTProfileEnable() bool {
	return utils.HasBit(r.CtrlFunc2Data[3], 0)
}

// SetAmplScaleFromSTProfileEnable configures the amplitude from single tone
// profile to be configured to be scaled by amplitude scalar factor if
// active is true.
func (r *AD9910) SetAmplScaleFromSTProfileEnable(active bool) {
	r.CtrlFunc2Data[3] = utils.UnsetBit(r.CtrlFunc2Data[3], 0)

	if active {
		r.CtrlFunc2Data[3] = utils.SetBit(r.CtrlFunc2Data[3], 0)
	}
}

// RefClockFeedbackDivider returns a 7 bit divide modulus of the reference
// clock phase locked loop feedback divider.
func (r *AD9910) RefClockFeedbackDivider() uint {
	return uint(utils.ReadBits(r.CtrlFunc3Data[0], 1, 7))
}

// SetRefClockFeedbackDivider sets the 7 bit divide modulus of the reference
// clock phase locked loop feedback divider.
func (r *AD9910) SetRefClockFeedbackDivider(d uint) {
	utils.WriteBits(r.CtrlFunc3Data[0], 1, 7, byte(d))
}

// PhaseLockedLoopEnable returns true if reference clock phase locked loop
// is configured to be enabled.
func (r *AD9910) PhaseLockedLoopEnable() bool {
	return utils.HasBit(r.CtrlFunc3Data[1], 0)
}

// SetPhaseLockedLoopEnable configures the reference clock phase locked loop
// (PLL) to be configured enabled if active is true.
func (r *AD9910) SetPhaseLockedLoopEnable(active bool) {
	r.CtrlFunc3Data[1] = utils.UnsetBit(r.CtrlFunc3Data[1], 0)

	if active {
		r.CtrlFunc3Data[1] = utils.SetBit(r.CtrlFunc3Data[1], 0)
	}
}

// PhaseFreqDetectorReset returns true if phase frequency detector (PFD) is
// configured to be disabled.
func (r *AD9910) PhaseFreqDetectorReset() bool {
	return utils.HasBit(r.CtrlFunc3Data[1], 2)
}

// SetPhaseFreqDetectorReset configures the phase frequency detector (PFD)
// to be enabled if active is true.
func (r *AD9910) SetPhaseFreqDetectorReset(active bool) {
	r.CtrlFunc3Data[1] = utils.UnsetBit(r.CtrlFunc3Data[1], 2)

	if active {
		r.CtrlFunc3Data[1] = utils.SetBit(r.CtrlFunc3Data[1], 2)
	}
}

// RefClockInputDividerReset returns true if reference clock input divider
// operates normally.
func (r *AD9910) RefClockInputDividerReset() bool {
	return utils.HasBit(r.CtrlFunc3Data[1], 6)
}

// SetRefClockInputDividerReset configures the reference clock input divider
// to operate normally if active is true and be reset else.
func (r *AD9910) SetRefClockInputDividerReset(active bool) {
	r.CtrlFunc3Data[1] = utils.UnsetBit(r.CtrlFunc3Data[1], 6)

	if active {
		r.CtrlFunc3Data[1] = utils.SetBit(r.CtrlFunc3Data[1], 6)
	}
}

// RefClockInputDividerBypass returns true if reference clock input divider
// is configured to be bypassed.
func (r *AD9910) RefClockInputDividerBypass() bool {
	return utils.HasBit(r.CtrlFunc3Data[1], 7)
}

// SetRefClockInputDividerBypass configures the reference clock input divider
// to be configured bypassed if active is true.
func (r *AD9910) SetRefClockInputDividerBypass(active bool) {
	r.CtrlFunc3Data[1] = utils.UnsetBit(r.CtrlFunc3Data[1], 7)

	if active {
		r.CtrlFunc3Data[1] = utils.SetBit(r.CtrlFunc3Data[1], 7)
	}
}

// AD9910PhaseLockedLoopCurrent defines the charge pump current values of the
// phase locked loop.
type AD9910PhaseLockedLoopCurrent int

// Allowed charge pump current values in micro amperes.
const (
	AD9910PhaseLockedLoopCurrent212 AD9910PhaseLockedLoopCurrent = iota
	AD9910PhaseLockedLoopCurrent237
	AD9910PhaseLockedLoopCurrent262
	AD9910PhaseLockedLoopCurrent287
	AD9910PhaseLockedLoopCurrent312
	AD9910PhaseLockedLoopCurrent337
	AD9910PhaseLockedLoopCurrent363
	AD9910PhaseLockedLoopCurrent387
)

// PhaseLockedLoopCurrent returns the configured phase locked loop current.
func (r *AD9910) PhaseLockedLoopCurrent() AD9910PhaseLockedLoopCurrent {
	b := utils.ReadBits(r.CtrlFunc3Data[2], 3, 3)

	return AD9910PhaseLockedLoopCurrent(b)
}

// SetPhaseLockedLoopCurrent configures the phase locked loop current.
func (r *AD9910) SetPhaseLockedLoopCurrent(c AD9910PhaseLockedLoopCurrent) {
	utils.WriteBits(r.CtrlFunc3Data[2], 3, 3, byte(c))
}

// AD9910RefClockVCORange defines allowed VCO frequency ranges.
type AD9910RefClockVCORange int

const (
	// AD9910RefClockVCORange0 typically ranges between 370 and 510 MHz.
	AD9910RefClockVCORange0 AD9910RefClockVCORange = iota
	// AD9910RefClockVCORange1 typically ranges between 420 and 590 MHz.
	AD9910RefClockVCORange1
	// AD9910RefClockVCORange2 typically ranges between 500 and 700 MHz.
	AD9910RefClockVCORange2
	// AD9910RefClockVCORange3 typically ranges between 600 and 880 MHz.
	AD9910RefClockVCORange3
	// AD9910RefClockVCORange4 typically ranges between 700 and 950 MHz.
	AD9910RefClockVCORange4
	// AD9910RefClockVCORange5 typically ranges between 820 and 1150 MHz.
	AD9910RefClockVCORange5
	// AD9910RefClockVCOPLLBypassed disabled internal PLL (unconfirmed?).
	AD9910RefClockVCOPLLBypassed
)

// RefClockVCORange returns the configured VCO frequency range.
// See AD910RefClockVCORange.
func (r *AD9910) RefClockVCORange() AD9910RefClockVCORange {
	b := utils.ReadBits(r.CtrlFunc3Data[3], 0, 3)

	return AD9910RefClockVCORange(b)
}

// SetRefClockVCORange configures the VCO frequency range.
// See AD910RefClockVCORange.
func (r *AD9910) SetRefClockVCORange(v AD9910RefClockVCORange) {
	utils.WriteBits(r.CtrlFunc3Data[3], 0, 3, byte(v))
}

// AD9910RefClockOutputMode defines the output mode of the reference clock
// output mode.
type AD9910RefClockOutputMode int

// Allowed configurations for AD9910RefClockOutputMode.
const (
	AD9910RefClockOutputModeDisabled AD9910RefClockOutputMode = iota
	AD9910RefClockOutputModeLowCurrent
	AD9910RefClockOutputModeMediumCurrent
	AD9910RefClockOutputModeHighCurrent
)

// RefClockOuputCtrl returns the configured reference clock output mode.
// See AD9910RefClockOutputMode.
func (r *AD9910) RefClockOuputCtrl() AD9910RefClockOutputMode {
	b := utils.ReadBits(r.CtrlFunc3Data[3], 4, 2)

	return AD9910RefClockOutputMode(b)
}

// SetRefClockOuputCtrl configures the reference clock output mode.
// See AD9910RefClockOutputMode.
func (r *AD9910) SetRefClockOuputCtrl(v AD9910RefClockOutputMode) {
	utils.WriteBits(r.CtrlFunc3Data[3], 4, 2, byte(v))
}

// FullScaleOutputCurrent returns the 8 bit code word related to the output
// current of the main DAC.
func (r *AD9910) FullScaleOutputCurrent() uint8 {
	return uint8(r.AuxDACCtrlData[0])
}

// SetFullScaleOutputCurrent configures the main DAC to ammend the output
// current. See datasheet for more details.
func (r *AD9910) SetFullScaleOutputCurrent(v uint8) {
	r.AuxDACCtrlData[0] = byte(v)
}

// IOUpdateRate returns the configured I/O update rate.
func (r *AD9910) IOUpdateRate() uint32 {
	return binary.LittleEndian.Uint32(r.IOUpdateRateData[:])
}

// SetIOUpdateRate configures the I/O update rate. See datasheet for specifics.
func (r *AD9910) SetIOUpdateRate(v uint32) {
	binary.LittleEndian.PutUint32(r.IOUpdateRateData[:], v)
}

// FreqTuningWord returns the configured frequency tuning word.
func (r *AD9910) FreqTuningWord() uint32 {
	return binary.LittleEndian.Uint32(r.FreqTuningWordData[:])
}

// SetFreqTuningWord configures the frequency tuning word.
func (r *AD9910) SetFreqTuningWord(v uint32) {
	binary.LittleEndian.PutUint32(r.FreqTuningWordData[:], v)
}

// PhaseOffsetWord returns the configured phase offset word.
func (r *AD9910) PhaseOffsetWord() uint16 {
	return binary.LittleEndian.Uint16(r.PhaseOffsetWordData[:])
}

// SetPhaseOffsetWord configures the phase offset word.
func (r *AD9910) SetPhaseOffsetWord(v uint16) {
	binary.LittleEndian.PutUint16(r.FreqTuningWordData[:], v)
}

// AD9910AmplStepSize defines the amplitude step size in OSK mode.
type AD9910AmplStepSize int

// Values for amplitude step size in OSK mode.
const (
	AD9910AmplStepSize1 AD9910AmplStepSize = iota
	AD9910AmplStepSize2
	AD9910AmplStepSize4
	AD9910AmplStepSize8
)

// AmplStepSize returns the configured amplitude step size.
func (r *AD9910) AmplStepSize() AD9910AmplStepSize {
	b := utils.ReadBits(r.AmplScaleFactorData[0], 0, 2)

	return AD9910AmplStepSize(b)
}

// SetAmplStepSize configures the amplitude step size.
func (r *AD9910) SetAmplStepSize(v AD9910AmplStepSize) {
	utils.WriteBits(r.AmplScaleFactorData[0], 0, 2, byte(v))
}

// AmplScaleFactor returns the 14 bit amplitude scale factor.
func (r *AD9910) AmplScaleFactor() uint16 {
	return binary.LittleEndian.Uint16(r.AmplScaleFactorData[:2]) << 2
}

// SetAmplScaleFactor configures the 14 bit amplitude scale factor.
func (r *AD9910) SetAmplScaleFactor(v uint16) {
	binary.LittleEndian.PutUint16(r.AmplScaleFactorData[2:], v>>2)
}

// AmplRampRate returns the 16 bit amplitude ramp rate.
func (r *AD9910) AmplRampRate() uint16 {
	return binary.LittleEndian.Uint16(r.AmplScaleFactorData[2:])
}

// SetAmplRampRate configures the 16 bit amplitude ramp rate.
func (r *AD9910) SetAmplRampRate(v uint16) {
	binary.LittleEndian.PutUint16(r.AmplScaleFactorData[2:], v)
}

// InputSyncReceiverDelay returns a 5 bit number configured to be the input
// delay of the sync receiver.
func (r *AD9910) InputSyncReceiverDelay() uint8 {
	return uint8(utils.ReadBits(r.MultichipSyncData[0], 3, 5))
}

// SetInputSyncReceiverDelay configures the sync receiver input delay.
func (r *AD9910) SetInputSyncReceiverDelay(v uint8) {
	utils.WriteBits(r.MultichipSyncData[0], 3, 5, byte(v))
}

// OutputSyncReceiverDelay returns a 5 bit number configured to be the output
// delay of the sync receiver.
func (r *AD9910) OutputSyncReceiverDelay() uint8 {
	return uint8(utils.ReadBits(r.MultichipSyncData[1], 3, 5))
}

// SetOutputSyncReceiverDelay configures the sync receiver output delay.
func (r *AD9910) SetOutputSyncReceiverDelay(v uint8) {
	utils.WriteBits(r.MultichipSyncData[1], 3, 5, byte(v))
}

// SyncStatePresetValue returns a 6 bit number configured to define the state
// of the internal clock generator for the sync pulse.
func (r *AD9910) SyncStatePresetValue() uint8 {
	return uint8(utils.ReadBits(r.MultichipSyncData[2], 2, 6))
}

// SetSyncStatePresetValue configures a 6 bit number to define the state of
// the internal clock generator.
func (r *AD9910) SetSyncStatePresetValue(v uint8) {
	utils.WriteBits(r.MultichipSyncData[2], 2, 6, byte(v))
}

// SyncGeneratorPolarity returns true if syncronisation clock generator is
// configured to coincident with the falling edge of the system clock.
func (r *AD9910) SyncGeneratorPolarity() bool {
	return utils.HasBit(r.MultichipSyncData[3], 1)
}

// SetSyncGeneratorPolarity configures the syncronisation clock generator to
// be configured to coincident with the falling edge of the system clock.
func (r *AD9910) SetSyncGeneratorPolarity(active bool) {
	utils.UnsetBit(r.MultichipSyncData[3], 1)

	if active {
		utils.SetBit(r.MultichipSyncData[3], 1)
	}
}

// SyncGeneratorEnable returns true if the syncronisation clock generator is
// configured to be enabled.
func (r *AD9910) SyncGeneratorEnable() bool {
	return utils.HasBit(r.MultichipSyncData[3], 2)
}

// SetSyncGeneratorEnable configures the syncronisation clock generator to be
// enabled if active is true.
func (r *AD9910) SetSyncGeneratorEnable(active bool) {
	utils.UnsetBit(r.MultichipSyncData[3], 2)

	if active {
		utils.SetBit(r.MultichipSyncData[3], 2)
	}
}

// SyncReceiverEnable returns true if the syncronisation clock receiver is
// configured to be enabled.
func (r *AD9910) SyncReceiverEnable() bool {
	return utils.HasBit(r.MultichipSyncData[3], 3)
}

// SetSyncReceiverEnable configures the syncronisation clock receiver to be
// configured enabled if active is true.
func (r *AD9910) SetSyncReceiverEnable(active bool) {
	utils.UnsetBit(r.MultichipSyncData[3], 3)

	if active {
		utils.SetBit(r.MultichipSyncData[3], 3)
	}
}

// SyncValidationDelay returns a 4 bit number configured to be the skew timing
// between SYSCLK and SYNC_INx signal.
func (r *AD9910) SyncValidationDelay() uint8 {
	return uint8(utils.ReadBits(r.MultichipSyncData[3], 4, 4))
}

// SetSyncValidationDelay configures a 4 bit number to be configured as the
// skew timing between SYSCLK and SYNC_INx.
func (r *AD9910) SetSyncValidationDelay(v uint8) {
	utils.WriteBits(r.MultichipSyncData[3], 4, 4, byte(v))
}

// DigitalRampLowerLimit returns the 32 bit number configured to be the
// lower limit of the digital ramp.
func (r *AD9910) DigitalRampLowerLimit() uint32 {
	return binary.LittleEndian.Uint32(r.DigitalRampLimitData[:4])
}

// SetDigitalRampLowerLimit configures the 32 bit number to be the lower
// limit of the digital ramp.
func (r *AD9910) SetDigitalRampLowerLimit(v uint32) {
	binary.LittleEndian.PutUint32(r.DigitalRampLimitData[:4], v)
}

// DigitalRampUpperLimit returns the 32 bit number configured to be the
// upper limit of the digital ramp.
func (r *AD9910) DigitalRampUpperLimit() uint32 {
	return binary.LittleEndian.Uint32(r.DigitalRampLimitData[4:])
}

// SetDigitalRampUpperLimit configures the 32 bit number to be the upper
// limit of the digital ramp.
func (r *AD9910) SetDigitalRampUpperLimit(v uint32) {
	binary.LittleEndian.PutUint32(r.DigitalRampLimitData[4:], v)
}

// DigitalRampIncrementStepSize returns the 32 bit number configured to be the
// increment step size of the digital ramp.
func (r *AD9910) DigitalRampIncrementStepSize() uint32 {
	return binary.LittleEndian.Uint32(r.DigitalRampStepSizeData[:4])
}

// SetDigitalRampIncrementStepSize configures the 32 bit number to be the
// increment step size of the digital ramp.
func (r *AD9910) SetDigitalRampIncrementStepSize(v uint32) {
	binary.LittleEndian.PutUint32(r.DigitalRampStepSizeData[:4], v)
}

// DigitalRampDecrementStepSize returns the 32 bit number configured to be the
// decrement step size of the digital ramp.
func (r *AD9910) DigitalRampDecrementStepSize() uint32 {
	return binary.LittleEndian.Uint32(r.DigitalRampStepSizeData[4:])
}

// SetDigitalRampDecrementStepSize configures the 32 bit number to be the
// decrement step size of the digital ramp.
func (r *AD9910) SetDigitalRampDecrementStepSize(v uint32) {
	binary.LittleEndian.PutUint32(r.DigitalRampStepSizeData[4:], v)
}

// DigitalRampPositiveSlopeRate returns the 16 bit number configured to be
// the positive slope rate of the digital ramp.
func (r *AD9910) DigitalRampPositiveSlopeRate() uint16 {
	return binary.LittleEndian.Uint16(r.DigitalRampRateData[:2])
}

// SetDigitalRampPositiveSlopeRate configures the 16 bit number to be
// the positive slope rate of the digital ramp.
func (r *AD9910) SetDigitalRampPositiveSlopeRate(v uint16) {
	binary.LittleEndian.PutUint16(r.DigitalRampRateData[:2], v)
}

// DigitalRampNegativeSlopeRate returns the 16 bit number configured to be
// the negative slope rate of the digital ramp.
func (r *AD9910) DigitalRampNegativeSlopeRate() uint16 {
	return binary.LittleEndian.Uint16(r.DigitalRampRateData[2:])
}

// SetDigitalRampNegativeSlopeRate configures the 16 bit number to be
// the negative slope rate of the digital ramp.
func (r *AD9910) SetDigitalRampNegativeSlopeRate(v uint16) {
	binary.LittleEndian.PutUint16(r.DigitalRampRateData[:2], v)
}

// STProfile0FreqTuningWord returns the configured 32 bit number to be
// the frequency tuning word of single tone profile 0.
func (r *AD9910) STProfile0FreqTuningWord() uint32 {
	return binary.LittleEndian.Uint32(r.STProfile0Data[:4])
}

// SetSTProfile0FreqTuningWord configures the 32 bit number to be
// the frequency tuning word of single tone profile 0.
func (r *AD9910) SetSTProfile0FreqTuningWord(v uint32) {
	binary.LittleEndian.PutUint32(r.STProfile0Data[:4], v)
}

// STProfile0PhaseOffsetWord returns the configured 16 bit number to be
// the phase offset word of single tone profile 0.
func (r *AD9910) STProfile0PhaseOffsetWord() uint16 {
	return binary.LittleEndian.Uint16(r.STProfile0Data[4:6])
}

// SetSTProfile0PhaseOffsetWord configures the 16 bit number to be
// the phase offset word of single tone profile 0.
func (r *AD9910) SetSTProfile0PhaseOffsetWord(v uint16) {
	binary.LittleEndian.PutUint16(r.STProfile0Data[4:6], v)
}

// STProfile0AmplScaleFactor returns the configured 14 bit number to be
// the amplitude scale factor single tone profile 0.
func (r *AD9910) STProfile0AmplScaleFactor() uint16 {
	return binary.LittleEndian.Uint16(r.STProfile0Data[6:])
}

// SetSTProfile0AmplScaleFactor configures the 14 bit number to be
// the amplitude scale factor of single tone profile 0.
func (r *AD9910) SetSTProfile0AmplScaleFactor(v uint16) {
	binary.LittleEndian.PutUint16(r.STProfile0Data[6:], v)
}
