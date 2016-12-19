package main

import (
	"flag"
	"os"

	"github.com/svetli-n/swarmy/web"
)

var host = flag.String("host", "", "Host to run test aginst")

// Parses command line options and creates web.Runner
// The web.Runner listens on a port and binds a web.Runner's webHandler
// The webHandler starts tests runner
// The tests runner creates test user with a task list. Each task is a separate request
// The tests runner runs the task list. Per task it creates a stats.Collect-or which receives stats from the running tasks.

func main() {
	flag.Parse()
	if err := web.Run(*host); err != nil {
		os.Exit(1)
	}
}
