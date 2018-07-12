// Package config provides global configuration.
package config

import (
	"bytes"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/gobuffalo/packr"

	"github.com/bodokaiser/houston/driver/dds"
	"github.com/bodokaiser/houston/driver/mux"
)

// Config combines different device configs into a single type.
type Config struct {
	Debug bool
	DDS   DDSConfig  `yaml:"dds"`
	Mux   mux.Config `yaml:"mux"`
}

type DDSConfig struct {
	dds.Config `yaml:",inline"`
	SPI        SPIConfig  `yaml:"spi"`
	GPIO       GPIOConfig `yaml:"gpio"`
}

type SPIConfig struct {
	Device string `yaml:"device"`
}

type GPIOConfig struct {
	Reset  string `yaml:"reset"`
	Update string `yaml:"update"`
}

// Ensure ensures configurations consistency.
//
// In particular we just pass on the top level debug flag into the
// nested configs.
func (c *Config) Ensure() {
	c.DDS.Debug = c.Debug
	c.Mux.Debug = c.Debug
}

// ReadFromFile reads a configuration from file.
//
// At this point we only support YAML formatted files.
func (c *Config) ReadFromFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(f); err != nil {
		return err
	}

	return yaml.Unmarshal(buf.Bytes(), c)
}

// ReadFromBox reads a configuration from a box container.
//
// At this point we only support YAML formatted files.
func (c *Config) ReadFromBox(filename string) error {
	b := packr.NewBox("./").Bytes("beagle.yaml")

	return yaml.Unmarshal(b, c)
}
