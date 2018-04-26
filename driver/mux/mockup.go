package mux

type Mockup struct {
	Selected uint8
}

func NewMockup() *Mockup {
	return new(Mockup)
}

func (d *Mockup) Init() error {
	return nil
}

func (d *Mockup) Select(n uint8) error {
	d.Selected = n

	return nil
}
