package temperature

import (
	"fmt"
	"strings"
	"time"

	"github.com/shirou/gopsutil/host"
)

type Temperature struct {
	sensor host.TemperatureStat
	value  float64
}

func New(device string) (temperature Temperature, err error) {

	if !(len(device) > 0) {
		err = fmt.Errorf("BUG: device is empty string")
		return
	}

	var sensors []host.TemperatureStat
	sensors, err = host.SensorsTemperatures()
	if err != nil {
		return
	}

	for _, sensor := range sensors {
		if sensor.SensorKey == device {
			temperature.sensor = sensor
			return
		}
	}

	err = fmt.Errorf("get temperature stats: device not found: %s", device)
	return
}

func (temperature *Temperature) Update(quit chan struct{}, duration time.Duration, value chan string, err chan error) {
	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for range ticker.C {
		select {
		case <-quit:
			return
		default:
			temperature_new := Temperature{}
			temperature_new.value = temperature.sensor.Temperature
			if temperature.value != temperature_new.value {
				temperature.value = temperature_new.value
				value <- temperature.str()
			}
		}
	}
}

func (temperature *Temperature) str() string {
	temperatureStr := fmt.Sprintf("%.f", temperature.value) + "°C"
	temperatureIcon := "🌡️"

	temperatureList := []string{temperatureIcon, temperatureStr}
	return strings.Join(temperatureList, " ")
}
