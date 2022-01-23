package battery

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/omeid/upower-notify/upower"
)

func TestNew(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("Skipping tests for linux based dbus/upower implementation")
	}

	type args struct {
		device string
	}

	device := "dummyDevice"
	dbusDevice, err := upower.New(device)
	if err != nil {
		t.Errorf("Unable to get dbus device object")
	}

	dbusBattery := Battery{}
	dbusBattery.device = dbusDevice

	tests := []struct {
		name        string
		args        args
		wantBattery Battery
		wantErr     bool
	}{
		{"emptyStr", args{device: ""}, Battery{device: nil, stats: upower.Update{}}, true},
		{device, args{device: device}, dbusBattery, false},
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

func BenchmarkTest(b *testing.B) {
	if runtime.GOOS != "linux" {
		b.Skip("Skipping tests for linux based dbus/upower implementation")
	}

	for n := 0; n < b.N; n++ {
		New("dummyDevice")
	}
}

func TestBattery_Update(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("Skipping tests for linux based upower/dbus implementation")
	}

	type args struct {
		quitChan  chan struct{}
		valueChan chan string
		errChan   chan error
	}

	battery, err := New("battery_BAT0")
	if err != nil {
		t.Errorf("Unable to get dbus battery object")
	}
	param := args{}
	param.quitChan = make(chan struct{})
	param.valueChan = make(chan string)
	param.errChan = make(chan error)

	tests := []struct {
		name    string
		battery *Battery
		args    args
	}{
		{"goroutine", &battery, param},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			go tt.battery.Update(tt.args.quitChan, tt.args.valueChan, tt.args.errChan)

			loop := true
			for loop {
				select {
				case err := <-tt.args.errChan:
					t.Error(err)
				case value := <-tt.args.valueChan:
					if !(len(value) > 0) {
						t.Errorf("Unable to retrieve battery string")
					}
					loop = false
				}
			}
			close(tt.args.quitChan)
			select {
			case _, ok := (<-tt.args.valueChan):
				if ok {
					t.Errorf("groutine not cleaned up properly")
				}
				break
			default:
			}
		})
	}
}

func BenchmarkBattery_Update(b *testing.B) {
	if runtime.GOOS != "linux" {
		b.Skip("Skipping tests for linux based dbus/upower implementation")
	}

	type args struct {
		quitChan  chan struct{}
		valueChan chan string
		errChan   chan error
	}

	battery, err := New("battery_BAT0")
	if err != nil {
		b.Errorf("Unable to get dbus battery object")
	}

	param := args{}
	param.quitChan = make(chan struct{})
	param.valueChan = make(chan string)
	param.errChan = make(chan error)

	type quit struct{}
	quitStruct := quit{}
	for n := 0; n < b.N; n++ {
		go battery.Update(param.quitChan, param.valueChan, param.errChan)
		param.quitChan <- quitStruct
	}
}

func TestBattery_str(t *testing.T) {
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

func BenchmarkBattery_str(b *testing.B) {
	if runtime.GOOS != "linux" {
		b.Skip("Skipping tests for linux based dbus/upower implementation")
	}

	battery, err := New("dummyDevice")
	if err != nil {
		b.Errorf("Unable to get dbus battery object")
	}

	for n := 0; n < b.N; n++ {
		_ = battery.str()
	}
}
