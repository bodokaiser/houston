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

var device = model.DDSDevice{
	Amplitude: model.DDSParam{
		DDSConst:    &model.DDSConst{},
		DDSSweep:    &model.DDSSweep{},
		DDSPlayback: &model.DDSPlayback{},
	},
	Frequency: model.DDSParam{
		DDSConst:    &model.DDSConst{},
		DDSSweep:    &model.DDSSweep{},
		DDSPlayback: &model.DDSPlayback{},
	},
	PhaseOffset: model.DDSParam{
		DDSConst:    &model.DDSConst{},
		DDSSweep:    &model.DDSSweep{},
		DDSPlayback: &model.DDSPlayback{},
	},
}

func main() {
	kingpin.Flag("select", "chip select").Required().UintVar(&device.ID)
	kingpin.Flag("config", "device config").Default("config.yaml").ExistingFileVar(&config.Filename)
	kingpin.Command("reset", "resets a dds")

	c := kingpin.Command("const", "outputs constant signal")
	c.Flag("amplitude", "").Default("1").Float64Var(&device.Amplitude.Value)
	c.Flag("frequency", "").Required().Float64Var(&device.Frequency.Value)
	c.Flag("phase", "").Default("0").Float64Var(&device.PhaseOffset.Value)

	sa := kingpin.Command("asweep", "sweeps amplitude")
	sa.Flag("frequency", "").Required().Float64Var(&device.Frequency.Value)
	sa.Flag("phase", "").Default("0").Float64Var(&device.PhaseOffset.Value)
	sa.Flag("duration", "").Required().DurationVar(&device.Amplitude.Duration)
	sa.Flag("start", "").Required().Float64Var(&device.Amplitude.Limits[0])
	sa.Flag("stop", "").Required().Float64Var(&device.Amplitude.Limits[1])
	sa.Flag("nodwell-low", "").Default("true").BoolVar(&device.Amplitude.NoDwells[0])
	sa.Flag("nodwell-high", "").Default("true").BoolVar(&device.Amplitude.NoDwells[1])

	sf := kingpin.Command("fsweep", "sweeps frequency")
	sf.Flag("amplitude", "").Default("1").Float64Var(&device.Amplitude.Value)
	sf.Flag("phase", "").Default("0").Float64Var(&device.PhaseOffset.Value)
	sf.Flag("duration", "").Required().DurationVar(&device.Frequency.Duration)
	sf.Flag("start", "").Required().Float64Var(&device.Frequency.Limits[0])
	sf.Flag("stop", "").Required().Float64Var(&device.Frequency.Limits[1])
	sf.Flag("nodwell-low", "").Default("true").BoolVar(&device.Frequency.NoDwells[0])
	sf.Flag("nodwell-high", "").Default("true").BoolVar(&device.Frequency.NoDwells[1])

	sp := kingpin.Command("psweep", "sweeps phase")
	sp.Flag("amplitude", "").Default("1").Float64Var(&device.Amplitude.Value)
	sp.Flag("frequency", "").Required().Float64Var(&device.Frequency.Value)
	sp.Flag("duration", "").Required().DurationVar(&device.PhaseOffset.Duration)
	sp.Flag("start", "").Required().Float64Var(&device.PhaseOffset.Limits[0])
	sp.Flag("stop", "").Required().Float64Var(&device.PhaseOffset.Limits[1])
	sp.Flag("nodwell-low", "").Default("true").BoolVar(&device.PhaseOffset.NoDwells[0])
	sp.Flag("nodwell-high", "").Default("true").BoolVar(&device.PhaseOffset.NoDwells[1])

	pa := kingpin.Command("aplayback", "playbacks amplitude")
	pa.Flag("frequency", "").Required().Float64Var(&device.Frequency.Value)
	pa.Flag("phase", "").Default("0").Float64Var(&device.PhaseOffset.Value)
	pa.Flag("interval", "").Required().DurationVar(&device.Amplitude.Interval)
	pa.Flag("trigger", "").Default("0").BoolVar(&device.Amplitude.Trigger)
	pa.Flag("duplex", "").Default("0").BoolVar(&device.Amplitude.Duplex)
	pa.Flag("data", "").Required().Float64ListVar(&device.Amplitude.Data)

	pf := kingpin.Command("fplayback", "playbacks frequency")
	pf.Flag("amplitude", "").Default("1").Float64Var(&device.Amplitude.Value)
	pf.Flag("phase", "").Default("0").Float64Var(&device.PhaseOffset.Value)
	pf.Flag("interval", "").Required().DurationVar(&device.Frequency.Interval)
	pf.Flag("trigger", "").Default("0").BoolVar(&device.Frequency.Trigger)
	pf.Flag("duplex", "").Default("0").BoolVar(&device.Frequency.Duplex)
	pf.Flag("data", "").Required().Float64ListVar(&device.Frequency.Data)

	pp := kingpin.Command("pplayback", "playbacks frequency")
	pp.Flag("amplitude", "").Default("1").Float64Var(&device.Amplitude.Value)
	pp.Flag("frequency", "").Default("0").Float64Var(&device.Frequency.Value)
	pp.Flag("interval", "").Required().DurationVar(&device.PhaseOffset.Interval)
	pp.Flag("trigger", "").Default("0").BoolVar(&device.PhaseOffset.Trigger)
	pp.Flag("duplex", "").Default("0").BoolVar(&device.PhaseOffset.Duplex)
	pp.Flag("data", "").Required().Float64ListVar(&device.PhaseOffset.Data)

	subcmd := kingpin.Parse()

	kingpin.FatalIfError(config.ReadFromFile(), "config")

	if _, err := host.Init(); err != nil {
		kingpin.FatalIfError(err, "host initialization")
	}

	m := mux.NewDigital(config.Mux)
	kingpin.FatalIfError(m.Init(), "mux initialization")
	kingpin.FatalIfError(m.Select(uint8(device.ID)), "mux chip select")

	d := ad9910.NewAD9910(config.DDS)
	kingpin.FatalIfError(d.Init(), "dds initialization")

	switch subcmd {
	case "reset":
		d.Reset()
	case "const":
		d.SetAmplitude(device.Amplitude.Value)
		d.SetFrequency(device.Frequency.Value)
		d.SetPhaseOffset(device.PhaseOffset.Value)
	case "asweep":
		d.Sweep(dds.SweepConfig{
			Limits:   device.Amplitude.Limits,
			NoDwells: device.Amplitude.NoDwells,
			Duration: device.Amplitude.Duration,
			Param:    dds.ParamAmplitude,
		})
		d.SetFrequency(device.Frequency.Value)
		d.SetPhaseOffset(device.PhaseOffset.Value)
	case "fsweep":
		d.Sweep(dds.SweepConfig{
			Limits:   device.Frequency.Limits,
			NoDwells: device.Frequency.NoDwells,
			Duration: device.Frequency.Duration,
			Param:    dds.ParamFrequency,
		})
		d.SetAmplitude(device.Amplitude.Value)
		d.SetPhaseOffset(device.PhaseOffset.Value)
	case "psweep":
		d.Sweep(dds.SweepConfig{
			Limits:   device.PhaseOffset.Limits,
			NoDwells: device.PhaseOffset.NoDwells,
			Duration: device.PhaseOffset.Duration,
			Param:    dds.ParamPhase,
		})
		d.SetAmplitude(device.Amplitude.Value)
		d.SetFrequency(device.Frequency.Value)
	case "aplayback":
		d.Playback(dds.PlaybackConfig{
			Trigger:  device.Amplitude.Trigger,
			Duplex:   device.Amplitude.Duplex,
			Duration: device.Amplitude.Interval,
			Data:     device.Amplitude.Data,
			Param:    dds.ParamAmplitude,
		})
		d.SetFrequency(device.Frequency.Value)
		d.SetPhaseOffset(device.PhaseOffset.Value)
	case "fplayback":
		d.Playback(dds.PlaybackConfig{
			Trigger:  device.Frequency.Trigger,
			Duplex:   device.Frequency.Duplex,
			Duration: device.Frequency.Interval,
			Data:     device.Frequency.Data,
			Param:    dds.ParamFrequency,
		})
		d.SetAmplitude(device.Amplitude.Value)
		d.SetPhaseOffset(device.PhaseOffset.Value)
	case "pplayback":
		d.Playback(dds.PlaybackConfig{
			Trigger:  device.PhaseOffset.Trigger,
			Duplex:   device.PhaseOffset.Duplex,
			Duration: device.PhaseOffset.Interval,
			Data:     device.PhaseOffset.Data,
			Param:    dds.ParamPhase,
		})
		d.SetAmplitude(device.Amplitude.Value)
		d.SetFrequency(device.Frequency.Value)
	}

	if subcmd != "reset" {
		kingpin.FatalIfError(d.Exec(), "failed to write config")
	}
}
