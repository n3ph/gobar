package pulseaudio

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/lawl/pulseaudio"
)

type Pulseaudio struct {
	init   bool
	level  float32
	mute   bool
	client *pulseaudio.Client
}

func New() (pa Pulseaudio, err error) {
	client, err := pulseaudio.NewClient()
	if err != nil {
		return pa, err
	}

	pa.client = client
	return pa, nil
}

func (pa *Pulseaudio) Update(quit chan struct{}, duration time.Duration, value chan string, err chan error) {
	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for range ticker.C {
		select {
		case <-quit:
			pa.client.Close()
			return
		default:
			var _err error

			pa_new := &Pulseaudio{}
			pa_new.level, _err = pa.client.Volume()
			if _err != nil {
				err <- _err
				return
			}

			pa_new.mute, _err = pa.client.Mute()
			if _err != nil {
				err <- _err
				return
			}

			if pa.level != pa_new.level || pa.mute != pa_new.mute {
				pa.level = pa_new.level
				pa.mute = pa_new.mute
				value <- pa.str()
			}

			if !pa.init {
				value <- pa.str()
				pa.init = true
			}
		}
	}
}

func (pa *Pulseaudio) str() (output string) {
	volumeStr := fmt.Sprintf("%.f", math.Ceil(float64(pa.level)*100)) + "%"

	var volumeIcon string
	switch {
	case pa.level == 0:
		volumeIcon = "🔇"
	case pa.level > 0 && pa.level < 0.3:
		volumeIcon = "🔈"
	case pa.level > 0.3 && pa.level < 0.7:
		volumeIcon = "🔉"
	case pa.level > 0.7:
		volumeIcon = "🔊"
	}

	if pa.mute {
		volumeIcon = "🔇"
	}

	volumeList := []string{volumeIcon, volumeStr}
	return strings.Join(volumeList, " ")
}
