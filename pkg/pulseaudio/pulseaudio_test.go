package pulseaudio

import (
	"runtime"
	"testing"

	"github.com/lawl/pulseaudio"
)

func TestNew(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("Skipping tests for linux based pulseaudio implementation")
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
			_, err := New()
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// pa.client != gitPa.client :-(
			// if !reflect.DeepEqual(gotPa, tt.wantPa) {
			// 	t.Errorf("New() = %v, want %v", gotPa, tt.wantPa)
			// }
		})
	}
}

// func BenchmarkNew(b *testing.B) {
// 	if runtime.GOOS != "linux" {
// 		b.Skip("Skipping tests for linux based pulseaudio implementation")
// 	}

// 	for n := 0; n < b.N; n++ {
// 		New()
// 	}
// }

// func TestPulseaudio_Update(t *testing.T) {
// 	if runtime.GOOS != "linux" {
// 		t.Skip("Skipping tests for linux based pulseaudio implementation")
// 	}

// 	type args struct {
// 		quit     chan struct{}
// 		duration time.Duration
// 		value    chan string
// 		err      chan error
// 	}

// 	pa, err := New()
// 	if err != nil {
// 		t.Errorf("Unable to get pulseaudio object: %s", err)
// 	}
// 	paArgs := args{}
// 	paArgs.quit = make(chan struct{})
// 	paArgs.duration = time.Millisecond
// 	paArgs.value = make(chan string)
// 	paArgs.err = make(chan error)

// 	tests := []struct {
// 		name string
// 		pa   *Pulseaudio
// 		args args
// 	}{
// 		{"goroutine", &pa, paArgs},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.pa.Update(tt.args.quit, tt.args.duration, tt.args.value, tt.args.err)

// 			loop := true
// 			for loop {
// 				select {
// 				case err := <-tt.args.err:
// 					t.Error(err)
// 				case value := <-tt.args.value:
// 					if !(len(value) > 0) {
// 						t.Errorf("Unable to retrieve pulseaudio string")
// 					}
// 					loop = false
// 				}
// 			}
// 			close(tt.args.quit)
// 			select {
// 			case _, ok := (<-tt.args.value):
// 				if ok {
// 					t.Errorf("groutine not cleaned up properly")
// 				}
// 				break
// 			default:
// 			}
// 		})
// 	}
// }

func TestStr(t *testing.T) {
	tests := []struct {
		name       string
		pa         *Pulseaudio
		wantOutput string
	}{
		{"emptyStr", &Pulseaudio{}, "🔇 0%"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := tt.pa.str(); gotOutput != tt.wantOutput {
				t.Errorf("Pulseaudio.str() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func BenchmarkStr(b *testing.B) {
	if runtime.GOOS != "linux" {
		b.Skip("Skipping tests for linux based pulseaudio implementation")
	}

	pa, err := New()
	if err != nil {
		b.Errorf("Unable to get pulseaudio object: %s", err)
	}

	for n := 0; n < b.N; n++ {
		pa.str()
	}
}
