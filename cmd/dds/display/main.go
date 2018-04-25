package main

import (
	"flag"
	"log"

	"github.com/bodokaiser/houston/cmd"
	"github.com/bodokaiser/houston/driver/dds/ad99xx"
	"github.com/bodokaiser/houston/driver/mux"
	"github.com/bodokaiser/houston/model"
	"periph.io/x/periph/host"
)

func main() {
	c := cmd.Config{}
	m := model.DDSDevice{}

	flag.UintVar(&m.ID, "select", 0, "address to select")
	flag.Var(&c, "config", "path to config file")
	flag.Parse()

	_, err := host.Init()
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

	err = sel.Select(uint8(m.ID))
	if err != nil {
		log.Fatal(err)
	}

	err = dev.ReadFromDev()
	if err != nil {
		log.Fatal(err)
	}

	err = dev.Update()
	if err != nil {
		log.Fatal(err)
	}

	err = dev.ReadFromDev()
	if err != nil {
		log.Fatal(err)
	}
}
