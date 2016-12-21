package runner

import t "github.com/svetli-n/swarmy/test"

func Run(host string, numUsers int, stop chan bool) {
	for i := 0; i < numUsers; i++ {
		go func(host string, stop chan bool) {
			u := t.MakeUser(host)
			u.RunTasks(stop)
		}(host, stop)
	}
}
