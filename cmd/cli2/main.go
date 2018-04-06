package main

import (
	"log"
	"time"

	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/beaglebone"
)

func main() {
	beagle := beaglebone.NewAdaptor()

	led := gpio.NewLedDriver(beagle, "P8_18")
	if err := led.Start(); err != nil {
		log.Fatal(err)
	}

	led.On()
	time.Sleep(time.Millisecond)
	led.Off()
}
