package dds

import (
	"math"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/spi"
)

const (
	ad9910Read  = 0x80
	ad9910Write = 0x00
)

const (
	ad9910AddrCFR1          = 0x00
	ad9910AddrCFR2          = 0x01
	ad9910AddrCFR3          = 0x02
	ad9910AddrAuxDAC        = 0x03
	ad9910AddrIOUpdateRate  = 0x04
	ad9910AddrFTW           = 0x07
	ad9910AddrPOW           = 0x08
	ad9910AddrASF           = 0x09
	ad9910AddrMultiChip     = 0x0a
	ad9910AddrDRampLimit    = 0x0b
	ad9910AddrDRampStepSize = 0x0c
	ad9910AddrDRampRate     = 0x0d
	ad9910AddrProfile0      = 0x0e
	ad9910AddrProfile1      = 0x0f
	ad9910AddrProfile2      = 0x10
	ad9910AddrProfile3      = 0x11
	ad9910AddrProfile4      = 0x12
	ad9910AddrProfile5      = 0x13
	ad9910AddrProfile6      = 0x14
	ad9910AddrProfile7      = 0x15
	ad9910AddrRAM           = 0x16
)

const (
	ad9910FlagRAMEnable         = 1 << 7
	ad9910FlagManualOSK         = 1 << 7
	ad9910FlagInverseSinc       = 1 << 6
	ad9910FlagSineOutput        = 1 << 0
	ad9910FlagAutomaticOSK      = 1 << 0
	ad9910FlagOSKEnable         = 1 << 1
	ad9910FlagLoadARR           = 1 << 2
	ad9910FlagClearPhaseAcc     = 1 << 3
	ad9910FlagClearDRampAcc     = 1 << 4
	ad9910FlagAutoclearPhaseAcc = 1 << 5
	ad9910FlagAutoclearDRampAcc = 1 << 6
	ad9910FlagLoadLRR           = 1 << 7
	ad9910FlagLSBFirst          = 1 << 0
	ad9910FlagSDIOInput         = 1 << 1
	ad9910FlagExtPowerDown      = 1 << 2
	ad9910FlagAuxDACPowerDown   = 1 << 3
	ad9910FlagREFCLKPowerDown   = 1 << 4
	ad9910FlagDACPowerDown      = 1 << 5
	ad9910FlagDigitalPowerDown  = 1 << 6
	ad9910FlagAmplScaleEnable   = 1 << 0
	ad9910FlagReadEffectiveFTW  = 1 << 0
	ad9910FlagDRampNoDwellLow   = 1 << 1
	ad9910FlagDRampNoDwellHigh  = 1 << 2
	ad9910FlagDRampEnable       = 1 << 3
	ad9910FlagSYNCCLKEnable     = 1 << 6
	ad9910FlagIntIOUpdateActive = 1 << 7
	ad9910FlagTxEnableInvert    = 1 << 1
	ad9910FlagPDCLKInvert       = 1 << 2
	ad9910FlagPDCLKEnable       = 1 << 3
	ad9910FlagPDataEnable       = 1 << 4
	ad9910FlagSyncValidDisable  = 1 << 5
	ad9910FlagDAssemblerLastVal = 1 << 6
	ad9910FlagMatchedLatEnable  = 1 << 7
	ad9910FlagPLLEnable         = 1 << 0
	ad9910FlagPFDReset          = 1 << 2
	ad9910FlagREFCLKDivReset    = 1 << 6
	ad9910FlagREFCLKDivBypass   = 1 << 7
)

type ad9910Config struct {
	CFR1          [5]byte
	CFR2          [5]byte
	CFR3          [5]byte
	AuxDAC        [5]byte
	IOUpdateRate  [5]byte
	FTW           [5]byte
	POW           [5]byte
	ASF           [5]byte
	MultiChip     [5]byte
	DRampLimit    [9]byte
	DRampStepSize [9]byte
	DRampRate     [9]byte
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

// Default clock parameters required for PLL and conversion of frequency to
// FTW.
const (
	AD9910DefaultREFCLK = 10e6
	AD9910DefaultSYSCLK = 1e9
)

// AD9910 is the SPI driver for the AD9910 DDS chip.
type AD9910 struct {
	name       string
	connector  spi.Connector
	connection spi.Connection
}

// NewAD9910 initializes a new instance of the AD9910 device driver.
func NewAD9910(c spi.Connector) *AD9910 {
	return &AD9910{
		name:      gobot.DefaultName("AD9910"),
		connector: c,
	}
}

// Name returns the name of the device.
func (d *AD9910) Name() string {
	return d.name
}

// SetName sets the name of the device.
func (d *AD9910) SetName(s string) {
	d.name = s
}

// Connection returns the spi connection used by the device.
func (d *AD9910) Connection() gobot.Connection {
	return d.connection.(gobot.Connection)
}

// RunSingleTone configures the AD9910 to run in single tone mode.
func (d *AD9910) RunSingleTone() (err error) {
	c := new(ad9910Config)

	c.CFR1[0] |= ad9910Write
	c.CFR1[2] |= ad9910FlagManualOSK
	c.CFR1[3] |= ad9910FlagOSKEnable
	c.CFR1[4] |= ad9910FlagSDIOInput

	err = d.connection.Tx(c.CFR1[:], make([]byte, len(c.CFR1)))
	if err != nil {
		return
	}

	c.CFR2[0] |= ad9910Write
	c.CFR2[1] |= ad9910FlagAmplScaleEnable
	c.CFR2[2] |= ad9910FlagSYNCCLKEnable
	c.CFR2[3] |= ad9910FlagPDCLKEnable
	c.CFR2[4] |= ad9910FlagSyncValidDisable

	err = d.connection.Tx(c.CFR2[:], make([]byte, len(c.CFR2)))
	if err != nil {
		return
	}

	c.POW[0] |= ad9910Write
	c.POW[1] = 0x00
	c.POW[2] = 0x00

	err = d.connection.Tx(c.POW[:], make([]byte, len(c.POW)))
	if err != nil {
		return
	}

	c.ASF[0] |= ad9910Write
	c.ASF[3] = 0xff
	c.ASF[4] = 0xfc

	err = d.connection.Tx(c.ASF[:], make([]byte, len(c.ASF)))
	if err != nil {
		return
	}

	c.FTW[0] |= ad9910Write
	c.FTW[1] = 0x19
	c.FTW[2] = 0x99
	c.FTW[3] = 0x99
	c.FTW[4] = 0x9a

	err = d.connection.Tx(c.FTW[:], make([]byte, len(c.FTW)))
	if err != nil {
		return
	}

	c.STProfile0[0] |= ad9910Write
	c.STProfile0[1] = 0xff
	c.STProfile0[4] = 0x19
	c.STProfile0[5] = 0x99
	c.STProfile0[6] = 0x99
	c.STProfile0[7] = 0x9a

	return d.connection.Tx(c.STProfile0[:], make([]byte, len(c.STProfile0)))
}

// SPI default paramaeters.
const (
	AD9910SPIDefaultBus   = 2
	AD9910SPIDefaultMode  = 0
	AD9910SPIDefaultSpeed = 50e6
)

// Start initializes the driver.
func (d *AD9910) Start() (err error) {
	d.connection, err = d.connector.GetSpiConnection(
		AD9910SPIDefaultBus,
		AD9910SPIDefaultMode,
		AD9910SPIDefaultSpeed)

	return
}

// Halt stops the driver.
func (d *AD9910) Halt() (err error) {
	return d.connection.Close()
}

func frequencyToFTW(frequency float64) uint32 {
	return uint32(math.Round(frequency) / AD9910DefaultSYSCLK * (1 << 32))
}

func amplitudeToASF(amplitude float64) uint16 {
	return uint16(math.Round(amplitude * (1 << 14)))
}
