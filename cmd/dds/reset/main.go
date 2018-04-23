// cli controls a DDS array via command line.
package main

import (
	"flag"
	"log"

	"periph.io/x/periph/host"

	"github.com/bodokaiser/houston/config"
	"github.com/bodokaiser/houston/driver/dds/ad99xx"
	"github.com/bodokaiser/houston/driver/mux"
	"github.com/bodokaiser/houston/model"
)

func main() {
	c := config.Config{}
	d := model.DDSDevice{}

	flag.Var(&c, "config", "path to config file")
	flag.UintVar(&d.ID, "select", 0, "address to select")
	flag.Parse()

	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	sel := mux.NewDigital(c.GPIO.Mux)
	err = sel.Init()
	if err != nil {
		log.Fatal(err)
	}

	dev := ad99xx.NewAD9910(ad99xx.Config{
		SysClock:  c.SysClock,
		RefClock:  c.RefClock,
		ResetPin:  c.GPIO.Reset,
		UpdatePin: c.GPIO.Update,
		SPI:       c.SPI,
	})

	err = dev.Init()
	if err != nil {
		log.Fatal(err)
	}

	err = sel.Select(uint8(d.ID))
	if err != nil {
		log.Fatal(err)
	}

	err = dev.Reset()
	if err != nil {
		log.Fatal(err)
	}
}
