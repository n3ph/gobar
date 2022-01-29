package temperature

import (
	"reflect"
	"testing"

	"github.com/shirou/gopsutil/host"
)

func TestNew(t *testing.T) {
	type args struct {
		device string
	}

	tests := []struct {
		name            string
		args            args
		wantTemperature Temperature
		wantErr         bool
	}{
		{"empty", args{device: ""}, Temperature{sensor: host.TemperatureStat{}, value: 0}, true},
		{"unknown", args{device: "unknown"}, Temperature{sensor: host.TemperatureStat{}, value: 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTemperature, err := New(tt.args.device)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTemperature, tt.wantTemperature) {
				t.Errorf("New() = %v, want %v", gotTemperature, tt.wantTemperature)
			}
		})
	}
}

func BenchmarkNew(b *testing.B) {
	var err error
	var sensors []host.TemperatureStat
	sensors, err = host.SensorsTemperatures()
	if err != nil {
		b.Errorf("Unable to get sensor objects: %s", err)
	}

	var sensor string
	if len(sensors) > 0 {
		last := len(sensors) - 1
		sensor = sensors[last].SensorKey
	} else {
		b.Error("Unable to get host temperature object: none avail")
	}

	for n := 0; n < b.N; n++ {
		_, err := New(sensor)
		if err != nil {
			b.Errorf("Unable to get host temperature object: %s", err)
		}
	}
}

func TestStr(t *testing.T) {
	tests := []struct {
		name        string
		temperature *Temperature
		want        string
	}{
		{"emptyStr", &Temperature{}, "🌡️ 0°C"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.temperature.str(); got != tt.want {
				t.Errorf("Temperature.str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkStr(b *testing.B) {
	var err error
	var sensors []host.TemperatureStat
	sensors, err = host.SensorsTemperatures()
	if err != nil {
		b.Errorf("Unable to get sensor objects: %s", err)
	}

	var temperature Temperature
	if len(sensors) > 0 {
		last := len(sensors) - 1
		temperature, err = New(sensors[last].SensorKey)
		if err != nil {
			b.Errorf("Unable to get host temperature object: %s", err)
		}
	} else {
		b.Error("Unable to get host temperature object: none avail")
	}

	for n := 0; n < b.N; n++ {
		temperature.str()
	}
}
