package runner

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	query1 = `{ "query": { "range": { "age": { "gte": 2000, "lte": 10000, "boost": 2 } } }, "aggs": { "age_agg": { "terms": { "field": "age", "size": 1000 } } }, "size": 1 }`
	query2 = `{"query":{"match_all":{}},"aggs":{"last_updated":{"histogram":{"field":"last_updated","interval":10000}}},"size":1}`
	index  = "test_data"
)

type Task interface {
	Request() (string, time.Duration)
}

type HttpTask struct {
	name, url, data string
}

type User struct {
	tasks []Task
}

func (t HttpTask) Request() (string, time.Duration) {
	req, err := http.NewRequest("POST", t.url, bytes.NewBuffer([]byte(t.data)))
	if err != nil {
		log.Fatal("Request error: ", err)
	}
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	client := &http.Client{}
	start := time.Now()
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal("Request error: ", err)
	}
	return t.name, time.Since(start)
}

func (u User) RunTasks() {
	for _, t := range u.tasks {
		fmt.Println(t.Request())
	}
}

func Run(host string, numUsers int) {
	url := host + "/test_data/_search"
	task1 := HttpTask{"Task 1", url, query1}
	task2 := HttpTask{"Task 2", url, query2}
	user := User{[]Task{task1, task2}}
	user.RunTasks()

}
