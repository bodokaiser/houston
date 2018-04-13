// cli controls a DDS array via command line.
package main

import (
	"flag"
	"log"

	"periph.io/x/periph/host"

	"github.com/bodokaiser/houston/driver/dds/ad99xx"
	"github.com/bodokaiser/houston/driver/mux"
)

const (
	defaultSysClock = 1e9
	defaultRefClock = 1e7
)

const (
	defaultSPIDevice  = "SPI1.0"
	defaultSPIMaxFreq = 5e6
	defaultSPIMode    = 0
)

const (
	defaultResetPin    = "65"
	defaultIOUpdatePin = "27"
)

var defaultMuxPins = []string{"48", "30", "60", "31", "50"}

type config struct {
	cselect   uint
	frequency float64
	amplitude float64
	phaseOff  float64
}

func main() {
	c := config{}

	flag.UintVar(&c.cselect, "select", 0, "address to select")
	flag.Float64Var(&c.frequency, "frequency", 200e6, "frequency in Hz")
	flag.Float64Var(&c.amplitude, "amplitude", 1.0, "amplitude from 0.0 to 1.0")
	flag.Float64Var(&c.phaseOff, "phase", 0.0, "phase offset from 0 to 2π")
	flag.Parse()

	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	csel, err := mux.NewDigital(defaultMuxPins)
	if err != nil {
		log.Fatal(err)
	}

	dds, err := ad99xx.NewAD9910(ad99xx.Config{
		SysClock:    defaultSysClock,
		RefClock:    defaultRefClock,
		ResetPin:    defaultResetPin,
		IOUpdatePin: defaultIOUpdatePin,
		SPIDevice:   defaultSPIDevice,
		SPIMaxFreq:  defaultSPIMaxFreq,
		SPIMode:     defaultSPIMode,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = csel.Select(uint8(c.cselect))
	if err != nil {
		log.Fatal(err)
	}

	err = dds.SingleTone(c.amplitude, c.frequency, c.phaseOff)
	if err != nil {
		log.Fatal(err)
	}
}
