package main

import (
	"flag"

	"github.com/svetli-n/swarmy/web"
)

var host = flag.String("host", "", "Host to run test aginst")

func main() {
	flag.Parse()
	runner := web.Runner{Host: *host}
	runner.Run()
}
