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
	d := model.DDSDevice{
		Amplitude: model.DDSParam{
			DDSConst: &model.DDSConst{},
		},
		Frequency: model.DDSParam{
			DDSConst: &model.DDSConst{},
		},
		PhaseOffset: model.DDSParam{
			DDSConst: &model.DDSConst{},
		},
	}

	flag.UintVar(&d.ID, "select", 0, "address to select")
	flag.Float64Var(&d.Frequency.Value, "frequency", 10e6, "Frequency [0, 400e6]")
	flag.Float64Var(&d.Amplitude.Value, "amplitude", 1.0, "Amplitude [0, 1]")
	flag.Float64Var(&d.PhaseOffset.Value, "phase", 0.0, "Phase [0, 2π]")
	flag.Var(&c, "config", "path to config file")
	flag.Parse()

	log.Printf("config:\n%s\n", c.Render())

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

	err = sel.Select(uint8(d.ID))
	if err != nil {
		log.Fatal(err)
	}

	dev.SetAmplitude(d.Amplitude.Value)
	dev.SetFrequency(d.Frequency.Value)
	dev.SetPhaseOffset(d.PhaseOffset.Value)

	err = dev.Exec()
	if err != nil {
		log.Fatal(err)
	}
}
