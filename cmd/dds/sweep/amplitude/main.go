// cli controls a DDS array via command line.
package main

import (
	"flag"
	"log"
	"time"

	"periph.io/x/periph/host"

	"github.com/bodokaiser/houston/config"
	"github.com/bodokaiser/houston/driver/dds"
	"github.com/bodokaiser/houston/driver/dds/ad99xx"
	"github.com/bodokaiser/houston/driver/mux"
	"github.com/bodokaiser/houston/model"
)

func main() {
	c := config.Config{}
	d := model.DDSDevice{
		Amplitude: model.DDSParam{
			DDSSweep: &model.DDSSweep{
				Limits:  [2]float64{},
				NoDwell: [2]bool{true, true},
			},
		},
		Frequency: model.DDSParam{
			DDSConst: &model.DDSConst{},
		},
		PhaseOffset: model.DDSParam{
			DDSConst: &model.DDSConst{},
		},
	}

	flag.Var(&c, "config", "path to config file")
	flag.UintVar(&d.ID, "select", 0,
		"Address of DDS chip")
	flag.Float64Var(&d.Frequency.Value, "frequency", 10e6,
		"Constant frequency in Hertz")
	flag.Float64Var(&d.Amplitude.Limits[0], "start", 0.5,
		"Start value for amplitude in relative units")
	flag.Float64Var(&d.Amplitude.Limits[1], "stop", 1.0,
		"Stop value for amplitude in relative units")
	flag.DurationVar(&d.Amplitude.Duration, "duration", time.Second,
		"Ramp Duration in Seconds")
	flag.Float64Var(&d.PhaseOffset.Value, "phase", 0.0,
		"Phase offset in Radiants")
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

	dev.SetFrequency(d.Frequency.Value)
	dev.SetPhaseOffset(d.PhaseOffset.Value)
	err = dev.SweepAmplitude(dds.DigitalRampConfig{
		SingleToneConfig: dds.SingleToneConfig{
			Frequency:   d.Frequency.Value,
			PhaseOffset: d.PhaseOffset.Value,
		},
		Limits:   d.Amplitude.Limits,
		NoDwell:  d.Amplitude.NoDwell,
		Duration: d.Amplitude.Duration,
	})
	if err != nil {
		log.Fatal(err)
	}
}
