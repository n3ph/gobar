package temperature

import (
	"reflect"
	"testing"
	"time"
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
		// TODO: Add test cases.
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

func TestTemperature_Update(t *testing.T) {
	type args struct {
		quit     chan struct{}
		duration time.Duration
		value    chan string
		err      chan error
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
			tt.temperature.Update(tt.args.quit, tt.args.duration, tt.args.value, tt.args.err)
		})
	}
}

func TestTemperature_str(t *testing.T) {
	tests := []struct {
		name        string
		temperature *Temperature
		want        string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.temperature.str(); got != tt.want {
				t.Errorf("Temperature.str() = %v, want %v", got, tt.want)
			}
		})
	}
}
