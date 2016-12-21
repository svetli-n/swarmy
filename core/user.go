package core

import (
	"time"

	"github.com/svetli-n/swarmy/stats"
)

type User interface {
	RunTasks(stop chan bool)
}

type WebUser struct {
	Tasks []Task
}

func (u WebUser) RunTasks(stop chan bool) {
	times := make(chan map[string]time.Duration)
	stats := stats.Stats{}
	go stats.Collect(stop, times)
	for _, t := range u.Tasks {
		go t.Request(stop, times)
	}
}
