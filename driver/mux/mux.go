// Package mux provides drivers for hardware multiplexers.
package mux

// Mux implements a hardware multiplexer driver.
type Mux interface {
	// Select configures the multiplexer to address given address
	Select(uint8) error
}

// DefaultDigitalPins are the GPIO pins to use for the digital multiplexer.
var DefaultDigitalPins = []string{"48", "30", "60", "31", "50"}
