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
	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for range ticker.C {
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
	if battery.stats.Percentage < 30 {
		batteryIcon = "đĒĢ"
	} else {
		batteryIcon = "đ"
	}

	switch battery.stats.State {
	case upower.Charging:
		batteryStatusIcon = "đ" + batteryIcon
	case upower.Discharging:
		batteryStatusIcon = batteryIcon
	case upower.Empty:
		batteryStatusIcon = "đ"
	case upower.FullCharged:
		batteryStatusIcon = "đ"
	case upower.Unknown:
		batteryStatusIcon = "â"
	}

	batteryList := []string{batteryStatusIcon, batteryStr}
	return strings.Join(batteryList, " ")
}
