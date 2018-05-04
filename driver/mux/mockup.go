package mux

import "log"

// Mockup implements Mux interface.
type Mockup struct {
	Debug    bool
	Selected uint8
}

// Init implements Mux interface.
func (d *Mockup) Init() error {
	if d.Debug {
		log.Println("init")
	}
	return nil
}

// Select implements Mux interface.
func (d *Mockup) Select(n uint8) error {
	d.Selected = n
	if d.Debug {
		log.Printf("selected %v\n", n)
	}

	return nil
}
