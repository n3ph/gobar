package battery

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/omeid/upower-notify/upower"
)

type Battery struct {
	stats upower.Update
}

func New() Battery {
	return Battery{}
}

func (battery *Battery) Update(drift chan bool, device string) {
	for range time.Tick(time.Millisecond * 250) {
		upower, err := upower.New(device)
		if err != nil {
			panic(err)
		}

		stats, err := upower.Get()
		if err != nil {
			panic(err)
		}

		if !reflect.DeepEqual(stats, battery.stats) {
			battery.stats = stats
			drift <- true
		}
	}
}

func (battery *Battery) Get() string {
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
