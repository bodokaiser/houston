package main

import (
	"log"

	"gobot.io/x/gobot/platforms/beaglebone"
)

func main() {
	beagle := beaglebone.NewAdaptor()

	c, err := beagle.GetSpiConnection(2, 5e6, 0)
	if err != nil {
		log.Fatal(err)
	}

	w := []byte{0x00}
	r := []byte{0x00}

	err = c.Tx(w, r)
	if err != nil {
		log.Fatal(err)
	}
}
