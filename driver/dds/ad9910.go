package dds

import (
	"math"

	"github.com/bodokaiser/beagle/driver/misc"
	"periph.io/x/periph/conn/spi"
	"periph.io/x/periph/conn/spi/spireg"
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
	AD9910DefaultREFCLK = 1e7 // 10 MHz
	AD9910DefaultSYSCLK = 1e9 // 1 GHz
)

// AD9910 is the SPI driver for the AD9910 DDS chip.
type AD9910 struct {
	conn spi.Conn
}

// NewAD9910 returns a new AD9910 device.
func NewAD9910() *AD9910 {
	return &AD9910{}
}

// Init initializes the AD9910 driver.
func (d *AD9910) Init() error {
	spi, err := spireg.Open("SPI1.0")
	if err != nil {
		return err
	}

	d.conn, err = spi.Connect(5e6, 0, 8)

	return err
}

// RunSingleTone configures the AD9910 to run in single tone mode.
func (d *AD9910) RunSingleTone(frequency float64) (err error) {
	c := misc.NewControl()

	err = c.Init()
	if err != nil {
		return
	}

	w := []byte{
		0, 0, 0, 0, 2, 2, 31, 63, 64, 0, 1, 0, 64, 8, 32, 11, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 127, 8, 0, 0, 8, 0, 0, 9, 0, 0, 0, 0, 12, 0, 0, 0, 0, 0, 0, 0, 0, 15, 0, 0, 0, 0, 0, 0, 0, 0, 14, 0, 0, 0, 0, 0, 0, 0, 0, 13, 0, 0, 0, 0, 7, 0, 0, 0, 0, 4, 255, 255, 255, 255, 15, 0, 0, 0, 0, 0, 0, 0, 0, 14, 8, 181, 0, 0, 0, 0, 0, 0, 17, 0, 0, 0, 0, 0, 0, 0, 0, 16, 0, 0, 0, 0, 0, 0, 0, 0, 19, 0, 0, 0, 0, 0, 0, 0, 0, 18, 0, 0, 0, 0, 0, 0, 0, 0, 21, 0, 0, 0, 0, 0, 0, 0, 0, 20, 0, 0, 0, 0, 0, 0, 0, 0,
	}
	r := make([]byte, len(w))

	err = d.conn.Tx(w, r)
	if err != nil {
		return
	}
	err = c.IOUpdate()
	if err != nil {
		return
	}

	w = []byte{0, 0, 128, 2, 2, 2, 29, 63, 65, 200, 1, 1, 64, 8, 32}
	r = make([]byte, len(w))

	err = d.conn.Tx(w, r)
	if err != nil {
		return
	}
	err = c.IOUpdate()
	if err != nil {
		return
	}

	w = []byte{
		8, 0, 0, 1, 3, 64, 8, 32, 14, 63, 255, 0, 0, 51, 51, 51, 51,
		0, 0, 128, 2, 2, 9, 0, 0, 255, 252, 7, 51, 51, 51, 51,
	}
	r = make([]byte, len(w))

	err = d.conn.Tx(w, r)
	if err != nil {
		return
	}
	return c.IOUpdate()
}

func frequencyToFTW(frequency float64) uint32 {
	return uint32(math.Round(frequency) / AD9910DefaultSYSCLK * (1 << 32))
}

func amplitudeToASF(amplitude float64) uint16 {
	return uint16(math.Round(amplitude * (1 << 14)))
}
