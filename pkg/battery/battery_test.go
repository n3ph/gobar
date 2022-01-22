package battery

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/omeid/upower-notify/upower"
)

func TestNew(t *testing.T) {
	type args struct {
		device string
	}
	tests := []struct {
		name        string
		args        args
		wantBattery Battery
		wantErr     bool
	}{
		{"emptyStr", args{device: ""}, Battery{device: &upower.UPower{}, stats: upower.Update{}}, true},
		{"dummyDevice", args{device: "dummyDevice"}, Battery{device: &upower.UPower{}, stats: upower.Update{}}, false},
	}
	for _, tt := range tests {
		if runtime.GOOS != "linux" {
			t.Skip("Skipping tests for linux based upower/dbus implementation")
		}
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

func TestBattery_Update(t *testing.T) {
	type args struct {
		quitChan  chan struct{}
		valueChan chan string
		errChan   chan error
	}

	battery, err := New("dummy")
	if err != nil {
		t.Errorf("nonoo")
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
		{"dummy123", &battery, param},
	}
	for _, tt := range tests {
		if runtime.GOOS != "linux" {
			t.Skip("Skipping tests for linux based upower/dbus implementation")
		}
		t.Run(tt.name, func(t *testing.T) {
			go tt.battery.Update(tt.args.quitChan, tt.args.valueChan, tt.args.errChan)
		})

		select {
		case err := <-tt.args.errChan:
			t.Error(err)
		}
		close(tt.args.quitChan)
	}
}

func TestBattery_str(t *testing.T) {
	tests := []struct {
		name    string
		battery *Battery
		want    string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.battery.str(); got != tt.want {
				t.Errorf("Battery.str() = %v, want %v", got, tt.want)
			}
		})
	}
}
