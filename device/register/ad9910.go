package register

import "github.com/bodokaiser/beagle/util"

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
	AD9910DRampLimitAddress
	AD9910DRampStepSizeAddress
	AD9910DRampRateAddress
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
	CtrlFunc1       [4]byte
	CtrlFunc2       [4]byte
	CtrlFunc3       [4]byte
	AuxDACCtrl      [4]byte
	IOUpdateRate    [4]byte
	FreqTuningWord  [4]byte
	PhaseOffsetWord [2]byte
	AmplScaleFactor [4]byte
	MultichipSync   [4]byte
	DRampLimit      [8]byte
	DRampStepSize   [8]byte
	DRampRate       [4]byte
	STProfile0      [8]byte
	STProfile1      [8]byte
	RAMProfile0     [8]byte
	RAMProfile1     [8]byte
	RAMMemory       [4]byte
}

// LSBFirst returns true if SPI byte order is configured to be
// LSB and false if SPI byte order is MSB.
func (r *AD9910) LSBFirst() bool {
	return util.HasBit(r.CtrlFunc1[0], 0)
}

// SetLSBFirst configures the SPI byte order to be LSB on
// true and MSB on false.
func (r *AD9910) SetLSBFirst(active bool) {
	r.CtrlFunc1[0] = util.UnsetBit(r.CtrlFunc1[0], 0)

	if active {
		r.CtrlFunc1[0] = util.SetBit(r.CtrlFunc1[0], 0)
	}
}

// SDIOInputOnly returns true if SPI uses 3-wire mode and
// false if SDIO pin is used bidirectional.
func (r *AD9910) SDIOInputOnly() bool {
	return util.HasBit(r.CtrlFunc1[0], 1)
}

// SetSDIOInputOnly configures SPI to use 3-wire on true and
// 2-wire on false.
func (r *AD9910) SetSDIOInputOnly(active bool) {
	r.CtrlFunc1[0] = util.UnsetBit(r.CtrlFunc1[0], 1)

	if active {
		r.CtrlFunc1[0] = util.SetBit(r.CtrlFunc1[0], 1)
	}
}

// AD9910PowerDownMode defines how the AD9910 responds to
// an external power down request.
type AD9910PowerDownMode int

// AD9910 can be configured to go in full or fast recovery
// power down.
const (
	FullPowerDown AD9910PowerDownMode = iota
	FastRecoveryPowerDown
)

// ExtPowerDownCtrl returns the configured power down
// behaviour (see AD9910PowerDownModee).
func (r *AD9910) ExtPowerDownCtrl() AD9910PowerDownMode {
	if util.HasBit(r.CtrlFunc1[0], 3) {
		return FastRecoveryPowerDown
	}
	return FullPowerDown
}

// SetExtPowerDownCtrl configures the power down
// behaviour (see AD9910PowerDownModee).
func (r *AD9910) SetExtPowerDownCtrl(m AD9910PowerDownMode) {
	r.CtrlFunc1[0] = util.UnsetBit(r.CtrlFunc1[0], 3)

	if FastRecoveryPowerDown == m {
		r.CtrlFunc1[0] = util.SetBit(r.CtrlFunc1[0], 3)
	}
}

// AuxiliaryDACPowerDown returns true if auxiliary DAC is
// configured to be powered down.
func (r *AD9910) AuxiliaryDACPowerDown() bool {
	return util.HasBit(r.CtrlFunc1[0], 4)
}

// SetAuxiliaryDACPowerDown configures the auxiliary DAC to
// be powered down.
func (r *AD9910) SetAuxiliaryDACPowerDown(active bool) {
	r.CtrlFunc1[0] = util.UnsetBit(r.CtrlFunc1[0], 4)

	if active {
		r.CtrlFunc1[0] = util.SetBit(r.CtrlFunc1[0], 4)
	}
}

// RefClockInputPowerDown returns true if external reference
// clock input pin is configured to be powered down.
func (r *AD9910) RefClockInputPowerDown() bool {
	return util.HasBit(r.CtrlFunc1[0], 5)
}

// SetRefClockInputPowerDown configures the external reference
// clock input pin to be configured powered down.
func (r *AD9910) SetRefClockInputPowerDown(active bool) {
	r.CtrlFunc1[0] = util.UnsetBit(r.CtrlFunc1[0], 5)

	if active {
		r.CtrlFunc1[0] = util.SetBit(r.CtrlFunc1[0], 5)
	}
}

// DACPowerDown returns true if main DAC is configured to be
// powered down.
func (r *AD9910) DACPowerDown() bool {
	return util.HasBit(r.CtrlFunc1[0], 6)
}

// SetDACPowerDown configures the main DAC to be configured
// powered down.
func (r *AD9910) SetDACPowerDown(active bool) {
	r.CtrlFunc1[0] = util.UnsetBit(r.CtrlFunc1[0], 6)

	if active {
		r.CtrlFunc1[0] = util.SetBit(r.CtrlFunc1[0], 6)
	}
}

// DigitalPowerDown returns true if core is configured to be
// powered down.
func (r *AD9910) DigitalPowerDown() bool {
	return util.HasBit(r.CtrlFunc1[0], 7)
}

// SetDigitalPowerDown configures the core to be powered down.
func (r *AD9910) SetDigitalPowerDown(active bool) {
	r.CtrlFunc1[0] = util.UnsetBit(r.CtrlFunc1[0], 7)

	if active {
		r.CtrlFunc1[0] = util.SetBit(r.CtrlFunc1[0], 7)
	}
}
