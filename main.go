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
	volume      string
	battery     string
	host        string
	timestamp   string
	temperature string
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

	volume := pulseaudio.Pulseaudio{}
	volumeDrift := make(chan bool)
	go volume.Update(volumeDrift)

	battery := battery.Battery{}
	batteryDrift := make(chan bool)
	go battery.Update(batteryDrift, "battery_BAT0")

	timestamp := timestamp.Timestamp{}
	timestampDrift := make(chan bool)
	go timestamp.Update(timestampDrift)

	host := host.Host{}
	hostDrift := make(chan bool)
	go host.Update(hostDrift)

	temperature := temperature.Temperature{}
	temperatureDrift := make(chan bool)
	go temperature.Update(temperatureDrift)

	for {
		select {
		case <-batteryDrift:
			stdout.battery = battery.Get()
			stdout.write()
		case <-volumeDrift:
			stdout.volume = volume.Get()
			stdout.write()
		case <-timestampDrift:
			stdout.timestamp = timestamp.Get()
			stdout.write()
		case <-hostDrift:
			stdout.host = host.Get()
			stdout.write()
		case <-temperatureDrift:
			stdout.temperature = temperature.Get()
			stdout.write()
		}
	}
}
