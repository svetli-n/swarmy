package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/svetli-n/swarmy/runner"
)

type webRunner struct {
	host        string
	runningTest bool
	stop        chan bool
}

func (wr *webRunner) webHandler(w http.ResponseWriter, r *http.Request) {
	if numUsers := r.FormValue("numUsers"); numUsers != "" {
		n, err := strconv.Atoi(numUsers)
		if err != nil {
			log.Fatal("numUsers error.", err)
		}
		runTest, err := strconv.ParseBool(r.FormValue("runTest"))
		if err != nil {
			log.Fatal("runTest error.", err)
		}
		if runTest {
			if !wr.runningTest {
				wr.stop = make(chan bool)
				fmt.Println("Running")
				runner.Run(wr.host, n, wr.stop)
				wr.runningTest = true
			}
		} else {
			if wr.runningTest {
				fmt.Println("Stop")
				close(wr.stop)
				wr.runningTest = false
			}
		}
	} else {
		t, _ := template.ParseFiles("web/templates/index.html")
		t.Execute(w, nil)
	}
}

func (wr *webRunner) start() error {
	http.HandleFunc("/", wr.webHandler)
	return http.ListenAndServe(":9999", nil)
}

func makeWebRunner(host string) *webRunner {
	return &webRunner{host: host}
}

func handleError(err error) error {
	if err != nil {
		log.Printf("Web error", err)
	}
	return err
}

func Run(host string) error {
	return handleError(makeWebRunner(host).start())
}
