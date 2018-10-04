// Command line interface to DDS devices.
package main

import (
	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/host"

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

var cmd = options{}

func main() {
	kingpin.Flag("select", "chip select").Required().Uint8Var(&cmd.ID)
	kingpin.Flag("config", "device config").ExistingFileVar(&cmd.Filename)
	kingpin.Flag("debug", "verbose logging").Default("false").BoolVar(&cmd.Debug)

	kingpin.Command("reset", "resets a dds")
	kingpin.Command("trigger", "triggers a dds")

	c := kingpin.Command("const", "constant parameters")
	c.Flag("amplitude", "").Default("1").Float64Var(&cmd.Amplitude)
	c.Flag("phase", "").Default("0").Float64Var(&cmd.PhaseOffset)
	c.Flag("frequency", "").Required().Float64Var(&cmd.Frequency)

	s := kingpin.Command("sweep", "sweeps single parameter")
	s.Flag("amplitude", "").Default("1").Float64Var(&cmd.Amplitude)
	s.Flag("phase", "").Default("0").Float64Var(&cmd.PhaseOffset)
	s.Flag("frequency", "").Default("10e6").Float64Var(&cmd.Frequency)
	s.Flag("start", "").Required().Float64Var(&cmd.SweepConfig.Limits[0])
	s.Flag("stop", "").Required().Float64Var(&cmd.SweepConfig.Limits[1])
	s.Flag("nodwell-low", "").Default("true").BoolVar(&cmd.SweepConfig.NoDwells[0])
	s.Flag("nodwell-high", "").Default("true").BoolVar(&cmd.SweepConfig.NoDwells[1])
	s.Flag("duration", "").Required().DurationVar(&cmd.SweepConfig.Duration)
	s.Flag("param", "").Required().EnumVar(&cmd.Param, "amplitude", "frequency", "phase")

	p := kingpin.Command("playback", "playbacks single parameter")
	p.Flag("amplitude", "").Default("1").Float64Var(&cmd.Amplitude)
	p.Flag("phase", "").Default("0").Float64Var(&cmd.PhaseOffset)
	p.Flag("frequency", "").Default("10e6").Float64Var(&cmd.Frequency)
	p.Flag("interval", "").Required().DurationVar(&cmd.PlaybackConfig.Interval)
	p.Flag("trigger", "").Default("0").BoolVar(&cmd.PlaybackConfig.Trigger)
	p.Flag("duplex", "").Default("0").BoolVar(&cmd.PlaybackConfig.Duplex)
	p.Flag("data", "").Required().Float64ListVar(&cmd.PlaybackConfig.Data)
	p.Flag("param", "").Required().EnumVar(&cmd.Param, "amplitude", "frequency", "phase")

	_, err := host.Init()
	if err != nil {
		kingpin.FatalIfError(err, "host initialization")
	}

	subcmd := kingpin.Parse()

	cmd.Ensure()
	cmd.ReadFromBox(cmd.Filename)

	m := mux.NewDigital(cmd.Mux)
	kingpin.FatalIfError(m.Init(), "mux initialization")
	kingpin.FatalIfError(m.Select(cmd.ID), "mux chip select")

	cmd.DDS.Config.SPIPort, _ = spireg.Open(cmd.DDS.SPI.Device)
	cmd.DDS.Config.ResetPin = gpioreg.ByName(cmd.DDS.GPIO.Reset)
	cmd.DDS.Config.UpdatePin = gpioreg.ByName(cmd.DDS.GPIO.Update)

	d := ad9910.NewAD9910(cmd.DDS.Config)
	kingpin.FatalIfError(d.Init(), "dds initialization")

	switch cmd.Param {
	case "amplitude":
		cmd.SweepConfig.Param = dds.ParamAmplitude
		cmd.PlaybackConfig.Param = dds.ParamAmplitude
	case "frequency":
		cmd.SweepConfig.Param = dds.ParamFrequency
		cmd.PlaybackConfig.Param = dds.ParamFrequency
	case "phase":
		cmd.SweepConfig.Param = dds.ParamPhase
		cmd.PlaybackConfig.Param = dds.ParamPhase
	}

	switch subcmd {
	case "reset":
		d.Reset()
	case "update":
		d.Update()
	case "const":
		d.SetAmplitude(cmd.Amplitude)
		d.SetFrequency(cmd.Frequency)
		d.SetPhaseOffset(cmd.PhaseOffset)

		kingpin.FatalIfError(d.Exec(), "failed to update device")
	case "sweep":
		d.SetAmplitude(cmd.Amplitude)
		d.SetFrequency(cmd.Frequency)
		d.SetPhaseOffset(cmd.PhaseOffset)
		d.SetSweep(cmd.SweepConfig)

		kingpin.FatalIfError(d.Exec(), "failed to update device")
	case "playback":
		d.SetAmplitude(cmd.Amplitude)
		d.SetFrequency(cmd.Frequency)
		d.SetPhaseOffset(cmd.PhaseOffset)
		d.SetPlayback(cmd.PlaybackConfig)

		kingpin.FatalIfError(d.Exec(), "failed to update device")
	}
}
