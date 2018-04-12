// cli controls a DDS array via command line.
package main

import (
	"flag"
	"log"

	"periph.io/x/periph/host"

	"github.com/bodokaiser/houston/driver/dds/ad99xx"
	"github.com/bodokaiser/houston/driver/mux"
)

type config struct {
	cselect   uint
	frequency float64
	amplitude float64
}

func main() {
	c := config{}

	flag.UintVar(&c.cselect, "select", 0, "address to select")
	flag.Float64Var(&c.frequency, "frequency", 200e6, "frequency in Hz")
	flag.Float64Var(&c.amplitude, "amplitude", 1.0, "amplitude from 0.0 to 1.0")
	flag.Parse()

	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	csel, err := mux.NewDigital(mux.DefaultDigitalPins)
	if err != nil {
		log.Fatal(err)
	}

	dds, err := ad99xx.NewAD9910(ad99xx.AD9910DefaultConfig)
	if err != nil {
		log.Fatal(err)
	}

	err = csel.Select(uint8(c.cselect))
	if err != nil {
		log.Fatal(err)
	}

	err = dds.SingleTone(c.frequency)
	if err != nil {
		log.Fatal(err)
	}
}
