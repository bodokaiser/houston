// cli controls a DDS array via command line.
package main

import (
	"log"
	"os"

	"periph.io/x/periph/host"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/bodokaiser/houston/cmd"
	"github.com/bodokaiser/houston/driver/dds"
	"github.com/bodokaiser/houston/driver/dds/ad99xx"
	"github.com/bodokaiser/houston/driver/mux"
	"github.com/bodokaiser/houston/model"
)

func main() {
	c := cmd.Config{}
	m := model.DDSDevice{
		Amplitude: model.DDSParam{
			DDSPlayback: &model.DDSPlayback{},
		},
		Frequency: model.DDSParam{
			DDSConst: &model.DDSConst{},
		},
		PhaseOffset: model.DDSParam{
			DDSConst: &model.DDSConst{},
		},
	}

	a := kingpin.New("playback", "playback of arbitrary waveforms")
	a.Flag("select", "chip select").
		Required().UintVar(&m.ID)
	a.Flag("frequency", "frequency in Hz").
		Required().Float64Var(&m.Frequency.Value)
	a.Flag("phase", "phase offset in rad").
		Default("0").Float64Var(&m.PhaseOffset.Value)
	a.Flag("duration", "playback duration").
		Required().DurationVar(&m.Amplitude.DDSPlayback.Duration)
	a.Flag("trigger", "playback on trigger").
		Default("0").BoolVar(&m.Amplitude.Trigger)
	a.Flag("duplex", "playback bidirectional").
		Default("0").BoolVar(&m.Amplitude.Duplex)
	a.Flag("data", "playback data").
		Required().Float64ListVar(&m.Amplitude.Data)
	kingpin.MustParse(a.Parse(os.Args[1:]))

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

	dev.Playback(dds.PlaybackConfig{
		Trigger:  m.Amplitude.Trigger,
		Duplex:   m.Amplitude.Duplex,
		Duration: m.Amplitude.DDSPlayback.Duration,
		Data:     m.Amplitude.Data,
		Param:    dds.ParamAmplitude,
	})
	dev.SetFrequency(m.Frequency.Value)
	dev.SetPhaseOffset(m.PhaseOffset.Value)

	err = dev.Exec()
	if err != nil {
		log.Fatal(err)
	}
}
