// Package dds provides device drivers for direct digital synthesizer.
package dds

type Synthesizer interface {
	Amplitude() float64
	Frequency() float64
	PhaseOffset() float64

	SetAmplitude(float64)
	SetFrequency(float64)
	SetPhaseOffset(float64)

	Sync() error
}

type SweepSynthesizer interface {
	Synthesizer
}

type ArbitrarySynthesizer interface {
	Synthesizer
}
