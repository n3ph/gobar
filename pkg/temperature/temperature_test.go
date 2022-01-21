package temperature

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want Temperature
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

func TestTemperature_Update(t *testing.T) {
	type args struct {
		drift chan string
	}
	tests := []struct {
		name        string
		temperature *Temperature
		args        args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.temperature.Update(tt.args.drift)
		})
	}
}

func TestTemperature_Get(t *testing.T) {
	tests := []struct {
		name        string
		temperature *Temperature
		want        string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.temperature.Get(); got != tt.want {
				t.Errorf("Temperature.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
