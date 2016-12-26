package runner

import (
	"time"

	t "github.com/svetli-n/swarmy/ptest"
	s "github.com/svetli-n/swarmy/stats"
)

func Run(host string, numUsers int, stop chan bool) {
	postData := make(chan string)
	go t.PostData(&postData)
	merger := s.Merger{}
	m := make(chan map[string][]time.Duration)
	go merger.Run(m)
	for i := 0; i < numUsers; i++ {
		go func(host string, stop chan bool) {
			u := t.MakeUser(host, &postData)
			u.RunTasks(stop, m)
		}(host, stop)
	}
}
