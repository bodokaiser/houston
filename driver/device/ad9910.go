package device

var (
	// ControlFunctionRegister1 represents the Control Function Register 1 (CFR1)
	// register which is used generic configuration.
	ControlFunctionRegister1 = &Register{
		0x00, []byte{0x00, 0x00, 0x00, 0x00}}

	// ControlFunctionRegister2 represents the Control Function Register 2 (CFR2)
	// register which is used generic configuration.
	ControlFunctionRegister2 = &Register{
		0x01, []byte{0x00, 0x40, 0x08, 0x20}}

	// ControlFunctionRegister3 represents the Control Function Register3 (CFR3)
	// register which is used generic configuration.
	ControlFunctionRegister3 = &Register{
		0x02, []byte{0x1f, 0x3f, 0x40, 0x00}}

	// AuxiliaryDACControl represents the Auxiliary DAC Control register which
	// is used to configure the auxiliary digital to analog converter (DAC)
	// responsible for control of the output current of the main DAC.
	AuxiliaryDACControl = &Register{
		0x03, []byte{0x00, 0x00, 0x00, 0x7f}}

	// IOUpdateRate represents the I/O Update Rate register which is used to
	// determine the pulse width of the I/O update signal.
	IOUpdateRate = &Register{
		0x04, []byte{0xff, 0xff, 0xff, 0xff}}

	// FrequencyTuningWord represents the Frequency Tuning Word (FTW) register
	// which is used to determine the output frequency.
	FrequencyTuningWord = &Register{
		0x07, []byte{0x00, 0x00, 0x00, 0x00}}

	// PhaseOffsetWord represents the Phase Offset Word (POW) register which is
	// used to determine the output phase offset.
	PhaseOffsetWord = &Register{
		0x08, []byte{0x00, 0x00}}

	// AmplitudeScaleFactor represents the Amplitude Scale Factor (ASF) register
	// which is used to scale the output amplitude.
	AmplitudeScaleFactor = &Register{
		0x09, []byte{0x00, 0x00, 0x00, 0x00}}

	// MultichipSync represents the Multichip Sync register which is used for
	// syncronisation of multiple DDS chips.
	MultichipSync = &Register{
		0x0a, []byte{0x00, 0x00, 0x00, 0x00}}

	// DigitalRampLimit represents the Digital Ramp Limit register which is used
	// to determine lower and upper limit of a digital ramp sweep.
	DigitalRampLimit = &Register{
		0x0b, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}}

	// DigitalRampStep represents the Digital Ramp Step register which is used
	// to determine increment and decrement step size of digital ramp sweep.
	DigitalRampStep = &Register{
		0x0c, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}}

	// DigitalRampRate represents the Digital Ramp Rate register which is used
	// to determine positive and negative slope rate of digital ramp sweep.
	DigitalRampRate = &Register{
		0x0d, []byte{0x00, 0x00, 0x00, 0x00}}

	// SingleToneProfile0 represents the Single Tone Profile 0 register which
	// is used to determine constant frequency generation parameters.
	SingleToneProfile0 = &Register{
		0x0e, []byte{0x08, 0xb5, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}}

	// RAMProfile0 represents the RAM Profile 0 register which is used to define
	// custom waveform from RAM frequency generation.
	RAMProfile0 = &Register{
		0x0e, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}}

	// SingleToneProfile1 represents Single Tone Profile 1 register. See
	// SingleToneProfile0.
	SingleToneProfile1 = &Register{
		0x0f, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}}

	// RAMProfile1 represents the RAM Profile 1 register. See RAMProfile0.
	RAMProfile1 = &Register{
		0x0f, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}}
)
