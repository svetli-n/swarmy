package runner

func Run(host string, numUsers int, stop chan bool) {
	for i := 0; i < numUsers; i++ {
		go func(host string, stop chan bool) {
			makeUser(host).runTasks(stop)
		}(host, stop)
	}
}
