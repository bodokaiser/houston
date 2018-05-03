package dds

import "log"

type Mockup struct {
	Debug     bool
	HadReset  bool
	HadExec   bool
	amplitude float64
	frequency float64
	phase     float64
	sweep     SweepConfig
	playback  PlaybackConfig
}

func (d *Mockup) Init() error {
	if d.Debug {
		log.Println("init")
	}
	return nil
}

func (d *Mockup) Reset() error {
	d.HadReset = true

	if d.Debug {
		log.Println("reset")
	}

	return nil
}

func (d *Mockup) Amplitude() float64 {
	return d.amplitude
}

func (d *Mockup) SetAmplitude(x float64) {
	d.amplitude = x

	if d.Debug {
		log.Printf("amplitude set to %v\n", x)
	}
}

func (d *Mockup) Frequency() float64 {
	return d.frequency
}

func (d *Mockup) SetFrequency(x float64) {
	d.frequency = x

	if d.Debug {
		log.Printf("frequency set to %v\n", x)
	}
}

func (d *Mockup) PhaseOffset() float64 {
	return d.phase
}

func (d *Mockup) SetPhaseOffset(x float64) {
	d.phase = x

	if d.Debug {
		log.Printf("phase offset set to %v\n", x)
	}
}

func (d *Mockup) Sweep() SweepConfig {
	return d.sweep
}

func (d *Mockup) SetSweep(c SweepConfig) {
	d.sweep = c

	if d.Debug {
		log.Printf("sweep set to %+v\n", c)
	}
}

func (d *Mockup) Playback() PlaybackConfig {
	return d.playback
}

func (d *Mockup) SetPlayback(c PlaybackConfig) {
	d.playback = c

	if d.Debug {
		log.Printf("playback set to %+v\n", c)
	}
}

func (d *Mockup) Exec() error {
	d.HadExec = true
	if d.Debug {
		log.Println("exec")
	}

	return nil
}
