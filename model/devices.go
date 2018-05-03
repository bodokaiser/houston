package model

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// DDSDevices is a collection of DDSDevices.
//
// DDSDevices implements the flag.Value interface so that you can parse a
// list of device addresses.
type DDSDevices []DDSDevice

// FindByID returns the first index of the DDSDevice with the given id
// or -1 if not match was found.
func (s *DDSDevices) FindByID(id uint8) int {
	for i, d := range *s {
		if d.ID == id {
			return i
		}
	}

	return -1
}

func (s *DDSDevices) FindByIDString(id string) int {
	i, err := strconv.Atoi(id)
	if err != nil {
		return -1
	}

	return s.FindByID(uint8(i))
}

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
		p[i] = strconv.Itoa(int(d.ID))
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
			ID:   uint8(id),
			Name: fmt.Sprintf("DDS%d", id),
		})
	}

	return nil
}
