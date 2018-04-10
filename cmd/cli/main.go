package main

import (
	"flag"
	"log"

	"periph.io/x/periph/host"

	"github.com/bodokaiser/beagle/driver/dds"
	"github.com/bodokaiser/beagle/driver/misc"
)

type config struct {
	frequency float64
	amplitude float64
}

func main() {
	c := config{}

	flag.Float64Var(&c.frequency, "frequency", 200e6, "")
	flag.Float64Var(&c.amplitude, "amplitude", 0, "")
	flag.Parse()

	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	s, err := misc.NewSelect()
	if err != nil {
		log.Fatal(err)
	}
	s.Address(0)

	d, err := dds.NewAD9910(dds.DefaultSysClock, dds.DefaultRefClock)
	if err != nil {
		log.Fatal(err)
	}

	err = d.SingleTone(c.frequency)
	if err != nil {
		log.Fatal(err)
	}
}
