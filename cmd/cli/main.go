package main

import (
	"flag"

	"gobot.io/x/gobot/platforms/beaglebone"

	"github.com/bodokaiser/beagle/device/driver"
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

	cs := driver.NewChipSelect(beagle)
	cs.Start()
	cs.Select(0)

	io := driver.NewControl(beagle)
	io.Start()

	dds0 := driver.NewAD9910(beagle)
	dds0.Start()
	dds0.RunSingleTone(1<<15, 1<<30, 0)

	io.IOUpdate()
}
