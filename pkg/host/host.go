package host

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/shirou/gopsutil/load"
)

type Host struct {
	load1  float64
	load5  float64
	load15 float64
}

func New() Host {
	return Host{}
}

func (host *Host) Update(quit chan struct{}, duration time.Duration, value chan string, err chan error) {
	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for range ticker.C {
		select {
		case <-quit:
			return
		default:
			host_new := &Host{}
			loadAvg, _err := load.Avg()
			if _err != nil {
				err <- _err
				return
			}

			host_new.load1 = loadAvg.Load1
			host_new.load5 = loadAvg.Load5
			host_new.load15 = loadAvg.Load15

			if !reflect.DeepEqual(host, host_new) {
				host.load1 = host_new.load1
				host.load5 = host_new.load5
				host.load15 = host_new.load15
				value <- host.str()
			}
		}
	}
}

func (host *Host) str() string {
	load1Str := fmt.Sprintf("%.2f", host.load1)
	load5Str := fmt.Sprintf("%.2f", host.load5)
	load15Str := fmt.Sprintf("%.2f", host.load15)

	loadList := []string{load1Str, load5Str, load15Str}
	return strings.Join(loadList, " ")
}
