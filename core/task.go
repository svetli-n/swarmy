package core

import (
	"fmt"
	"time"
)

type Task interface {
	Request(chan bool, chan<- map[string]time.Duration) (string, time.Duration)
}

type HttpTask struct {
	Name, Url, Data string
	F               func(string, string, string)
}

func (t HttpTask) Request(stop chan bool, times chan<- map[string]time.Duration) (string, time.Duration) {
	for {
		select {
		case <-stop:
			fmt.Println("Stopping")
			return "", 0
		default:
			start := time.Now()
			t.F(t.Name, t.Url, t.Data)
			times <- map[string]time.Duration{t.Name: time.Since(start)}
			time.Sleep(time.Second * 1)
		}
	}
}
