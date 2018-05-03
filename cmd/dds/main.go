package main

import (
	"periph.io/x/periph/host"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/bodokaiser/houston/cmd"
	"github.com/bodokaiser/houston/driver/dds"
	"github.com/bodokaiser/houston/driver/dds/ad9910"
	"github.com/bodokaiser/houston/driver/mux"
	"github.com/bodokaiser/houston/model"
)

var config = cmd.Config{}

var device = model.DDSDevice{}

func main() {
	kingpin.Flag("select", "chip select").Required().Uint8Var(&device.ID)
	kingpin.Flag("config", "device config").Default("config.yaml").ExistingFileVar(&config.Filename)
	kingpin.Flag("debug", "verbose logging").Default("false").BoolVar(&config.DDS.Debug)
	kingpin.Command("reset", "resets a dds")

	c := kingpin.Command("const", "outputs constant signal")
	c.Flag("amplitude", "").Default("1").Float64Var(&device.Amplitude.Const.Value)
	c.Flag("frequency", "").Required().Float64Var(&device.Frequency.Const.Value)
	c.Flag("phase", "").Default("0").Float64Var(&device.PhaseOffset.Const.Value)

	sa := kingpin.Command("asweep", "sweeps amplitude")
	sa.Flag("frequency", "").Required().Float64Var(&device.Frequency.Const.Value)
	sa.Flag("phase", "").Default("0").Float64Var(&device.PhaseOffset.Const.Value)
	sa.Flag("duration", "").Required().DurationVar(&device.Amplitude.Sweep.Duration)
	sa.Flag("start", "").Required().Float64Var(&device.Amplitude.Sweep.Limits[0])
	sa.Flag("stop", "").Required().Float64Var(&device.Amplitude.Sweep.Limits[1])
	sa.Flag("nodwell-low", "").Default("true").BoolVar(&device.Amplitude.Sweep.NoDwells[0])
	sa.Flag("nodwell-high", "").Default("true").BoolVar(&device.Amplitude.Sweep.NoDwells[1])

	sf := kingpin.Command("fsweep", "sweeps frequency")
	sf.Flag("amplitude", "").Default("1").Float64Var(&device.Amplitude.Const.Value)
	sf.Flag("phase", "").Default("0").Float64Var(&device.PhaseOffset.Const.Value)
	sf.Flag("duration", "").Required().DurationVar(&device.Frequency.Sweep.Duration)
	sf.Flag("start", "").Required().Float64Var(&device.Frequency.Sweep.Limits[0])
	sf.Flag("stop", "").Required().Float64Var(&device.Frequency.Sweep.Limits[1])
	sf.Flag("nodwell-low", "").Default("true").BoolVar(&device.Frequency.Sweep.NoDwells[0])
	sf.Flag("nodwell-high", "").Default("true").BoolVar(&device.Frequency.Sweep.NoDwells[1])

	sp := kingpin.Command("psweep", "sweeps phase")
	sp.Flag("amplitude", "").Default("1").Float64Var(&device.Amplitude.Const.Value)
	sp.Flag("frequency", "").Required().Float64Var(&device.Frequency.Const.Value)
	sp.Flag("duration", "").Required().DurationVar(&device.PhaseOffset.Sweep.Duration)
	sp.Flag("start", "").Required().Float64Var(&device.PhaseOffset.Sweep.Limits[0])
	sp.Flag("stop", "").Required().Float64Var(&device.PhaseOffset.Sweep.Limits[1])
	sp.Flag("nodwell-low", "").Default("true").BoolVar(&device.PhaseOffset.Sweep.NoDwells[0])
	sp.Flag("nodwell-high", "").Default("true").BoolVar(&device.PhaseOffset.Sweep.NoDwells[1])

	pa := kingpin.Command("aplayback", "playbacks amplitude")
	pa.Flag("frequency", "").Required().Float64Var(&device.Frequency.Const.Value)
	pa.Flag("phase", "").Default("0").Float64Var(&device.PhaseOffset.Const.Value)
	pa.Flag("interval", "").Required().DurationVar(&device.Amplitude.Playback.Interval)
	pa.Flag("trigger", "").Default("0").BoolVar(&device.Amplitude.Playback.Trigger)
	pa.Flag("duplex", "").Default("0").BoolVar(&device.Amplitude.Playback.Duplex)
	pa.Flag("data", "").Required().Float64ListVar(&device.Amplitude.Playback.Data)

	pf := kingpin.Command("fplayback", "playbacks frequency")
	pf.Flag("amplitude", "").Default("1").Float64Var(&device.Amplitude.Const.Value)
	pf.Flag("phase", "").Default("0").Float64Var(&device.PhaseOffset.Const.Value)
	pf.Flag("interval", "").Required().DurationVar(&device.Frequency.Playback.Interval)
	pf.Flag("trigger", "").Default("0").BoolVar(&device.Frequency.Playback.Trigger)
	pf.Flag("duplex", "").Default("0").BoolVar(&device.Frequency.Playback.Duplex)
	pf.Flag("data", "").Required().Float64ListVar(&device.Frequency.Playback.Data)

	pp := kingpin.Command("pplayback", "playbacks frequency")
	pp.Flag("amplitude", "").Default("1").Float64Var(&device.Amplitude.Const.Value)
	pp.Flag("frequency", "").Default("0").Float64Var(&device.Frequency.Const.Value)
	pp.Flag("interval", "").Required().DurationVar(&device.PhaseOffset.Playback.Interval)
	pp.Flag("trigger", "").Default("0").BoolVar(&device.PhaseOffset.Playback.Trigger)
	pp.Flag("duplex", "").Default("0").BoolVar(&device.PhaseOffset.Playback.Duplex)
	pp.Flag("data", "").Required().Float64ListVar(&device.PhaseOffset.Playback.Data)

	subcmd := kingpin.Parse()
	config.Mux.Debug = config.DDS.Debug

	kingpin.FatalIfError(config.ReadFromFile(), "config")

	if _, err := host.Init(); err != nil {
		kingpin.FatalIfError(err, "host initialization")
	}

	m := mux.NewDigital(config.Mux)
	kingpin.FatalIfError(m.Init(), "mux initialization")
	kingpin.FatalIfError(m.Select(device.ID), "mux chip select")

	d := ad9910.NewAD9910(config.DDS)
	kingpin.FatalIfError(d.Init(), "dds initialization")

	switch subcmd {
	case "reset":
		d.Reset()
	case "const":
		d.SetAmplitude(device.Amplitude.Const.Value)
		d.SetFrequency(device.Frequency.Const.Value)
		d.SetPhaseOffset(device.PhaseOffset.Const.Value)
	case "asweep":
		d.SetSweep(dds.SweepConfig{
			Limits:   device.Amplitude.Sweep.Limits,
			NoDwells: device.Amplitude.Sweep.NoDwells,
			Duration: device.Amplitude.Sweep.Duration,
			Param:    dds.ParamAmplitude,
		})
		d.SetFrequency(device.Frequency.Const.Value)
		d.SetPhaseOffset(device.PhaseOffset.Const.Value)
	case "fsweep":
		d.SetSweep(dds.SweepConfig{
			Limits:   device.Frequency.Sweep.Limits,
			NoDwells: device.Frequency.Sweep.NoDwells,
			Duration: device.Frequency.Sweep.Duration,
			Param:    dds.ParamFrequency,
		})
		d.SetAmplitude(device.Amplitude.Const.Value)
		d.SetPhaseOffset(device.PhaseOffset.Const.Value)
	case "psweep":
		d.SetSweep(dds.SweepConfig{
			Limits:   device.PhaseOffset.Sweep.Limits,
			NoDwells: device.PhaseOffset.Sweep.NoDwells,
			Duration: device.PhaseOffset.Sweep.Duration,
			Param:    dds.ParamPhase,
		})
		d.SetAmplitude(device.Amplitude.Const.Value)
		d.SetFrequency(device.Frequency.Const.Value)
	case "aplayback":
		d.SetPlayback(dds.PlaybackConfig{
			Trigger:  device.Amplitude.Playback.Trigger,
			Duplex:   device.Amplitude.Playback.Duplex,
			Duration: device.Amplitude.Playback.Interval,
			Data:     device.Amplitude.Playback.Data,
			Param:    dds.ParamAmplitude,
		})
		d.SetFrequency(device.Frequency.Const.Value)
		d.SetPhaseOffset(device.PhaseOffset.Const.Value)
	case "fplayback":
		d.SetPlayback(dds.PlaybackConfig{
			Trigger:  device.Frequency.Playback.Trigger,
			Duplex:   device.Frequency.Playback.Duplex,
			Duration: device.Frequency.Playback.Interval,
			Data:     device.Frequency.Playback.Data,
			Param:    dds.ParamFrequency,
		})
		d.SetAmplitude(device.Amplitude.Const.Value)
		d.SetPhaseOffset(device.PhaseOffset.Const.Value)
	case "pplayback":
		d.SetPlayback(dds.PlaybackConfig{
			Trigger:  device.PhaseOffset.Playback.Trigger,
			Duplex:   device.PhaseOffset.Playback.Duplex,
			Duration: device.PhaseOffset.Playback.Interval,
			Data:     device.PhaseOffset.Playback.Data,
			Param:    dds.ParamPhase,
		})
		d.SetAmplitude(device.Amplitude.Const.Value)
		d.SetFrequency(device.Frequency.Const.Value)
	}

	if subcmd != "reset" {
		kingpin.FatalIfError(d.Exec(), "failed to write config")
	}
}
