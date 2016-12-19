package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/svetli-n/swarmy/runner"
)

type Runner struct {
	Host        string
	runningTest bool
	stop        chan bool
}

func (ru *Runner) webHandler(w http.ResponseWriter, r *http.Request) {
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
			if !ru.runningTest {
				ru.stop = make(chan bool)
				fmt.Println("runTest, !runningTest")
				runner.Run(ru.Host, n, ru.stop)
				ru.runningTest = true
			}
		} else {
			if ru.runningTest {
				fmt.Println("Stop")
				close(ru.stop)
			}
		}
	} else {
		t, _ := template.ParseFiles("web/templates/index.html")
		t.Execute(w, nil)
	}
}

func (ru *Runner) Run() {
	http.HandleFunc("/", ru.webHandler)
	http.ListenAndServe(":9999", nil)
}
