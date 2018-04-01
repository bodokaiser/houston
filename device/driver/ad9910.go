package driver

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/spi"

	"github.com/bodokaiser/beagle/device/register"
)

// AD9910 communication ops.
const (
	AD9910Write byte = 0x80
	AD9910Read       = 0x00
)

// AD9910 register addresses.
const (
	AD9910CtrlFunc1 = iota
	AD9910CtrlFunc2
	AD9910CtrlFunc3
	AD9910AuxDACCtrl
	AD9910IOUpdateRate
	_
	_
	AD9910FreqTuningWord
	AD9910PhaseOffsetWord
	AD9910AmplScaleFactor
	AD9910MultichipSync
	AD9910DigitalRampLimit
	AD9910DigitalRampStepSize
	AD9910DigitalRampRate
	AD9910Profile0
	AD9910Profile1
	AD9910Profile2
	AD9910Profile3
	AD9910Profile4
	AD9910Profile5
	AD9910Profile6
	AD9910Profile7
	AD9910RAM
)

// AD9910 is the gobot driver for the AD9910 DDS.
type AD9910 struct {
	name       string
	connector  spi.Connector
	connection spi.Connection
	registers  *register.AD9910
}

// NewAD9910 returns an initialized AD9910.
func NewAD9910(c spi.Connector) *AD9910 {
	r := new(register.AD9910)
	r.CtrlFunc1Data = register.AD9910CtrlFunc1Default
	r.CtrlFunc2Data = register.AD9910CtrlFunc2Default
	r.CtrlFunc3Data = register.AD9910CtrlFunc3Default
	r.AuxDACCtrlData = register.AD9910AuxDACCtrlDefault
	r.StProfile1Data = register.AD9910StProfile0Default

	return &AD9910{
		name:      gobot.DefaultName("AD9910"),
		connector: c,
		registers: r,
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

// RunSingleTone configures AD9910 to run in single tone mode with amplite,
// frequency and phase.
func (d *AD9910) RunSingleTone(ampl uint16, freq uint32, phase uint16) error {
	d.registers.SetRAMEnable(false)
	d.registers.SetDigitalRampEnable(false)
	d.registers.SetParallelPortEnable(false)

	d.registers.SetStProfile0AmplScaleFactor(ampl)
	d.registers.SetStProfile0FreqTuningWord(freq)
	d.registers.SetStProfile0PhaseOffsetWord(phase)

	instructions := [][]byte{
		append([]byte{AD9910CtrlFunc1 | AD9910Write}, d.registers.CtrlFunc1Data[:]...),
		append([]byte{AD9910CtrlFunc2 | AD9910Write}, d.registers.CtrlFunc2Data[:]...),
		append([]byte{AD9910CtrlFunc3 | AD9910Write}, d.registers.CtrlFunc3Data[:]...),
	}

	w := []byte{}
	for _, i := range instructions {
		w = append(w, i...)
	}
	r := make([]byte, len(w))

	return d.connection.Tx(w, r)
}

// Start initializes the driver.
func (d *AD9910) Start() (err error) {
	bus := d.connector.GetSpiDefaultBus()
	mode := d.connector.GetSpiDefaultMode()
	speed := d.connector.GetSpiDefaultMaxSpeed()

	d.connection, err = d.connector.GetSpiConnection(bus, mode, speed)

	return err
}

// Halt stops the driver.
func (d *AD9910) Halt() (err error) {
	return d.connection.Close()
}
