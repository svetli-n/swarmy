package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/svetli-n/swarmy/runner"
)

var Host string
var runningTest bool
var stop = make(chan bool)

func webHandler(w http.ResponseWriter, r *http.Request) {
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
			if !runningTest {
				fmt.Println("runTest, !runningTest")
				runner.Run(Host, n, stop)
				runningTest = true
			}
		} else {
			if runningTest {
				fmt.Println("Stop")
				close(stop)
			}
		}
	} else {
		t, _ := template.ParseFiles("web/templates/index.html")
		t.Execute(w, nil)
	}
}

func Run() {
	http.HandleFunc("/", webHandler)
	http.ListenAndServe(":9999", nil)
}
