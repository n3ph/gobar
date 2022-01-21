package pulseaudio

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want Pulseaudio
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

func TestPulseaudio_Update(t *testing.T) {
	type args struct {
		value chan string
	}
	tests := []struct {
		name string
		pa   *Pulseaudio
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pa.Update(tt.args.value)
		})
	}
}

func TestPulseaudio_Get(t *testing.T) {
	tests := []struct {
		name       string
		pa         *Pulseaudio
		wantOutput string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := tt.pa.Get(); gotOutput != tt.wantOutput {
				t.Errorf("Pulseaudio.Get() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
