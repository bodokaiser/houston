package ad9910

import (
	"bytes"
	"math"
	"testing"
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpiotest"
	"periph.io/x/periph/conn/spi/spitest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/bodokaiser/houston/driver/dds"
)

type AD9910TestSuite struct {
	suite.Suite

	buffer    *bytes.Buffer
	spiPort   *spitest.RecordRaw
	updatePin *gpiotest.Pin
	resetPin  *gpiotest.Pin

	driver *AD9910
}

func (s *AD9910TestSuite) SetupSuite() {
	s.updatePin = &gpiotest.Pin{N: "RESET", Num: 0}
	s.resetPin = &gpiotest.Pin{N: "UPDATE", Num: 1}
}

func (s *AD9910TestSuite) SetupTest() {
	s.buffer = &bytes.Buffer{}
	s.spiPort = spitest.NewRecordRaw(s.buffer)

	c := dds.Config{
		ResetPin:  s.resetPin,
		UpdatePin: s.updatePin,
		SPIPort:   s.spiPort,
	}
	c.SysClock = 1e9
	c.RefClock = 1e7
	c.SPIMaxTxSize = 256

	s.driver = NewAD9910(c)
}

func (s *AD9910TestSuite) TestInit() {
	assert.NoError(s.T(), s.driver.Init())
	assert.Equal(s.T(), gpio.Low, s.updatePin.Read())
	assert.Equal(s.T(), gpio.Low, s.resetPin.Read())
}

func (s *AD9910TestSuite) TestInitErrors() {
	d1 := NewAD9910(dds.Config{
		ResetPin:  s.resetPin,
		UpdatePin: s.updatePin,
	})
	assert.EqualError(s.T(), d1.Init(), "failed to find SPI port")

	d2 := NewAD9910(dds.Config{
		UpdatePin: s.updatePin,
		SPIPort:   s.spiPort,
	})
	assert.EqualError(s.T(), d2.Init(), "failed to find reset GPIO pin")

	d3 := NewAD9910(dds.Config{
		ResetPin: s.resetPin,
		SPIPort:  s.spiPort,
	})
	assert.EqualError(s.T(), d3.Init(), "failed to find update GPIO pin")
}

func (s *AD9910TestSuite) TestMaxTxSize() {
	assert.NoError(s.T(), s.driver.Init())
	assert.Equal(s.T(), 256, s.driver.MaxTxSize())
}

func (s *AD9910TestSuite) TestExecDefaults() {
	assert.NoError(s.T(), s.driver.Init())
	assert.NoError(s.T(), s.driver.Exec())

	assert.Equal(s.T(), []byte{
		// CFR1
		0x00, 0x00, 0x00, 0x00, 0x00,
		// CFR2
		0x01, 0x01, 0x40, 0x08, 0x20,
		// CFR3
		0x02, 0x1f, 0x3f, 0x40, 0x00,
		// AuxDAC
		0x03, 0x00, 0x00, 0x00, 0x7f,
		// IOUpdateRate
		0x04, 0xff, 0xff, 0xff, 0xff,
		// STProfile0
		0x0e, 0x08, 0xb5, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}, s.buffer.Bytes())
}

func (s *AD9910TestSuite) TestExecSingleTone1() {
	s.driver.SetAmplitude(1.0)
	s.driver.SetFrequency(100e6)
	s.driver.SetPhaseOffset(0)

	assert.NoError(s.T(), s.driver.Init())
	assert.NoError(s.T(), s.driver.Exec())

	assert.Equal(s.T(), []byte{
		// CFR1
		0x00, 0x00, 0x00, 0x00, 0x00,
		// CFR2
		0x01, 0x01, 0x40, 0x08, 0x20,
		// CFR3
		0x02, 0x1f, 0x3f, 0x40, 0x00,
		// AuxDAC
		0x03, 0x00, 0x00, 0x00, 0x7f,
		// IOUpdateRate
		0x04, 0xff, 0xff, 0xff, 0xff,
		// FTW
		0x07, 0x19, 0x99, 0x99, 0x9b,
		// ASF
		0x09, 0x00, 0x00, 0xff, 0xfc,
		// STProfile0
		0x0e, 0x3f, 0xff, 0x00, 0x00, 0x19, 0x99, 0x99, 0x9b,
	}, s.buffer.Bytes())
}

func (s *AD9910TestSuite) TestExecSingleTone2() {
	s.driver.SetAmplitude(1.0)
	s.driver.SetFrequency(100e6)
	s.driver.SetPhaseOffset(math.Pi / 2)

	assert.NoError(s.T(), s.driver.Init())
	assert.NoError(s.T(), s.driver.Exec())

	assert.Equal(s.T(), []byte{
		// CFR1
		0x00, 0x00, 0x00, 0x00, 0x00,
		// CFR2
		0x01, 0x01, 0x40, 0x08, 0x20,
		// CFR3
		0x02, 0x1f, 0x3f, 0x40, 0x00,
		// AuxDAC
		0x03, 0x00, 0x00, 0x00, 0x7f,
		// IOUpdateRate
		0x04, 0xff, 0xff, 0xff, 0xff,
		// FTW
		0x07, 0x19, 0x99, 0x99, 0x9b,
		// POW
		0x08, 0x40, 0x00,
		// ASF
		0x09, 0x00, 0x00, 0xff, 0xfc,
		// STProfile0
		0x0e, 0x3f, 0xff, 0x40, 0x00, 0x19, 0x99, 0x99, 0x9b,
	}, s.buffer.Bytes())
}

func (s *AD9910TestSuite) TestExecSweep() {
	s.driver.SetAmplitude(1.0)
	s.driver.SetSweep(dds.SweepConfig{
		Limits:   [2]float64{100e6, 200e6},
		Duration: 2 * time.Second,
		NoDwells: [2]bool{true, true},
		Param:    dds.ParamFrequency,
	})
	s.driver.SetPhaseOffset(0)

	assert.NoError(s.T(), s.driver.Init())
	assert.NoError(s.T(), s.driver.Exec())

	assert.Equal(s.T(), []byte{
		// CFR1
		0x00, 0x00, 0x00, 0x00, 0x00,
		// CFR2
		0x01, 0x01, 0x4e, 0x08, 0x20,
		// CFR3
		0x02, 0x1f, 0x3f, 0x40, 0x00,
		// AuxDAC
		0x03, 0x00, 0x00, 0x00, 0x7f,
		// IOUpdateRate
		0x04, 0xff, 0xff, 0xff, 0xff,
		// ASF
		0x09, 0x00, 0x00, 0xff, 0xfc,
		// STProfile0
		0x0e, 0x3f, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		// DRampLimit
		0x0b, 0x33, 0x33, 0x33, 0x34, 0x19, 0x99, 0x99, 0x9b,
		// DRampStep
		0x0c, 0x00, 0x00, 0x33, 0xce, 0x00, 0x00, 0x33, 0xce,
		// DRampRate
		0x0d, 0x3c, 0x4f, 0x3c, 0x4f,
	}, s.buffer.Bytes())
}

func (s *AD9910TestSuite) TestExecPlayback() {
	s.driver.SetAmplitude(1.0)
	s.driver.SetPlayback(dds.PlaybackConfig{
		Data:     []float64{1.00, 0.75, 0.50, 0.25, 0.00},
		Interval: 200 * time.Nanosecond,
		Duplex:   false,
		Param:    dds.ParamAmplitude,
	})
	s.driver.SetPhaseOffset(0)

	assert.NoError(s.T(), s.driver.Init())
	assert.NoError(s.T(), s.driver.Exec())

	assert.Equal(s.T(), []byte{
		// CFR1
		0x00, 0xc0, 0x00, 0x00, 0x00,
		// CFR2
		0x01, 0x00, 0x40, 0x08, 0x20,
		// CFR3
		0x02, 0x1f, 0x3f, 0x40, 0x00,
		// AuxDAC
		0x03, 0x00, 0x00, 0x00, 0x7f,
		// IOUpdateRate
		0x04, 0xff, 0xff, 0xff, 0xff,
		// ASF
		0x09, 0x00, 0x00, 0xff, 0xfc,
		// RAMProfile0
		0x0e, 0x00, 0x00, 0x32, 0x01, 0x00, 0x00, 0x00, 0x24,
		// RAM
		0x16,
	}, s.buffer.Bytes())
}

func TestAD9910Suite(t *testing.T) {
	suite.Run(t, new(AD9910TestSuite))
}
