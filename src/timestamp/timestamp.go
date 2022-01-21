package timestamp

import "time"

type Timestamp struct {
	now time.Time
}

func New() Timestamp {
	return Timestamp{}
}

func (timestamp *Timestamp) Update(value chan string) {
	for range time.Tick(time.Second) {
		timestamp.now = time.Now()
		value <- timestamp.Get()
	}
}

func (timestamp *Timestamp) Get() string {
	return timestamp.now.Format("02.01.2006 15:04:05")
}
