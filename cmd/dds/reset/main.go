// cli controls a DDS array via command line.
package main

import (
	"flag"
	"log"

	"periph.io/x/periph/host"

	"github.com/bodokaiser/houston/cmd"
	"github.com/bodokaiser/houston/driver/dds/ad99xx"
	"github.com/bodokaiser/houston/driver/mux"
	"github.com/bodokaiser/houston/model"
)

func main() {
	c := cmd.Config{}
	d := model.DDSDevice{}

	flag.StringVar(&c.Filename, "config", "config.yaml", "path to config file")
	flag.UintVar(&d.ID, "select", 0, "address to select")
	flag.Parse()

	err := c.ReadFromFile()
	if err != nil {
		log.Fatal(err)
	}

	_, err = host.Init()
	if err != nil {
		log.Fatal(err)
	}

	sel := mux.NewDigital(c.Mux)
	err = sel.Init()
	if err != nil {
		log.Fatal(err)
	}

	dev := ad99xx.NewAD9910(c.DDS)
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
