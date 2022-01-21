package gobar

import (
	"fmt"

	"github.com/n3ph/gobar/pkg/host"
	"github.com/n3ph/gobar/pkg/temperature"
	"github.com/n3ph/gobar/pkg/timestamp"
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

func Gobar() {
	var stdout Elements
	var drift bool

	host := host.New()
	hostValue := make(chan string)
	go host.Update(hostValue)

	temperature := temperature.New()
	temperatureValue := make(chan string)
	go temperature.Update(temperatureValue)

	// battery := battery.New()
	// batteryValue := make(chan bool)
	// go battery.Update(batteryValue, "battery_BAT0")

	// volume := pulseaudio.New()
	// volumeValue := make(chan bool)
	// go volume.Update(volumeValue)

	timestamp := timestamp.New()
	timestampValue := make(chan string)
	go timestamp.Update(timestampValue)

	for {
		select {
		case value := <-hostValue:
			stdout.host = value
			drift = true
		case value := <-temperatureValue:
			stdout.temperature = value
			drift = true
		// case value := <-batteryValue:
		// 	stdout.battery = value
		// 	drift = true
		// case value := <-volumeValue:
		// 	stdout.volume = value
		// 	drift = true
		case value := <-timestampValue:
			stdout.timestamp = value
			drift = true
		}

		if drift {
			stdout.write()
			drift = false
		}
	}
}
