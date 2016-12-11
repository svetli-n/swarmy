package web

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/svetli-n/swarmy/runner"
)

var Host string

func webHandler(w http.ResponseWriter, r *http.Request) {
	if numUsers := r.FormValue("numUsers"); numUsers != "" {
		n, err := strconv.Atoi(numUsers)
		if err != nil {
			log.Fatal("numUsers error.")
		}
		runner.Run(Host, n)
	} else {
		t, _ := template.ParseFiles("web/templates/index.html")
		t.Execute(w, nil)
	}
}

func Run() {
	http.HandleFunc("/", webHandler)
	http.ListenAndServe(":9999", nil)
}
