package runner

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/svetli-n/swarmy/stats"
)

const (
	query1 = `{ "query": { "range": { "age": { "gte": 2000, "lte": 10000, "boost": 2 } } }, "aggs": { "age_agg": { "terms": { "field": "age", "size": 1000 } } }, "size": 1 }`
	query2 = `{"query":{"match_all":{}},"aggs":{"last_updated":{"histogram":{"field":"last_updated","interval":10000}}},"size":1}`
	index  = "test_data"
)

type Task interface {
	Request(chan bool, chan<- map[string]time.Duration) (string, time.Duration)
}

type HttpTask struct {
	name, url, data string
}

type WebUser struct {
	tasks []Task
}

func (t HttpTask) Request(stop chan bool, times chan<- map[string]time.Duration) (string, time.Duration) {
	for {
		select {
		case <-stop:
			fmt.Println("Stopping")
			return "", 0
		default:
			req, err := http.NewRequest("POST", t.url, bytes.NewBuffer([]byte(t.data)))
			if err != nil {
				log.Fatal("Request error: ", err)
			}
			req.Header.Add("If-None-Match", `W/"wyzzy"`)
			client := &http.Client{}
			start := time.Now()
			resp, err := client.Do(req)
			_ = resp
			if err != nil {
				log.Fatal("Request error: ", err)
			}
			times <- map[string]time.Duration{t.name: time.Since(start)}
			time.Sleep(time.Second * 1)
		}
	}
}

func (u WebUser) RunTasks(stop chan bool) {
	times := make(chan map[string]time.Duration)
	stats := stats.Stats{}
	go stats.Collect(stop, times)
	for _, t := range u.tasks {
		go t.Request(stop, times)
	}
}

func Run(host string, numUsers int, stop chan bool) {
	path := "/test_data/_search"
	url := host + path
	task1 := HttpTask{"Task 1", url, query1}
	task2 := HttpTask{"Task 2", url, query2}
	u := WebUser{[]Task{task1, task2}}
	u.RunTasks(stop)

}
