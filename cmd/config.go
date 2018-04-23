package cmd

import (
	"bytes"
	"io"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/bodokaiser/houston/driver/dds"
	"github.com/bodokaiser/houston/driver/mux"
)

type Config struct {
	DDS dds.Config `yaml:"dds"`
	Mux mux.Config `yaml:"mux"`
}

func (c *Config) ReadFrom(r io.Reader) (n int64, err error) {
	buf := new(bytes.Buffer)
	n, err = buf.ReadFrom(r)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(buf.Bytes(), c)

	return
}

func (c *Config) ReadFromFile(p string) error {
	f, err := os.Open(p)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = c.ReadFrom(f)

	return err
}

func (c *Config) Set(s string) error {
	_, err := os.Stat(s)
	if err != nil {
		return err
	}

	return c.ReadFromFile(s)
}

func (c *Config) String() string {
	return "config.yaml"
}

func (c *Config) Render() string {
	b, err := yaml.Marshal(c)
	if err != nil {
		return ""
	}
	return string(b)
}
