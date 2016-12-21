package runner

import (
	"fmt"
	"time"
)

type Task interface {
	Request(chan bool, chan<- map[string]time.Duration) (string, time.Duration)
}

type HttpTask struct {
	name, url, data string
}

func (t HttpTask) Request(stop chan bool, times chan<- map[string]time.Duration) (string, time.Duration) {
	for {
		select {
		case <-stop:
			fmt.Println("Stopping")
			return "", 0
		default:
			start := time.Now()
			makeRequest(t.name, t.url, t.data)
			times <- map[string]time.Duration{t.name: time.Since(start)}
			time.Sleep(time.Second * 1)
		}
	}
}
