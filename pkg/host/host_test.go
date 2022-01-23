package host

import (
	"reflect"
	"runtime"
	"testing"
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
		quitChan  chan struct{}
		valueChan chan string
		errChan   chan error
	}

	host := New()
	param := args{}
	param.quitChan = make(chan struct{})
	param.valueChan = make(chan string)
	param.errChan = make(chan error)

	tests := []struct {
		name string
		host *Host
		args args
	}{
		{"goroutine", &host, param},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go tt.host.Update(tt.args.quitChan, tt.args.valueChan, tt.args.errChan)

			loop := true
			for loop {
				select {
				case err := <-tt.args.errChan:
					t.Error(err)
				case value := <-tt.args.valueChan:
					if !(len(value) > 0) {
						t.Errorf("Unable to retrieve host string")
					}
					loop = false
				}
			}
			close(tt.args.quitChan)
			select {
			case _, ok := (<-tt.args.valueChan):
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
		quitChan  chan struct{}
		valueChan chan string
		errChan   chan error
	}

	host := New()
	param := args{}
	param.quitChan = make(chan struct{})
	param.valueChan = make(chan string)
	param.errChan = make(chan error)

	type quit struct{}
	quitStruct := quit{}
	for n := 0; n < b.N; n++ {
		go host.Update(param.quitChan, param.valueChan, param.errChan)
		param.quitChan <- quitStruct
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
	if runtime.GOOS != "linux" {
		b.Skip("Skipping tests for linux based dbus/upower implementation")
	}

	host := New()
	for n := 0; n < b.N; n++ {
		_ = host.str()
	}
}
