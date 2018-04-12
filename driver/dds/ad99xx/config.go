package ad99xx

import "periph.io/x/periph/conn/spi"

// Config is the config for the AD99xx family drivers.
type Config struct {
	SysClock    float64
	RefClock    float64
	ResetPin    string
	IOUpdatePin string
	SPIDevice   string
	SPIMaxFreq  int64
	SPIMode     spi.Mode
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
