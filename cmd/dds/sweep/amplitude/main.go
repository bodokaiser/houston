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
		Amplitude: model.DDSParam{
			DDSSweep: &model.DDSSweep{
				Limits:   [2]float64{},
				NoDwells: [2]bool{},
			},
		},
		Frequency: model.DDSParam{
			DDSConst: &model.DDSConst{},
		},
		PhaseOffset: model.DDSParam{
			DDSConst: &model.DDSConst{},
		},
	}

	flag.UintVar(&d.ID, "select", 0,
		"Address of DDS chip")
	flag.Float64Var(&d.Frequency.Value, "frequency", 10e6,
		"Constant frequency in Hertz")
	flag.Float64Var(&d.Amplitude.Limits[0], "start", 0.5,
		"Start value for amplitude in relative units")
	flag.Float64Var(&d.Amplitude.Limits[1], "stop", 1.0,
		"Stop value for amplitude in relative units")
	flag.BoolVar(&d.Amplitude.NoDwells[0], "no-dwell-low", true,
		"ramp does not remain at lower end")
	flag.BoolVar(&d.Amplitude.NoDwells[1], "no-dwell-high", true,
		"ramp does not remain at upper end")
	flag.DurationVar(&d.Amplitude.DDSSweep.Duration, "duration", time.Second,
		"Ramp Duration in Seconds")
	flag.Float64Var(&d.PhaseOffset.Value, "phase", 0.0,
		"Phase offset in Radiants")
	flag.StringVar(&c.Filename, "config", "config.yaml", "path to config file")
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

	dev.Sweep(dds.SweepConfig{
		Limits:   d.Amplitude.Limits,
		NoDwells: d.Amplitude.NoDwells,
		Duration: d.Amplitude.DDSSweep.Duration,
		Param:    dds.ParamAmplitude,
	})
	dev.SetFrequency(d.Frequency.Value)
	dev.SetPhaseOffset(d.PhaseOffset.Value)

	err = dev.Exec()
	if err != nil {
		log.Fatal(err)
	}
}
