package runner

import (
	"time"

	"github.com/svetli-n/swarmy/stats"
)

type User interface {
	runTasks(stop chan bool)
}

type WebUser struct {
	tasks []Task
}

func (u WebUser) runTasks(stop chan bool) {
	times := make(chan map[string]time.Duration)
	stats := stats.Stats{}
	go stats.Collect(stop, times)
	for _, t := range u.tasks {
		go t.Request(stop, times)
	}
}

func makeUser(host string) User {
	tasks := makeTasks(host)
	return WebUser{tasks}
}
