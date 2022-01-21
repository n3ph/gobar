package main

import (
	"fmt"

	"main/src/battery"
	"main/src/host"
	"main/src/pulseaudio"
	"main/src/temperature"
	"main/src/timestamp"
)

type Elements struct {
	battery     string
	host        string
	temperature string
	timestamp   string
	volume      string
}

func (elements Elements) write() {
	var stdout string

	for _, element := range []string{elements.host, elements.temperature, elements.battery, elements.volume, elements.timestamp} {
		stdout += element + " | "
	}

	fmt.Println(stdout)
}

func main() {
	var stdout Elements

	host := host.New()
	hostDrift := make(chan bool)
	go host.Update(hostDrift)

	temperature := temperature.New()
	temperatureDrift := make(chan bool)
	go temperature.Update(temperatureDrift)

	battery := battery.New()
	batteryDrift := make(chan bool)
	go battery.Update(batteryDrift, "battery_BAT0")

	volume := pulseaudio.New()
	volumeDrift := make(chan bool)
	go volume.Update(volumeDrift)

	timestamp := timestamp.New()
	timestampDrift := make(chan bool)
	go timestamp.Update(timestampDrift)

	for {
		select {
		case <-hostDrift:
			stdout.host = host.Get()
			stdout.write()
		case <-temperatureDrift:
			stdout.temperature = temperature.Get()
			stdout.write()
		case <-batteryDrift:
			stdout.battery = battery.Get()
			stdout.write()
		case <-volumeDrift:
			stdout.volume = volume.Get()
			stdout.write()
		case <-timestampDrift:
			stdout.timestamp = timestamp.Get()
			stdout.write()
		}
	}
}
