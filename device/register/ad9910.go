package register

// AD9910 register default values.
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

// AD9910 register addresses.
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

// AD9910 represents the registers of the AD9910 direct digital
// synthesizer.
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
	return r.CtrlFunc1[0]&1 == 1
}

// SetLSBFirst configures the SPI byte order to be LSB on
// true and MSB on false.
func (r *AD9910) SetLSBFirst(active bool) {
  r.CtrlFunc1[0] &= ~(1 << 0)

	if active {
		r.CtrlFunc1[0] |= 1
	}
}

func (r *AD9910) SDIOInputOnly() bool {
	return r.CtrlFunc1[0]&1<<1 == 1
}
