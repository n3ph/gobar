package main

import (
	"fmt"
	"math"
	"time"

	"github.com/mafik/pulseaudio"
	"github.com/omeid/upower-notify/upower"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
)

func date() string {
	currentTime := time.Now()
	return currentTime.Format("02.01.2006 15:04:05")
}

func volume() string {
	paClient, err := pulseaudio.NewClient()
	if err != nil {
		panic(err)
	}
	defer paClient.Close()

	volumeFloat, err := paClient.Volume()
	if err != nil {
		panic(err)
	}
	volumeStr := fmt.Sprintf("%.f", math.Ceil(float64(volumeFloat)*100)) + "%"

	mute, err := paClient.Mute()
	if err != nil {
		panic(err)
	}

	var volumeIcon string
	switch {
	case volumeFloat > 0.8:
		volumeIcon = "🔊"
	case volumeFloat > 0.4:
		volumeIcon = "🔉"
	case volumeFloat > 0.2:
		volumeIcon = "🔈"
	case volumeFloat == 0:
		volumeIcon = "🔇"
	}
	if mute {
		volumeIcon = "🔇"
	}

	return volumeIcon + " " + volumeStr
}

func avgLoad() string {
	loadAvg, err := load.Avg()
	if err != nil {
		panic(err)
	}

	load1 := fmt.Sprintf("%.2f", loadAvg.Load1)
	load5 := fmt.Sprintf("%.2f", loadAvg.Load5)
	load15 := fmt.Sprintf("%.2f", loadAvg.Load15)

	return load1 + " " + load5 + " " + load15
}

func temp() string {
	sensors, err := host.SensorsTemperatures()
	if err != nil {
		panic(err)
	}

	var tempStr string
	for _, sensor := range sensors {
		if sensor.SensorKey == "amdgpu_edge_input" {
			tempStr = fmt.Sprintf("%.f", sensor.Temperature) + "°C"
		}
	}

	tempIcon := "🌡️"
	return tempIcon + " " + tempStr
}

func battery() string {
	up, err := upower.New("battery_BAT0")
	if err != nil {
		panic(err)
	}

	status, err := up.Get()
	if err != nil {
		panic(err)
	}

	batteryStr := fmt.Sprintf("%.f", status.Percentage) + "%"
	var batteryIcon string

	if status.Percentage < 0.3 {
		batteryIcon = "🪫"
	} else {
		batteryIcon = "🔋"
	}

	switch status.State {
	case 0: // unknown
		batteryIcon = "❓"
	case 1: // charging
		batteryIcon = "🔌" + batteryIcon
	case 3: // empty
		batteryIcon = "💔"
	case 4: // fully charged
		batteryIcon = "🔌"
	}

	return batteryIcon + " " + batteryStr
}

func main() {
	var stdout string
	elements := []string{avgLoad(), temp(), battery(), volume(), date()}
	for _, element := range elements {
		stdout += element + " | "
	}

	fmt.Println(stdout)
}
