package runner

func Run(host string, numUsers int, stop chan bool) {
	for i := 0; i < numUsers; i++ {
		go runUser(host, stop)
	}
}

func runUser(host string, stop chan bool) {
	tasks := makeTasks(host)
	u := WebUser{tasks}
	u.runTasks(stop)
}
