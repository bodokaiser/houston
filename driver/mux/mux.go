// Package mux provides drivers for hardware multiplexers.
package mux

import "github.com/bodokaiser/houston/driver"

type Config struct {
	GPIO GPIOConfig `yaml:"gpio"`
}

type GPIOConfig struct {
	CS []string `yaml:"cs"`
}

// Mux implements a hardware multiplexer driver.
type Mux interface {
	driver.Driver
	// Select configures the multiplexer to address given address
	Select(uint8) error
}
