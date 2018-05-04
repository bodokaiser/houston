// Package mux provides drivers for hardware multiplexers.
package mux

import "github.com/bodokaiser/houston/driver"

// Config holds configuration for multiplexer devices.
type Config struct {
	GPIO  GPIOConfig `yaml:"gpio"`
	Debug bool       `yaml:"debug"`
}

// GPIOConfig defines the GPIO pines to use as multiplexer lines.
type GPIOConfig struct {
	CS []string `yaml:"cs"`
}

// Mux implements a hardware multiplexer driver.
type Mux interface {
	driver.Driver
	// Select configures the multiplexer to address given address
	Select(uint8) error
}
