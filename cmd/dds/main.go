package main

import (
	"periph.io/x/periph/host"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/bodokaiser/houston/cmd"
	"github.com/bodokaiser/houston/driver/dds"
	"github.com/bodokaiser/houston/driver/dds/ad9910"
	"github.com/bodokaiser/houston/driver/mux"
)

var config = cmd.Config{}

var (
	id uint8

	param       string
	amplitude   float64
	frequency   float64
	phaseOffset float64

	sweep    dds.SweepConfig
	playback dds.PlaybackConfig
)

func main() {
	kingpin.Flag("select", "chip select").Required().Uint8Var(&id)
	kingpin.Flag("config", "device config").Default("config.yaml").ExistingFileVar(&config.Filename)
	kingpin.Flag("debug", "verbose logging").Default("false").BoolVar(&config.DDS.Debug)
	kingpin.Command("reset", "resets a dds")

	c := kingpin.Command("const", "constant parameters")
	c.Flag("amplitude", "").Default("1").Float64Var(&amplitude)
	c.Flag("frequency", "").Required().Float64Var(&frequency)
	c.Flag("phase", "").Default("0").Float64Var(&phaseOffset)

	s := kingpin.Command("sweep", "sweeps single parameter")
	s.Flag("amplitude", "").Default("1").Float64Var(&amplitude)
	s.Flag("frequency", "").Required().Float64Var(&frequency)
	s.Flag("phase", "").Default("0").Float64Var(&phaseOffset)
	s.Flag("start", "").Required().Float64Var(&sweep.Limits[0])
	s.Flag("stop", "").Required().Float64Var(&sweep.Limits[1])
	s.Flag("nodwell-low", "").Default("true").BoolVar(&sweep.NoDwells[0])
	s.Flag("nodwell-high", "").Default("true").BoolVar(&sweep.NoDwells[1])
	s.Flag("param", "").Required().EnumVar(&param, "amplitude", "frequency", "phase")

	p := kingpin.Command("playback", "playbacks single parameter")
	s.Flag("amplitude", "").Default("1").Float64Var(&amplitude)
	s.Flag("frequency", "").Required().Float64Var(&frequency)
	s.Flag("phase", "").Default("0").Float64Var(&phaseOffset)
	p.Flag("interval", "").Required().DurationVar(&playback.Interval)
	p.Flag("trigger", "").Default("0").BoolVar(&playback.Trigger)
	p.Flag("duplex", "").Default("0").BoolVar(&playback.Duplex)
	p.Flag("data", "").Required().Float64ListVar(&playback.Data)
	p.Flag("param", "").Required().EnumVar(&param, "amplitude", "frequency", "phase")

	var m mux.Mux
	var d dds.DDS

	if _, err := host.Init(); err != nil {
		kingpin.FatalIfError(err, "host initialization")
	}

	switch kingpin.Parse() {
	default:
		config.Mux.Debug = config.DDS.Debug

		kingpin.FatalIfError(config.ReadFromFile(), "config")

		m = mux.NewDigital(config.Mux)
		kingpin.FatalIfError(m.Init(), "mux initialization")
		kingpin.FatalIfError(m.Select(id), "mux chip select")

		d = ad9910.NewAD9910(config.DDS)
		kingpin.FatalIfError(d.Init(), "dds initialization")

		switch param {
		case "amplitude":
			sweep.Param = dds.ParamAmplitude
			playback.Param = dds.ParamAmplitude
		case "frequency":
			sweep.Param = dds.ParamFrequency
			playback.Param = dds.ParamFrequency
		case "phase":
			sweep.Param = dds.ParamPhase
			playback.Param = dds.ParamPhase
		}

		fallthrough
	case "reset":
		d.Reset()
	case "const":
		d.SetAmplitude(amplitude)
		d.SetFrequency(frequency)
		d.SetPhaseOffset(phaseOffset)

		kingpin.FatalIfError(d.Exec(), "failed to update device")
	case "sweep":
		d.SetAmplitude(amplitude)
		d.SetFrequency(frequency)
		d.SetPhaseOffset(phaseOffset)
		d.SetSweep(sweep)

		kingpin.FatalIfError(d.Exec(), "failed to update device")
	case "playback":
		d.SetAmplitude(amplitude)
		d.SetFrequency(frequency)
		d.SetPhaseOffset(phaseOffset)
		d.SetPlayback(playback)

		kingpin.FatalIfError(d.Exec(), "failed to update device")
	}
}
