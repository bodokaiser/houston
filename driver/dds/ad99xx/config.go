package ad99xx

import "github.com/bodokaiser/houston/driver/spi"

// Config contains the configuration for the AD99xx family drivers.
type Config struct {
	spi.Config

	// SysClock is given by reference clock / divider.
	SysClock uint32
	// RefClock is the external clock signal provided to the AD9xx.
	RefClock uint32
	// ResetPin is the digital pin used to trigger resets.
	ResetPin string
	// IOUpdatePin is the digital pin used to trigger I/O updates.
	IOUpdatePin string
}
