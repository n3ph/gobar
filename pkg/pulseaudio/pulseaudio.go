package pulseaudio

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"time"

	"github.com/lawl/pulseaudio"
)

type Pulseaudio struct {
	level float32
	mute  bool
}

func New() Pulseaudio {
	return Pulseaudio{}
}

func (pa *Pulseaudio) Update(quitChan chan struct{}, valueChan chan string, errChan chan error) {
	client, err := pulseaudio.NewClient()
	if err != nil {
		errChan <- err
		return
	}
	defer client.Close()

	for range time.Tick(time.Millisecond * 250) {
		select {
		case <-quitChan:
			return
		default:
			var err error
			pa_new := &Pulseaudio{}
			pa_new.level, err = client.Volume()
			if err != nil {
				return
			}
			pa_new.mute, err = client.Mute()
			if err != nil {
				return
			}

			if !reflect.DeepEqual(pa, pa_new) {
				pa.level = pa_new.level
				pa.mute = pa_new.mute
				valueChan <- pa.str()
			}
		}
	}
}

func (pa *Pulseaudio) str() (output string) {
	volumeStr := fmt.Sprintf("%.f", math.Ceil(float64(pa.level)*100)) + "%"

	var volumeIcon string
	switch {
	case pa.level > 0.8:
		volumeIcon = "🔊"
	case pa.level > 0.4:
		volumeIcon = "🔉"
	case pa.level > 0.2:
		volumeIcon = "🔈"
	case pa.level == 0:
		volumeIcon = "🔇"
	}
	if pa.mute {
		volumeIcon = "🔇"
	}

	volumeList := []string{volumeIcon, volumeStr}
	return strings.Join(volumeList, " ")
}
