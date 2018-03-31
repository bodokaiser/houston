package register

// AD9910 represents the registers of the AD9910 digital direct synthesizer
// (DDS) chip.
type AD9910 struct {
	// ControlFunction1 represents the Control Function Register 1 (CFR1)
	// register which is used generic configuration.
	CtrlFunc1 [4]byte

	// ControlFunction2 represents the Control Function Register 2 (CFR2)
	// register which is used generic configuration.
	CtrlFunc2 [4]byte

	// ControlFunction3 represents the Control Function Register3 (CFR3)
	// register which is used generic configuration.
	CtrlFunc3 [4]byte

	// AuxiliaryDACControl represents the Auxiliary DAC Control register which
	// is used to configure the auxiliary digital to analog converter (DAC)
	// responsible for control of the output current of the main DAC.
	AuxDACCtrl [4]byte

	// IOUpdateRate represents the I/O Update Rate register which is used to
	// determine the pulse width of the I/O update signal.
	IOUpdateRate [4]byte

	// FrequencyTuningWord represents the Frequency Tuning Word (FTW) register
	// which is used to determine the output frequency.
	FreqTuningWord [4]byte

	// PhaseOffsetWord represents the Phase Offset Word (POW) register which is
	// used to determine the output phase offset.
	PhaseOffsetWord [2]byte

	// AmplitudeScaleFactor represents the Amplitude Scale Factor (ASF) register
	// which is used to scale the output amplitude.
	AmplScaleFactor [4]byte

	// MultichipSync represents the Multichip Sync register which is used for
	// syncronisation of multiple DDS chips.
	MultichipSync [4]byte

	// DigitalRampLimit represents the Digital Ramp Limit register which is used
	// to determine lower and upper limit of a digital ramp sweep.
	DigitalRampLimit [8]byte

	// DigitalRampStep represents the Digital Ramp Step register which is used
	// to determine increment and decrement step size of digital ramp sweep.
	DigitalRampStep [8]byte

	// DigitalRampRate represents the Digital Ramp Rate register which is used
	// to determine positive and negative slope rate of digital ramp sweep.
	DigitalRampRate [4]byte

	// SingleToneProfile0 represents the Single Tone Profile 0 register which
	// is used to determine constant frequency generation parameters.
	STProfile0 [8]byte

	// RAMProfile0 represents the RAM Profile 0 register which is used to define
	// custom waveform from RAM frequency generation.
	RAMProfile0 [8]byte
}

// NewAD9910 returns a new AD9910 initialized with the default values in
// the registers.
func NewAD9910() *AD9910 {
	return &AD9910{
		[]byte{0x00, 0x00, 0x00, 0x00}},
		Register{0x01, []byte{0x00, 0x40, 0x08, 0x20}},
		Register{0x02, []byte{0x1f, 0x3f, 0x40, 0x00}},
		Register{0x03, []byte{0x00, 0x00, 0x00, 0x7f}},
		Register{0x04, []byte{0xff, 0xff, 0xff, 0xff}},
		Register{0x07, []byte{0x00, 0x00, 0x00, 0x00}},
		Register{0x08, []byte{0x00, 0x00}},
		Register{0x09, []byte{0x00, 0x00, 0x00, 0x00}},
		Register{0x0a, []byte{0x00, 0x00, 0x00, 0x00}},
		Register{0x0b, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		Register{0x0c, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		Register{0x0d, []byte{0x00, 0x00, 0x00, 0x00}},
		Register{0x0e, []byte{0x08, 0xb5, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		Register{0x0e, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		Register{0x0f, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		Register{0x0f, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
	}
}

func (d *AD9910) SPIByteOrder() {}

func (d *AD9910) SPIInputMode() {}

func (d *AD9910) ExtPowerDown() {}

func (d *AD9910) DACPowerDown() {}

func (d *AD9910) AuxDACPowerDown() {}

func (d *AD9910) CorePowerDown() {}

func (d *AD9910) RefClockInputPowerDown() {}

func (d *AD9910) ClearPhaseAccumulator() {}

func (d *AD9910) ClearDigitalRampAccumulator() {}

func (d *AD9910) AutoClearPhaseAccumulator() {}

func (d *AD9910) AutoClearDigitalRampAccumulator() {}

func (d *AD9910) AutoClearDigitalRampAccumulator() {}

func (d *AD9910) SelectSineOutput() {}
