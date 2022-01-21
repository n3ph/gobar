package gobar

import (
	"fmt"
	"time"

	"github.com/n3ph/gobar/pkg/battery"
	"github.com/n3ph/gobar/pkg/host"
	"github.com/n3ph/gobar/pkg/pulseaudio"
	"github.com/n3ph/gobar/pkg/temperature"
)

func getTimestamp() string {
	return time.Now().Format("02.01.2006 15:04:05.0000")
}

type Elements struct {
	battery     string
	host        string
	temperature string
	timestamp   string
	volume      string
}

func (elements Elements) write() {
	var stdout string

	for _, element := range []string{elements.host, elements.temperature, elements.battery, elements.volume, getTimestamp()} {
		stdout += element + " | "
	}

	fmt.Println(stdout)
}

func Gobar() {
	var stdout Elements
	var drift bool
	quitChan := make(chan struct{})

	host := host.New()
	hostValueChan := make(chan string)
	hostErrChan := make(chan error)
	go host.Update(quitChan, hostValueChan, hostErrChan)

	temperature := temperature.New("amdgpu_edge_input")
	temperatureValueChan := make(chan string)
	temperatureErrChan := make(chan error)
	go temperature.Update(quitChan, temperatureValueChan, temperatureErrChan)

	battery := battery.New("battery_BAT0")
	batteryValueChan := make(chan string)
	batteryErrChan := make(chan error)
	go battery.Update(quitChan, batteryValueChan, batteryErrChan)

	volume := pulseaudio.New()
	volumeValueChan := make(chan string)
	volumeErrChan := make(chan error)
	go volume.Update(quitChan, volumeValueChan, volumeErrChan)

	for {
		select {
		case value := <-hostValueChan:
			stdout.host = value
			drift = true
		case value := <-temperatureValueChan:
			stdout.temperature = value
			drift = true
		case value := <-batteryValueChan:
			stdout.battery = value
			drift = true
		case value := <-volumeValueChan:
			stdout.volume = value
			drift = true
		case err := <-hostErrChan:
			fmt.Println(err)
		case err := <-temperatureErrChan:
			fmt.Println(err)
		case err := <-batteryErrChan:
			fmt.Println(err)
		case err := <-volumeErrChan:
			fmt.Println(err)
		case <-time.Tick(time.Second):
			drift = true
		}

		if drift {
			stdout.write()
			drift = false
		}
	}

	close(quitChan)
}
