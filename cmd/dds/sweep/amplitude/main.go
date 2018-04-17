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

	err = dev.DigitalRamp(dds.DigitalRampConfig{
		SingleToneConfig: dds.SingleToneConfig{
			Frequency:   d.Frequency.Value,
			PhaseOffset: d.PhaseOffset.Value,
		},
		Limits:      d.Amplitude.Limits,
		NoDwell:     d.Amplitude.NoDwell,
		Duration:    d.Amplitude.Duration,
		Destination: dds.Amplitude,
	})
	if err != nil {
		log.Fatal(err)
	}
}
