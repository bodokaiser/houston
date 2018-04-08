package main

import (
	"flag"

	"gobot.io/x/gobot/platforms/beaglebone"

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

	beagle := beaglebone.NewAdaptor()

	cs := misc.NewChipSelect(beagle)
	cs.Start()
	cs.Select(0)

	io := misc.NewIOControl(beagle)
	io.Start()

	dds0 := dds.NewAD9910(beagle)
	dds0.Start()
	dds0.RunSingleTone()

	io.Update()
}
