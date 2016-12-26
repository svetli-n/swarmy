package core

import (
	"time"

	"github.com/svetli-n/swarmy/stats"
)

type User interface {
	RunTasks(stop chan bool, merge chan map[string][]time.Duration)
}

type WebUser struct {
	Tasks []Task
}

func (u WebUser) RunTasks(stop chan bool, merge chan map[string][]time.Duration) {
	times := make(chan map[string]time.Duration)
	stats := stats.Stats{}
	go stats.Collect(stop, times, merge)
	for _, t := range u.Tasks {
		go t.Request(stop, times)
	}
}
