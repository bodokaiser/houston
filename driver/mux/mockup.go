package mux

import "log"

type Mockup struct {
	Debug    bool
	Selected uint8
}

func (d *Mockup) Init() error {
	if d.Debug {
		log.Println("init")
	}
	return nil
}

func (d *Mockup) Select(n uint8) error {
	d.Selected = n
	if d.Debug {
		log.Printf("selected %v\n", n)
	}

	return nil
}
