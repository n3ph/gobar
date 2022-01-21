package timestamp

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want Timestamp
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

func TestTimestamp_Update(t *testing.T) {
	type args struct {
		value chan string
	}
	tests := []struct {
		name      string
		timestamp *Timestamp
		args      args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.timestamp.Update(tt.args.value)
		})
	}
}

func TestTimestamp_Get(t *testing.T) {
	tests := []struct {
		name      string
		timestamp *Timestamp
		want      string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.timestamp.Get(); got != tt.want {
				t.Errorf("Timestamp.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
