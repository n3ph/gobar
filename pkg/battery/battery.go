package battery

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/omeid/upower-notify/upower"
)

type Battery struct {
	device *upower.UPower
	stats  upower.Update
}

func New(device string) (battery Battery, err error) {
	if !(len(device) > 0) {
		err = fmt.Errorf("BUG: device is empty string")
		return
	}

	upowerDevice, err := upower.New(device)
	if err != nil {
		return
	}

	_, err = upowerDevice.Get()
	if err != nil {
		return
	}

	battery.device = upowerDevice
	return
}

func (battery *Battery) Update(quit chan struct{}, duration time.Duration, value chan string, err chan error) {
	for range time.Tick(duration) {
		select {
		case <-quit:
			return
		default:
			stats, _err := battery.device.Get()
			if _err != nil {
				err <- _err
				return
			}

			if !reflect.DeepEqual(stats, battery.stats) {
				battery.stats = stats
				value <- battery.str()
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
