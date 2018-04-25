// cli controls a DDS array via command line.
package main

import (
	"flag"
	"log"
	"time"

	"periph.io/x/periph/host"

	"github.com/bodokaiser/houston/cmd"
	"github.com/bodokaiser/houston/driver/dds"
	"github.com/bodokaiser/houston/driver/dds/ad99xx"
	"github.com/bodokaiser/houston/driver/mux"
	"github.com/bodokaiser/houston/model"
)

func main() {
	c := cmd.Config{}
	d := model.DDSDevice{
		Frequency: model.DDSParam{
			DDSSweep: &model.DDSSweep{
				Limits:   [2]float64{},
				NoDwells: [2]bool{},
			},
		},
		Amplitude: model.DDSParam{
			DDSConst: &model.DDSConst{},
		},
		PhaseOffset: model.DDSParam{
			DDSConst: &model.DDSConst{},
		},
	}

	flag.UintVar(&d.ID, "select", 0,
		"Address of DDS chip")
	flag.Float64Var(&d.Amplitude.Value, "amplitude", 1.0,
		"Constant amplitude in relative units")
	flag.Float64Var(&d.PhaseOffset.Value, "phase", 0.0,
		"Phase offset in Radiants")
	flag.Float64Var(&d.Frequency.Limits[0], "start", 12e5,
		"Start value for frequency in Hertz")
	flag.Float64Var(&d.Frequency.Limits[1], "stop", 16e5,
		"Stop value for frequency in Hertz")
	flag.DurationVar(&d.Frequency.DDSSweep.Duration, "duration", time.Second,
		"Ramp Duration in Seconds")
	flag.BoolVar(&d.Frequency.NoDwells[0], "no-dwell-low", true,
		"ramp does not remain at lower end")
	flag.BoolVar(&d.Frequency.NoDwells[1], "no-dwell-high", true,
		"ramp does not remain at upper end")
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

	err = sel.Select(uint8(d.ID))
	if err != nil {
		log.Fatal(err)
	}

	dev.SetAmplitude(d.Amplitude.Value)
	dev.SetPhaseOffset(d.PhaseOffset.Value)
	dev.Sweep(dds.SweepConfig{
		Limits:   d.Frequency.Limits,
		NoDwells: d.Frequency.NoDwells,
		Duration: d.Frequency.DDSSweep.Duration,
		Param:    dds.ParamFrequency,
	})

	err = dev.WriteToDev()
	if err != nil {
		log.Fatal(err)
	}
	err = dev.Update()
	if err != nil {
		log.Fatal(err)
	}
}
