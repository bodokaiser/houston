package config

import (
	"bytes"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/bodokaiser/houston/driver/dds"
	"github.com/bodokaiser/houston/driver/mux"
)

type Config struct {
	Debug bool
	DDS   dds.Config `yaml:"dds"`
	Mux   mux.Config `yaml:"mux"`
}

func (c *Config) Ensure() {
	c.DDS.Debug = c.Debug
	c.Mux.Debug = c.Debug
}

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
