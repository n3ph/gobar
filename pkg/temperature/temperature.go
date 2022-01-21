package temperature

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/shirou/gopsutil/host"
)

type Temperature struct {
	device string
	value  float64
}

func New(device string) (temperature Temperature) {
	temperature.device = device
	return
}

func (temperature *Temperature) Update(quitChan chan struct{}, valueChan chan string, errChan chan error) {
	for range time.Tick(time.Second) {
		select {
		case <-quitChan:
			return
		default:
			temperature_new := &Temperature{}
			sensors, err := host.SensorsTemperatures()
			if err != nil {
				errChan <- err
				return
			}

			for _, sensor := range sensors {
				if sensor.SensorKey == temperature.device {
					temperature_new.value = sensor.Temperature
				} else {
					errChan <- fmt.Errorf("get temperature stats: device not found: %s", temperature.device)
					return
				}
			}

			if !reflect.DeepEqual(temperature, temperature_new) {
				temperature.value = temperature_new.value
				valueChan <- temperature.str()
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
