package pulseaudio

import (
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/lawl/pulseaudio"
)

func TestNew(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("Skipping tests for linux based dbus/upower implementation")
	}

	client, err := pulseaudio.NewClient()
	if err != nil {
		t.Errorf("Failed to build pulseaudio client: %s", err)
	}

	pa := Pulseaudio{}
	pa.client = client

	tests := []struct {
		name    string
		wantPa  Pulseaudio
		wantErr bool
	}{
		{"generic", pa, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPa, err := New()
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPa, tt.wantPa) {
				t.Errorf("New() = %v, want %v", gotPa, tt.wantPa)
			}
		})
	}
}

func TestPulseaudio_Update(t *testing.T) {
	type args struct {
		quit     chan struct{}
		duration time.Duration
		value    chan string
		err      chan error
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
			tt.pa.Update(tt.args.quit, tt.args.duration, tt.args.value, tt.args.err)
		})
	}
}

func TestPulseaudio_str(t *testing.T) {
	tests := []struct {
		name       string
		pa         *Pulseaudio
		wantOutput string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := tt.pa.str(); gotOutput != tt.wantOutput {
				t.Errorf("Pulseaudio.str() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
