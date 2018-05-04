package spi

// Config holds SPI configuration.
type Config struct {
	Device  string
	MaxFreq int64
	Duplex  bool
}
