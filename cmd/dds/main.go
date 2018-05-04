package main

import (
	"periph.io/x/periph/host"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/bodokaiser/houston/config"
	"github.com/bodokaiser/houston/driver/dds"
	"github.com/bodokaiser/houston/driver/dds/ad9910"
	"github.com/bodokaiser/houston/driver/mux"
)

type options struct {
	config.Config

	ID          uint8
	Amplitude   float64
	Frequency   float64
	PhaseOffset float64
	Param       string

	SweepConfig    dds.SweepConfig
	PlaybackConfig dds.PlaybackConfig

	Filename string
}

func main() {
	o := options{}

	kingpin.Flag("select", "chip select").Required().Uint8Var(&o.ID)
	kingpin.Flag("config", "device config").Default("config.yaml").ExistingFileVar(&o.Filename)
	kingpin.Flag("debug", "verbose logging").Default("false").BoolVar(&o.Debug)

	kingpin.Command("reset", "resets a dds")

	c := kingpin.Command("const", "constant parameters")
	c.Flag("amplitude", "").Default("1").Float64Var(&o.Amplitude)
	c.Flag("frequency", "").Required().Float64Var(&o.Frequency)
	c.Flag("phase", "").Default("0").Float64Var(&o.PhaseOffset)

	s := kingpin.Command("sweep", "sweeps single parameter")
	s.Flag("amplitude", "").Default("1").Float64Var(&o.Amplitude)
	s.Flag("frequency", "").Required().Float64Var(&o.Frequency)
	s.Flag("phase", "").Default("0").Float64Var(&o.PhaseOffset)
	s.Flag("start", "").Required().Float64Var(&o.SweepConfig.Limits[0])
	s.Flag("stop", "").Required().Float64Var(&o.SweepConfig.Limits[1])
	s.Flag("nodwell-low", "").Default("true").BoolVar(&o.SweepConfig.NoDwells[0])
	s.Flag("nodwell-high", "").Default("true").BoolVar(&o.SweepConfig.NoDwells[1])
	s.Flag("param", "").Required().EnumVar(&o.Param, "amplitude", "frequency", "phase")

	p := kingpin.Command("playback", "playbacks single parameter")
	s.Flag("amplitude", "").Default("1").Float64Var(&o.Amplitude)
	s.Flag("frequency", "").Required().Float64Var(&o.Frequency)
	s.Flag("phase", "").Default("0").Float64Var(&o.PhaseOffset)
	p.Flag("interval", "").Required().DurationVar(&o.PlaybackConfig.Interval)
	p.Flag("trigger", "").Default("0").BoolVar(&o.PlaybackConfig.Trigger)
	p.Flag("duplex", "").Default("0").BoolVar(&o.PlaybackConfig.Duplex)
	p.Flag("data", "").Required().Float64ListVar(&o.PlaybackConfig.Data)
	p.Flag("param", "").Required().EnumVar(&o.Param, "amplitude", "frequency", "phase")

	var m mux.Mux
	var d dds.DDS

	if _, err := host.Init(); err != nil {
		kingpin.FatalIfError(err, "host initialization")
	}

	switch kingpin.Parse() {
	default:
		o.Ensure()

		kingpin.FatalIfError(o.ReadFromFile(o.Filename), "config")

		m = mux.NewDigital(o.Mux)
		kingpin.FatalIfError(m.Init(), "mux initialization")
		kingpin.FatalIfError(m.Select(o.ID), "mux chip select")

		d = ad9910.NewAD9910(o.DDS)
		kingpin.FatalIfError(d.Init(), "dds initialization")

		switch o.Param {
		case "amplitude":
			o.SweepConfig.Param = dds.ParamAmplitude
			o.PlaybackConfig.Param = dds.ParamAmplitude
		case "frequency":
			o.SweepConfig.Param = dds.ParamFrequency
			o.PlaybackConfig.Param = dds.ParamFrequency
		case "phase":
			o.SweepConfig.Param = dds.ParamPhase
			o.PlaybackConfig.Param = dds.ParamPhase
		}

		fallthrough
	case "reset":
		d.Reset()
	case "const":
		d.SetAmplitude(o.Amplitude)
		d.SetFrequency(o.Frequency)
		d.SetPhaseOffset(o.PhaseOffset)

		kingpin.FatalIfError(d.Exec(), "failed to update device")
	case "sweep":
		d.SetAmplitude(o.Amplitude)
		d.SetFrequency(o.Frequency)
		d.SetPhaseOffset(o.PhaseOffset)
		d.SetSweep(o.SweepConfig)

		kingpin.FatalIfError(d.Exec(), "failed to update device")
	case "playback":
		d.SetAmplitude(o.Amplitude)
		d.SetFrequency(o.Frequency)
		d.SetPhaseOffset(o.PhaseOffset)
		d.SetPlayback(o.PlaybackConfig)

		kingpin.FatalIfError(d.Exec(), "failed to update device")
	}
}
