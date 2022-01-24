package battery

import (
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/omeid/upower-notify/upower"
)

func TestNew(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("Skipping tests for linux based dbus/upower implementation")
	}

	type args struct {
		device string
	}

	tests := []struct {
		name        string
		args        args
		wantBattery Battery
		wantErr     bool
	}{
		{"empty", args{device: ""}, Battery{device: nil, stats: upower.Update{}}, true},
		{"unknown", args{device: "unknown"}, Battery{device: nil, stats: upower.Update{}}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBattery, err := New(tt.args.device)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBattery, tt.wantBattery) {
				t.Errorf("New() = %v, want %v", gotBattery, tt.wantBattery)
			}
		})
	}
}

func BenchmarkNew(b *testing.B) {
	if runtime.GOOS != "linux" {
		b.Skip("Skipping tests for linux based dbus/upower implementation")
	}

	for n := 0; n < b.N; n++ {
		New("dummyDevice")
	}
}

func TestUpdate(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("Skipping tests for linux based dbus/upower implementation")
	}

	type args struct {
		quit     chan struct{}
		duration time.Duration
		value    chan string
		err      chan error
	}

	battery, err := New("battery_BAT0")
	if err != nil {
		t.Errorf("Unable to get dbus battery object: %s", err)
	}
	batteryArgs := args{}
	batteryArgs.quit = make(chan struct{})
	batteryArgs.duration = time.Millisecond
	batteryArgs.value = make(chan string)
	batteryArgs.err = make(chan error)

	tests := []struct {
		name    string
		battery *Battery
		args    args
	}{
		{"goroutine", &battery, batteryArgs},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			go tt.battery.Update(tt.args.quit, tt.args.duration, tt.args.value, tt.args.err)

			loop := true
			for loop {
				select {
				case err := <-tt.args.err:
					t.Error(err)
				case value := <-tt.args.value:
					if !(len(value) > 0) {
						t.Errorf("Unable to retrieve battery string")
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
	if runtime.GOOS != "linux" {
		b.Skip("Skipping tests for linux based dbus/upower implementation")
	}

	type args struct {
		quit     chan struct{}
		duration time.Duration
		value    chan string
		err      chan error
	}

	battery, err := New("battery_BAT0")
	if err != nil {
		b.Errorf("Unable to get dbus battery object %s", err)
	}

	batteryArgs := args{}
	batteryArgs.quit = make(chan struct{})
	batteryArgs.duration = time.Millisecond
	batteryArgs.value = make(chan string)
	batteryArgs.err = make(chan error)

	type quit struct{}
	quitStruct := quit{}
	for n := 0; n < b.N; n++ {
		go battery.Update(batteryArgs.quit, batteryArgs.duration, batteryArgs.value, batteryArgs.err)
		batteryArgs.quit <- quitStruct
	}
}

func TestStr(t *testing.T) {
	tests := []struct {
		name    string
		battery *Battery
		want    string
	}{
		{"emptyStr", &Battery{}, "❓ 0%"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.battery.str(); got != tt.want {
				t.Errorf("Battery.str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkStr(b *testing.B) {
	if runtime.GOOS != "linux" {
		b.Skip("Skipping tests for linux based dbus/upower implementation")
	}

	battery, err := New("battery_BAT0")
	if err != nil {
		b.Errorf("Unable to get dbus battery object: %s", err)
	}

	for n := 0; n < b.N; n++ {
		battery.str()
	}
}
