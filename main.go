package main

import (
	"flag"
	"log"

	"github.com/svetli-n/swarmy/web"
)

var host string

func init() {
	flag.StringVar(&host, "host", "", "Host to run test aginst")
	flag.Parse()
	if host == "" {
		log.Fatal("No host name")
	}
}

func main() {
	web.Host = host
	web.Run()
}
