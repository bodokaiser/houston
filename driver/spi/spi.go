package spi

import "periph.io/x/periph/conn/spi"

type Config struct {
	// SPIDevice is the SPI chip and bus to use as serial connection.
	SPIDevice string
	// SPIMaxFreq is the maximum frequency in Hz to run the serial connection with.
	SPIMaxFreq int64
	// SPIMode is the SPI mode. See spi.Mode for details.
	SPIMode spi.Mode
}
