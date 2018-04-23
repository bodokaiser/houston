package spi

import "periph.io/x/periph/conn/spi"

type Config struct {
	Device  string
	MaxFreq int64
	Mode    spi.Mode
}
