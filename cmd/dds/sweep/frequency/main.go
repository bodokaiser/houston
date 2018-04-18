// cli controls a DDS array via command line.
package main

import (
	"flag"
	"log"
	"time"

	"periph.io/x/periph/host"

	"github.com/bodokaiser/houston/driver/dds"
	"github.com/bodokaiser/houston/driver/dds/ad99xx"
	"github.com/bodokaiser/houston/driver/mux"
	"github.com/bodokaiser/houston/model"
)

const (
	defaultSysClock = 1e9
	defaultRefClock = 1e7
)

const (
	defaultSPIDevice  = "SPI1.0"
	defaultSPIMaxFreq = 5e6
	defaultSPIMode    = 0
)

const (
	defaultResetPin    = "65"
	defaultIOUpdatePin = "27"
)

var defaultMuxPins = []string{"48", "30", "60", "31", "50"}

func main() {
	d := &model.DDSDevice{
		Frequency: model.DDSParam{
			DDSSweep: &model.DDSSweep{
				Limits:  [2]float64{},
				NoDwell: [2]bool{true, true},
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
	flag.Float64Var(&d.Frequency.Limits[0], "start", 12e5,
		"Start value for frequency in Hertz")
	flag.Float64Var(&d.Frequency.Limits[1], "stop", 16e5,
		"Stop value for frequency in Hertz")
	flag.DurationVar(&d.Frequency.Duration, "duration", time.Second,
		"Ramp Duration in Seconds")
	flag.Float64Var(&d.PhaseOffset.Value, "phase", 0.0,
		"Phase offset in Radiants")
	flag.Parse()

	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	sel, err := mux.NewDigital(defaultMuxPins)
	if err != nil {
		log.Fatal(err)
	}

	dev, err := ad99xx.NewAD9910(ad99xx.Config{
		SysClock:    defaultSysClock,
		RefClock:    defaultRefClock,
		ResetPin:    defaultResetPin,
		IOUpdatePin: defaultIOUpdatePin,
		SPIDevice:   defaultSPIDevice,
		SPIMaxFreq:  defaultSPIMaxFreq,
		SPIMode:     defaultSPIMode,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = sel.Select(uint8(d.ID))
	if err != nil {
		log.Fatal(err)
	}

	err = dev.SweepFrequency(dds.DigitalRampConfig{
		SingleToneConfig: dds.SingleToneConfig{
			Amplitude:   d.Amplitude.Value,
			PhaseOffset: d.PhaseOffset.Value,
		},
		Limits:   d.Frequency.Limits,
		NoDwell:  d.Frequency.NoDwell,
		Duration: d.Frequency.Duration,
	})
	if err != nil {
		log.Fatal(err)
	}
}
