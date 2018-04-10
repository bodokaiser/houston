package misc

import (
	"os/exec"
)

// DefaultConfig enables the first SPI bus.
var DefaultConfig = Config{
	"p9.17": "spi_cs",
	"p9.21": "spi",
	"p9.18": "spi",
	"p9.22": "spi_sclk",
}

// Config describes a pin layout configuration for the capemanager.
type Config map[string]string

// Exec executes the shell commadn to configure the pin layout.
func (c Config) Exec() error {
	for pin, mode := range c {
		err := exec.Command("/usr/bin/config-pin", pin, mode).Run()
		if err != nil {
			return err
		}
	}

	return nil
}
