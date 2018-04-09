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

	csel := misc.NewSelect()
	ctrl := misc.NewControl()
	fgen := dds.NewAD9910()

	err = csel.Init()
	if err != nil {
		log.Fatal(err)
	}
	err = ctrl.Init()
	if err != nil {
		log.Fatal(err)
	}
	err = fgen.Init()
	if err != nil {
		log.Fatal(err)
	}

	err = fgen.RunSingleTone(c.frequency)
	if err != nil {
		log.Fatal(err)
	}

	err = ctrl.IOUpdate()
	if err != nil {
		log.Fatal(err)
	}
}
