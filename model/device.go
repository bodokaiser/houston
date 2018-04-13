package model

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// DDSDevice is a public entity to a digital synthesizer device.
//
// Instead of exposing an entity for every support operation mode we will
// decide what mode to use from the defined properties. For example in
// singletone mode it only makes sense to provide an amplitude and frequency
// property wherein in sweep mode we expect to have a frequency range defined
// over a single frequency.
type DDSDevice struct {
	Name      string  `json:"name"             validate:"required"`
	Address   uint8   `json:"-"                validate:"gte=0,lte=31"`
	Amplitude float64 `json:"amplitude"        validate:"gte=0,lte=1"`
	Frequency float64 `json:"frequency"        validate:"gt=0,lte=5e8"`
	Phase     float64 `json:"phase"            validate:"gte=0"`
}

// Validate validates DDSDevice.
func (d DDSDevice) Validate() error {
	return validate.Struct(d)
}

// DDSDevices is a collection of DDSDevices.
//
// DDSDevices implements the flag.Value interface so that you can parse a
// list of device addresses.
type DDSDevices []DDSDevice

// FindByName returns the first index of the DDSDevice with the given name.
//
// If no DDSDevice with given name is found -1 is returned.
func (s *DDSDevices) FindByName(name string) int {
	for i, d := range *s {
		if d.Name == name {
			return i
		}
	}

	return -1
}

// String implements flag.Value.
func (s *DDSDevices) String() string {
	p := make([]string, len(*s))

	for i, d := range *s {
		p[i] = strconv.Itoa(int(d.Address))
	}

	return strings.Join(p, ",")
}

// Set implements flag.Value.
func (s *DDSDevices) Set(v string) error {
	p := strings.Split(strings.TrimSpace(v), ",")

	if len(p) < 1 {
		return errors.New("no device ids provided")
	}

	for _, u := range p {
		id, err := strconv.Atoi(u)
		if err != nil {
			return err
		}

		*s = append(*s, DDSDevice{
			Name:    fmt.Sprintf("DDS%d", id),
			Address: uint8(id),
		})
	}

	return nil
}
