package ad99xx

import "periph.io/x/periph/conn/spi"

// Config contains the configuration for the AD99xx family drivers.
type Config struct {
	// SysClock is given by reference clock / divider.
	SysClock float64
	// RefClock is the external clock signal provided to the AD9xx.
	RefClock float64
	// ResetPin is the digital pin used to trigger resets.
	ResetPin string
	// IOUpdatePin is the digital pin used to trigger I/O updates.
	IOUpdatePin string
	// SPIDevice is the SPI chip and bus to use as serial connection.
	SPIDevice string
	// SPIMaxFreq is the maximum frequency in Hz to run the serial connection with.
	SPIMaxFreq int64
	// SPIMode is the SPI mode. See spi.Mode for details.
	SPIMode spi.Mode
}

// AD9910DefaultConfig is the default configuration for the AD9910 dds.
var AD9910DefaultConfig = Config{
	SysClock:    1e9,
	RefClock:    1e7,
	ResetPin:    "65",
	IOUpdatePin: "27",
	SPIDevice:   "SPI1.0",
	SPIMaxFreq:  5e6,
	SPIMode:     0,
}
