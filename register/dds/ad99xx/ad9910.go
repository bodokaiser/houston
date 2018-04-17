package ad99xx

// AD9910 register addresses.
const (
	// The control function registers hold the main configuration parameters.
	AddrCFR1 = 0x00
	AddrCFR2 = 0x01
	AddrCFR3 = 0x02

	// The auxiliary digital analogue converter register can be used to control
	// the output current but it can also be disabled.
	AddrAuxDAC = 0x03

	// The I/O update rate register configures the pulse width required to apply
	// on the I/O update pin to trigger an operation update.
	AddrIOUpdateRate = 0x04

	// The frequency tuning word registers are used to configure the output
	// frequency.
	AddrFTW = 0x07

	// The phase offset word register is used to configure the output phase offset.
	AddrPOW = 0x08

	// The amplitude scale factor register is used to configure the output
	// amplitude scale.
	AddrASF = 0x09

	// The multi chip register is used to configure additional syncronisation
	// features when operating multiple AD9910.
	AddrMultiChip = 0x0a

	// The digital ramp register is used to define a sweep of amplitude, phase
	// or frequency.
	AddrDRampLimit    = 0x0b
	AddrDRampStepSize = 0x0c
	AddrDRampRate     = 0x0d

	// The profile register is used to save different single tone or RAM
	// configurations.
	AddrProfile0 = 0x0e
	AddrProfile1 = 0x0f
	AddrProfile2 = 0x10
	AddrProfile3 = 0x11
	AddrProfile4 = 0x12
	AddrProfile5 = 0x13
	AddrProfile6 = 0x14
	AddrProfile7 = 0x15

	// The RAM register is used to save custom waveform shapes.
	AddrRAM = 0x16
)

// AD9910 control function register byte flags.
//
// Check the last pages of the datasheet to get more detailed information of
// what these flags will do.
const (
	FlagRAMEnable         = 1 << 7
	FlagManualOSK         = 1 << 7
	FlagInverseSinc       = 1 << 6
	FlagSineOutput        = 1 << 0
	FlagAutomaticOSK      = 1 << 0
	FlagOSKEnable         = 1 << 1
	FlagLoadARR           = 1 << 2
	FlagClearPhaseAcc     = 1 << 3
	FlagClearDRampAcc     = 1 << 4
	FlagAutoclearPhaseAcc = 1 << 5
	FlagAutoclearDRampAcc = 1 << 6
	FlagLoadLRR           = 1 << 7
	FlagLSBFirst          = 1 << 0
	FlagSDIOInput         = 1 << 1
	FlagExtPowerDown      = 1 << 2
	FlagAuxDACPowerDown   = 1 << 3
	FlagREFCLKPowerDown   = 1 << 4
	FlagDACPowerDown      = 1 << 5
	FlagDigitalPowerDown  = 1 << 6
	FlagAmplScaleEnable   = 1 << 0
	FlagReadEffectiveFTW  = 1 << 0
	FlagDRampNoDwellLow   = 1 << 1
	FlagDRampNoDwellHigh  = 1 << 2
	FlagDRampEnable       = 1 << 3
	FlagSYNCCLKEnable     = 1 << 6
	FlagIntIOUpdateActive = 1 << 7
	FlagTxEnableInvert    = 1 << 1
	FlagPDCLKInvert       = 1 << 2
	FlagPDCLKEnable       = 1 << 3
	FlagPDataEnable       = 1 << 4
	FlagSyncValidDisable  = 1 << 5
	FlagDAssemblerLastVal = 1 << 6
	FlagMatchedLatEnable  = 1 << 7
	FlagPLLEnable         = 1 << 0
	FlagPFDReset          = 1 << 2
	FlagREFCLKDivReset    = 1 << 6
	FlagREFCLKDivBypass   = 1 << 7
)

// AD9910 modes to control REFCLK_OUT signal.
const (
	ModeDRV0Disabled = iota << 4
	ModeDRV0OutputCurrentLow
	ModeDRV0OutputCurrentMedium
	ModeDRV0OutputCurrentHigh
)

// AD9910 modes to control VCO range.
const (
	ModeVCORange0 = iota
	ModeVCORange1
	ModeVCORange2
	ModeVCORange3
	ModeVCORange4
	ModeVCORange5
	ModeVCORangeByPassed1
	ModeVCORangeByPassed2
)

// Ad9910 modes to contorl PLL charge pump current.
const (
	ModeChargePumpCurrent212 = iota << 3
	ModeChargePumpCurrent237
	ModeChargePumpCurrent262
	ModeChargePumpCurrent287
	ModeChargePumpCurrent312
	ModeChargePumpCurrent337
	ModeChargePumpCurrent363
	ModeChargePumpCurrent387
)

// AD9910 images the registers of the AD9910 DDS chip with the first
// byte being the register address.
type AD9910 struct {
	CFR1          [5]byte
	CFR2          [5]byte
	CFR3          [5]byte
	AuxDAC        [5]byte
	IOUpdateRate  [5]byte
	FTW           [5]byte
	POW           [3]byte
	ASF           [5]byte
	MultiChip     [5]byte
	DRampLimit    [9]byte
	DRampStepSize [9]byte
	DRampRate     [5]byte
	STProfile0    [9]byte
	STProfile1    [9]byte
	STProfile2    [9]byte
	STProfile3    [9]byte
	STProfile4    [9]byte
	STProfile5    [9]byte
	STProfile6    [9]byte
	STProfile7    [9]byte
	RAMProfile0   [9]byte
	RAMProfile1   [9]byte
	RAMProfile2   [9]byte
	RAMProfile3   [9]byte
	RAMProfile4   [9]byte
	RAMProfile5   [9]byte
	RAMProfile6   [9]byte
	RAMProfile7   [9]byte
	RAMWord       [5]byte
}

// NewAD9910 returns an initialized AD9910 with register addresses and
// default values from the datasheet.
func NewAD9910() *AD9910 {
	return &AD9910{
		CFR1: [5]byte{AddrCFR1},
		CFR2: [5]byte{AddrCFR2,
			0,
			FlagSYNCCLKEnable,
			FlagPDCLKEnable,
			FlagSyncValidDisable,
		},
		CFR3: [5]byte{AddrCFR3,
			ModeDRV0OutputCurrentLow | ModeVCORangeByPassed2,
			ModeChargePumpCurrent387,
			FlagREFCLKDivReset,
			0,
		},
		AuxDAC: [5]byte{AddrAuxDAC,
			0, 0, 0, 0x7f,
		},
		IOUpdateRate: [5]byte{AddrIOUpdateRate,
			0xff, 0xff, 0xff, 0xff,
		},
		FTW:           [5]byte{AddrFTW},
		POW:           [3]byte{AddrPOW},
		ASF:           [5]byte{AddrASF},
		MultiChip:     [5]byte{AddrMultiChip},
		DRampLimit:    [9]byte{AddrDRampLimit},
		DRampStepSize: [9]byte{AddrDRampStepSize},
		DRampRate:     [5]byte{AddrDRampRate},
		STProfile0: [9]byte{AddrProfile0,
			0x08, 0xb5,
		},
		STProfile1:  [9]byte{AddrProfile1},
		STProfile2:  [9]byte{AddrProfile2},
		STProfile3:  [9]byte{AddrProfile3},
		STProfile4:  [9]byte{AddrProfile4},
		STProfile5:  [9]byte{AddrProfile5},
		STProfile6:  [9]byte{AddrProfile6},
		STProfile7:  [9]byte{AddrProfile7},
		RAMProfile0: [9]byte{AddrProfile0},
		RAMProfile1: [9]byte{AddrProfile1},
		RAMProfile2: [9]byte{AddrProfile2},
		RAMProfile3: [9]byte{AddrProfile3},
		RAMProfile4: [9]byte{AddrProfile4},
		RAMProfile5: [9]byte{AddrProfile5},
		RAMProfile6: [9]byte{AddrProfile6},
		RAMProfile7: [9]byte{AddrProfile7},
		RAMWord:     [5]byte{AddrRAM},
	}
}
