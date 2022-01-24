package temperature

import (
	"reflect"
	"testing"
	"time"

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
	for n := 0; n < b.N; n++ {
		New("amdgpu_edge_input")
	}
}

func TestUpdate(t *testing.T) {
	type args struct {
		quit     chan struct{}
		duration time.Duration
		value    chan string
		err      chan error
	}

	var err error
	var sensors []host.TemperatureStat
	sensors, err = host.SensorsTemperatures()
	if err != nil {
		t.Errorf("Unable to get sensor objects: %s", err)
	}

	var temperature Temperature
	if len(sensors) > 0 {
		last := len(sensors) - 1
		temperature, err = New(sensors[last].SensorKey)
		if err != nil {
			t.Errorf("Unable to get host temperature object: %s", err)
		}
	} else {
		t.Error("Unable to get host temperature object: none avail")
	}

	temperatureArgs := args{}
	temperatureArgs.quit = make(chan struct{})
	temperatureArgs.duration = time.Millisecond
	temperatureArgs.value = make(chan string)
	temperatureArgs.err = make(chan error)

	tests := []struct {
		name        string
		temperature *Temperature
		args        args
	}{
		{"goroutine", &temperature, temperatureArgs},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go tt.temperature.Update(tt.args.quit, tt.args.duration, tt.args.value, tt.args.err)

			loop := true
			for loop {
				select {
				case err := <-tt.args.err:
					t.Error(err)
				case value := <-tt.args.value:
					if !(len(value) > 0) {
						t.Errorf("Unable to retrieve temperature string")
					}
					loop = false
				}
			}
			close(tt.args.quit)
			select {
			case _, ok := (<-tt.args.value):
				if ok {
					t.Errorf("groutine not cleaned up properly")
				}
				break
			default:
			}
		})
	}
}

func BenchmarkUpdate(b *testing.B) {
	type args struct {
		quit     chan struct{}
		duration time.Duration
		value    chan string
		err      chan error
	}

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

	temperatureArgs := args{}
	temperatureArgs.quit = make(chan struct{})
	temperatureArgs.duration = time.Millisecond
	temperatureArgs.value = make(chan string)
	temperatureArgs.err = make(chan error)

	type quit struct{}
	quitStruct := quit{}
	for n := 0; n < b.N; n++ {
		go temperature.Update(temperatureArgs.quit, temperatureArgs.duration, temperatureArgs.value, temperatureArgs.err)
		temperatureArgs.quit <- quitStruct
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
