// Package mux provides drivers for hardware multiplexers.
package mux

// Mux implements a hardware multiplexer driver.
type Mux interface {
	// Select configures the multiplexer to address given address
	Select(uint8) error
}
