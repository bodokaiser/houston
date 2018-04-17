// cli controls a DDS array via command line.
package main

import (
	"flag"
	"log"

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
				Limits:    [2]float64{0.8, 1.0},
				StepSize:  [2]float64{1, 1},
				SlopeRate: [2]float64{1e5, 1e6},
				NoDwell:   [2]bool{true, true},
			},
		},
		Frequency: model.DDSParam{
			DDSConst: &model.DDSConst{Value: 100e3},
		},
		PhaseOffset: model.DDSParam{
			DDSConst: &model.DDSConst{Value: 0.0},
		},
	}

	flag.UintVar(&d.ID, "select", 0, "address to select")
	flag.Float64Var(&d.Amplitude.Limits[0], "lower-limit", 0.0,
		"Lower limit for amplitude [0, 1]")
	flag.Float64Var(&d.Amplitude.Limits[1], "upper-limit", 1.0,
		"Upper limit for amplitude [0, 1]")
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
		StepSize:    d.Amplitude.StepSize,
		SlopeRate:   d.Amplitude.SlopeRate,
		NoDwell:     d.Amplitude.NoDwell,
		Destination: dds.Amplitude,
	})
	if err != nil {
		log.Fatal(err)
	}
}
