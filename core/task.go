package core

import (
	"fmt"
	"time"

	r "github.com/svetli-n/swarmy/requests"
)

type Task interface {
	Request(chan bool, chan<- map[string]time.Duration) (string, time.Duration)
}

type HttpTask struct {
	Name, Url   string
	Data        *chan string
	SleepMillis time.Duration
}

func (t HttpTask) post(name, url, data string) {
	r.Post(name, url, data)
}

func (t HttpTask) Request(stop chan bool, times chan<- map[string]time.Duration) (string, time.Duration) {
	for {
		select {
		case <-stop:
			fmt.Println("Stopping")
			return "", 0
		default:
			start := time.Now()
			data := <-*t.Data
			t.post(t.Name, t.Url, data)
			times <- map[string]time.Duration{t.Name: time.Since(start)}
			time.Sleep(time.Millisecond * t.SleepMillis)
		}
	}
}
