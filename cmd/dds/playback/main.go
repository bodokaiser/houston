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

func main() {
	d := &model.DDSDevice{
		Amplitude: model.DDSParam{
			DDSPlayback: &model.DDSPlayback{
				WithTrigger: false,
				WithDuplex:  false,
				Interval:    time.Second,
				Data:        []float64{0.5, 0, 1},
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
	flag.Float64Var(&d.PhaseOffset.Value, "phase", 0.0,
		"Phase offset in Radiants")
	flag.DurationVar(&d.Amplitude.Interval, "duration", time.Second,
		"Sample Interval of playback")
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

	err = dev.Playback(dds.PlaybackConfig{
		SingleToneConfig: dds.SingleToneConfig{
			Frequency:   d.Frequency.Value,
			PhaseOffset: d.PhaseOffset.Value,
		},
		Duration:    d.Amplitude.Interval,
		Data:        d.Amplitude.Data,
		Destination: dds.Amplitude,
	})
	if err != nil {
		log.Fatal(err)
	}
}
