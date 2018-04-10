package main

import (
	"flag"
	"log"

	"periph.io/x/periph/host"

	"github.com/bodokaiser/beagle/driver/dds"
	"github.com/bodokaiser/beagle/driver/misc"
)

type config struct {
	cselect   uint
	frequency float64
	amplitude float64
}

func main() {
	c := config{}

	flag.UintVar(&c.cselect, "select", 0, "DDS chip to configure")
	flag.Float64Var(&c.frequency, "frequency", 200e6, "Frequency [Hz]")
	flag.Float64Var(&c.amplitude, "amplitude", 0, "Amplitude [0, 1]")
	flag.Parse()

	err := misc.DefaultConfig.Exec()
	if err != nil {
		log.Fatal(err)
	}

	_, err = host.Init()
	if err != nil {
		log.Fatal(err)
	}

	s, err := misc.NewSelect()
	if err != nil {
		log.Fatal(err)
	}

	err = s.Address(c.cselect)
	if err != nil {
		log.Fatal(err)
	}

	d, err := dds.NewAD9910(dds.DefaultSysClock, dds.DefaultRefClock)
	if err != nil {
		log.Fatal(err)
	}

	err = d.SingleTone(c.frequency)
	if err != nil {
		log.Fatal(err)
	}
}
