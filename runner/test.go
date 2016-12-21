package runner

import (
	"github.com/svetli-n/swarmy/requests"
)

const (
	query1 = `{ "query": { "range": { "age": { "gte": 2000, "lte": 10000, "boost": 2 } } }, "aggs": { "age_agg": { "terms": { "field": "age", "size": 1000 } } }, "size": 1 }`
	query2 = `{"query":{"match_all":{}},"aggs":{"last_updated":{"histogram":{"field":"last_updated","interval":10000}}},"size":1}`
	query3 = `{ "query": { "range": { "age": { "gte": 2000, "lte": 10000, "boost": 2 } } }, "aggs": { "age_agg": { "terms": { "field": "age", "size": 1000 } } }, "size": 1 }`
	index  = "test_data"
)

func makeRequest(name, url, data string) {
	requests.Post(name, url, data)
}

func makeTasks(host string) []Task {
	path := "/test_data/_search"
	url := host + path
	task1 := HttpTask{"Task 1", url, query1}
	task2 := HttpTask{"Task 2", url, query2}
	task3 := HttpTask{"Task 3", url, query3}
	return []Task{task1, task2, task3}
}
