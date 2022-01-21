package battery

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/omeid/upower-notify/upower"
)

type Battery struct {
	device string
	stats  upower.Update
}

func New(device string) (battery Battery) {
	battery.device = device
	return Battery{}
}

func (battery *Battery) Update(quitChan chan struct{}, valueChan chan string, errChan chan error) {
	for range time.Tick(time.Millisecond * 250) {
		select {
		case <-quitChan:
			return
		default:
			upower, err := upower.New(battery.device)
			if err != nil {
				errChan <- err
				return
			}

			stats, err := upower.Get()
			if err != nil {
				errChan <- err
				return
			}

			if !reflect.DeepEqual(stats, battery.stats) {
				battery.stats = stats
				valueChan <- battery.str()
			}
		}
	}
}

func (battery *Battery) str() string {
	batteryStr := fmt.Sprintf("%.f", battery.stats.Percentage) + "%"

	var batteryIcon, batteryStatusIcon string
	if battery.stats.Percentage < 0.3 {
		batteryIcon = "🪫"
	} else {
		batteryIcon = "🔋"
	}

	switch battery.stats.State {
	case upower.Charging:
		batteryStatusIcon = "🔌" + batteryIcon
	case upower.Discharging:
		batteryStatusIcon = batteryIcon
	case upower.Empty:
		batteryStatusIcon = "💔"
	case upower.FullCharged:
		batteryStatusIcon = "🔌"
	case upower.Unknown:
		batteryStatusIcon = "❓"
	}

	batteryList := []string{batteryStatusIcon, batteryStr}
	return strings.Join(batteryList, " ")
}
