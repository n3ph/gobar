package host

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want Host
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

func TestHost_Update(t *testing.T) {
	type args struct {
		drift chan string
	}
	tests := []struct {
		name string
		host *Host
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.host.Update(tt.args.drift)
		})
	}
}

func TestHost_Get(t *testing.T) {
	tests := []struct {
		name string
		host *Host
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.host.Get(); got != tt.want {
				t.Errorf("Host.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
