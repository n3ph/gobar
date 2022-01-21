package battery

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want Battery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBattery_Update(t *testing.T) {
	type args struct {
		value  chan string
		device string
	}
	tests := []struct {
		name    string
		battery *Battery
		args    args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.battery.Update(tt.args.value, tt.args.device)
		})
	}
}

func TestBattery_Get(t *testing.T) {
	tests := []struct {
		name    string
		battery *Battery
		want    string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.battery.Get(); got != tt.want {
				t.Errorf("Battery.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
