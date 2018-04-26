// cli controls a DDS array via command line.
package main

import (
	"flag"
	"log"
	"math"
	"time"

	"periph.io/x/periph/host"

	"github.com/bodokaiser/houston/cmd"
	"github.com/bodokaiser/houston/driver/dds"
	"github.com/bodokaiser/houston/driver/dds/ad9910"
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
			DDSSweep: &model.DDSSweep{
				Limits:   [2]float64{},
				NoDwells: [2]bool{},
			},
		},
	}

	flag.UintVar(&d.ID, "select", 0,
		"Address of DDS chip")
	flag.Float64Var(&d.Amplitude.Value, "amplitude", 1.0,
		"Constant amplitude in relative units")
	flag.Float64Var(&d.Frequency.Value, "frequency", 10e6,
		"Constant frequency in Hertz")
	flag.Float64Var(&d.PhaseOffset.Limits[0], "start", 0.0,
		"Start value of phase offset")
	flag.Float64Var(&d.PhaseOffset.Limits[1], "stop", 2*math.Pi,
		"Stop value of phase offset")
	flag.BoolVar(&d.PhaseOffset.NoDwells[0], "no-dwell-low", true,
		"ramp does not remain at lower end")
	flag.BoolVar(&d.PhaseOffset.NoDwells[1], "no-dwell-high", true,
		"ramp does not remain at upper end")
	flag.DurationVar(&d.PhaseOffset.Duration, "duration", time.Second,
		"Ramp Duration in Seconds")
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

	dev := ad9910.NewAD9910(c.DDS)
	err = dev.Init()
	if err != nil {
		log.Fatal(err)
	}

	err = sel.Select(uint8(d.ID))
	if err != nil {
		log.Fatal(err)
	}

	dev.Sweep(dds.SweepConfig{
		Limits:   d.PhaseOffset.Limits,
		NoDwells: d.PhaseOffset.NoDwells,
		Duration: d.PhaseOffset.Duration,
		Param:    dds.ParamPhase,
	})
	dev.SetAmplitude(d.Amplitude.Value)
	dev.SetFrequency(d.Frequency.Value)

	err = dev.Exec()
	if err != nil {
		log.Fatal(err)
	}
}
