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
	volume      string
}

func (elements Elements) write() {
	var stdout string

	for _, element := range []string{elements.host, elements.temperature, elements.battery, elements.volume, getTimestamp()} {
		stdout += element + " | "
	}

	fmt.Println(stdout)
}

type args struct {
	duration time.Duration
	value    chan string
	err      chan error
}

func Gobar() {
	var stdout Elements
	var drift bool
	quitChan := make(chan struct{})

	hostArgs := args{}
	hostArgs.duration = time.Millisecond * 250
	hostArgs.value = make(chan string)
	hostArgs.err = make(chan error)
	host := host.New()
	go host.Update(quitChan, hostArgs.duration, hostArgs.value, hostArgs.err)

	temperatureArgs := args{}
	temperatureArgs.duration = time.Millisecond * 250
	temperatureArgs.value = make(chan string)
	temperatureArgs.err = make(chan error)
	temperature, err := temperature.New("amdgpu_edge_input")
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		go temperature.Update(quitChan, temperatureArgs.duration, temperatureArgs.value, temperatureArgs.err)
	}

	batteryArgs := args{}
	batteryArgs.duration = time.Millisecond * 250
	batteryArgs.value = make(chan string)
	batteryArgs.err = make(chan error)
	battery, err := battery.New("battery_BAT0")
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		go battery.Update(quitChan, batteryArgs.duration, batteryArgs.value, batteryArgs.err)
	}

	volumeArgs := args{}
	volumeArgs.duration = time.Millisecond * 250
	volumeArgs.value = make(chan string)
	volumeArgs.err = make(chan error)
	volume, err := pulseaudio.New()
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		go volume.Update(quitChan, volumeArgs.duration, volumeArgs.value, volumeArgs.err)
	}

	for {
		select {
		case value := <-hostArgs.value:
			stdout.host = value
			drift = true
		case err := <-hostArgs.err:
			fmt.Println(err)

		case value := <-temperatureArgs.value:
			stdout.temperature = value
			drift = true
		case err := <-temperatureArgs.err:
			fmt.Println(err)

		case value := <-batteryArgs.value:
			stdout.battery = value
			drift = true
		case err := <-batteryArgs.err:
			fmt.Println(err)

		case value := <-volumeArgs.value:
			stdout.volume = value
			drift = true
		case err := <-volumeArgs.err:
			fmt.Println(err)

		case <-time.Tick(time.Second):
			drift = true
		}

		// TODO: cleanup go routines before exit
		// close(quitChan)
		// close pa.client

		if drift {
			stdout.write()
			drift = false
		}
	}
}
