package host

import (
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want Host
	}{
		{"generic", Host{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkNew(b *testing.B) {
	for n := 0; n < b.N; n++ {
		New()
	}
}

func TestUpdate(t *testing.T) {
	type args struct {
		quit     chan struct{}
		duration time.Duration
		value    chan string
		err      chan error
	}

	host := New()
	hostArgs := args{}
	hostArgs.quit = make(chan struct{})
	hostArgs.duration = time.Millisecond
	hostArgs.value = make(chan string)
	hostArgs.err = make(chan error)

	tests := []struct {
		name string
		host *Host
		args args
	}{
		{"goroutine", &host, hostArgs},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go tt.host.Update(tt.args.quit, tt.args.duration, tt.args.value, tt.args.err)

			loop := true
			for loop {
				select {
				case err := <-tt.args.err:
					t.Error(err)
				case value := <-tt.args.value:
					if !(len(value) > 0) {
						t.Errorf("Unable to retrieve host string")
					}
					loop = false
				}

			}
			close(tt.args.quit)
			select {
			case _, ok := (<-tt.args.value):
				if ok {
					t.Errorf("groutine not cleaned up properly")
				}
				break
			default:
			}
		})
	}
}

func BenchmarkUpdate(b *testing.B) {
	type args struct {
		quit     chan struct{}
		duration time.Duration
		value    chan string
		err      chan error
	}

	host := New()
	hostArgs := args{}
	hostArgs.quit = make(chan struct{})
	hostArgs.duration = time.Millisecond
	hostArgs.value = make(chan string)
	hostArgs.err = make(chan error)

	type quit struct{}
	_quit := quit{}
	for n := 0; n < b.N; n++ {
		go host.Update(hostArgs.quit, hostArgs.duration, hostArgs.value, hostArgs.err)
		hostArgs.quit <- _quit
	}
}

func TestStr(t *testing.T) {

	tests := []struct {
		name string
		host *Host
		want string
	}{
		{"generic", &Host{}, "0.00 0.00 0.00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.host.str(); got != tt.want {
				t.Errorf("Host.str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkStr(b *testing.B) {
	host := New()
	for n := 0; n < b.N; n++ {
		_ = host.str()
	}
}
