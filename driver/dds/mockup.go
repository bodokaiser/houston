package dds

import "log"

// Mockup implements DDS interface.
//
// Mockup can be used to test the DDS interface without the need of fully
// functioning hardware present.
type Mockup struct {
	Debug     bool
	HadReset  bool
	HadUpdate  bool
	HadExec   bool
	amplitude float64
	frequency float64
	phase     float64
	sweep     SweepConfig
	playback  PlaybackConfig
}

// Init implements driver.Driver interface.
func (d *Mockup) Init() error {
	if d.Debug {
		log.Println("init")
	}
	return nil
}

// Reset implements DDS interface.
func (d *Mockup) Reset() error {
	d.HadReset = true

	if d.Debug {
		log.Println("reset")
	}

	return nil
}

// Update implements DDS interface.
func (d *Mockup) Update() error {
	d.HadUpdate = true

	if d.Debug {
		log.Println("update")
	}

	return nil
}

// Amplitude implements DDS interface.
func (d *Mockup) Amplitude() float64 {
	return d.amplitude
}

// SetAmplitude implements DDS interface.
func (d *Mockup) SetAmplitude(x float64) {
	d.amplitude = x

	if d.Debug {
		log.Printf("amplitude set to %v\n", x)
	}
}

// Frequency implements DDS interface.
func (d *Mockup) Frequency() float64 {
	return d.frequency
}

// SetFrequency implements DDS interface.
func (d *Mockup) SetFrequency(x float64) {
	d.frequency = x

	if d.Debug {
		log.Printf("frequency set to %v\n", x)
	}
}

// PhaseOffset implements DDS interface.
func (d *Mockup) PhaseOffset() float64 {
	return d.phase
}

// SetPhaseOffset implements DDS interface.
func (d *Mockup) SetPhaseOffset(x float64) {
	d.phase = x

	if d.Debug {
		log.Printf("phase offset set to %v\n", x)
	}
}

// Sweep implements DDS interface.
func (d *Mockup) Sweep() SweepConfig {
	return d.sweep
}

// SetSweep implements DDS interface.
func (d *Mockup) SetSweep(c SweepConfig) {
	d.sweep = c

	if d.Debug {
		log.Printf("sweep set to %+v\n", c)
	}
}

// Playback implements DDS interface.
func (d *Mockup) Playback() PlaybackConfig {
	return d.playback
}

// SetPlayback implements DDS interface.
func (d *Mockup) SetPlayback(c PlaybackConfig) {
	d.playback = c

	if d.Debug {
		log.Printf("playback set to %+v\n", c)
	}
}

// Exec implements DDS interface.
func (d *Mockup) Exec() error {
	d.HadExec = true
	if d.Debug {
		log.Println("exec")
	}

	return nil
}
