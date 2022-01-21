package gobar

import "testing"

func TestElements_write(t *testing.T) {
	tests := []struct {
		name     string
		elements Elements
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.elements.write()
		})
	}
}

func TestGobar(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Gobar()
		})
	}
}
