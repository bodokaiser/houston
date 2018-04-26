package dds

type Mockup struct {
	Debug     bool
	amplitude float64
	frequency float64
	phase     float64
	sweep     SweepConfig
	playback  PlaybackConfig
}

func NewMockup(c Config) *Mockup {
	return &Mockup{
		Debug:     c.Debug,
		amplitude: 1.0,
		frequency: 1e6,
		phase:     0,
	}
}

func (d *Mockup) Init() error {
	return nil
}

func (d *Mockup) Amplitude() float64 {
	return d.amplitude
}

func (d *Mockup) SetAmplitude(x float64) {
	d.amplitude = x
}

func (d *Mockup) Frequency() float64 {
	return d.frequency
}

func (d *Mockup) SetFrequency(x float64) {
	d.frequency = x
}

func (d *Mockup) PhaseOffset() float64 {
	return d.phase
}

func (d *Mockup) SetPhaseOffset(x float64) {
	d.phase = x
}

func (d *Mockup) Sweep() SweepConfig {
	return d.sweep
}

func (d *Mockup) SetSweep(c SweepConfig) {
	d.sweep = c
}

func (d *Mockup) Playback() PlaybackConfig {
	return d.playback
}

func (d *Mockup) SetPlayback(c PlaybackConfig) {
	d.playback = c
}
