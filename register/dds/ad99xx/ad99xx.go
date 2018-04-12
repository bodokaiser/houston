// Package ad99xx provides the device register for direct digital synthesizer
// of the AD99xx family.
//
// Up to the moment of writing we only support the AD9910 chip.
package ad99xx

import "math"

// AD99xx instruction byte flags.
//
// The instruction byte flags define whether the byte sequence we send through
// the serial bus should be interpreted as write or read operation.
// They are applied by bitwise OR on the address byte. To direct a read
// you still have to write the number of bytes register after the address byte
// though their values will not be interpreted.
const (
	FlagOpRead  = 0x80
	FlagOpWrite = 0x00
)

// FrequencyToFTW returns the 32 bit FTW given desired output frequency
// and system frequency.
func FrequencyToFTW(sys float64, out float64) uint32 {
	return uint32(math.Round(out / sys * (1 << 32)))
}

// AmplitudeToASF returns the 14 bit ASF given amplitude scale from 0 to 1.
func AmplitudeToASF(ampl float64) uint16 {
	return uint16(math.Round(ampl * (1 << 14)))
}

// PhaseToPOW returns the 16 bit POW given phase in radiants.
func PhaseToPOW(phase float64) uint16 {
	return uint16(math.Round(phase / (2 * math.Pi) * (1 << 16)))
}
