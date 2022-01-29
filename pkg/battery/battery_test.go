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

	battery := &Battery{}
	for n := 0; n < b.N; n++ {
		battery.str()
	}
}
