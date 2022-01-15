package temperature

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/shirou/gopsutil/host"
)

type Temperature struct {
	value float64
}

func (temperature *Temperature) Update(drift chan bool) {
	for range time.Tick(time.Second) {
		temperature_new := &Temperature{}
		sensors, err := host.SensorsTemperatures()
		if err != nil {
			panic(err)
		}

		for _, sensor := range sensors {
			if sensor.SensorKey == "amdgpu_edge_input" {
				temperature_new.value = sensor.Temperature
			}
		}

		if !reflect.DeepEqual(temperature, temperature_new) {
			temperature.value = temperature_new.value
			drift <- true
		}
	}
}

func (temperature *Temperature) Get() string {
	temperatureStr := fmt.Sprintf("%.f", temperature.value) + "°C"
	temperatureIcon := "🌡️"

	temperatureList := []string{temperatureIcon, temperatureStr}
	return strings.Join(temperatureList, " ")
}
