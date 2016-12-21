//Package test contains the performance test specific code.
//One needs to define tasks in func makeTasks. Each tasks makes a request
//to a urls with ceratin data
package test

import (
	c "github.com/svetli-n/swarmy/core"
	r "github.com/svetli-n/swarmy/requests"
)

const (
	query1 = `{ "query": { "range": { "age": { "gte": 2000, "lte": 10000, "boost": 2 } } }, "aggs": { "age_agg": { "terms": { "field": "age", "size": 1000 } } }, "size": 1 }`
	query2 = `{"query":{"match_all":{}},"aggs":{"last_updated":{"histogram":{"field":"last_updated","interval":10000}}},"size":1}`
	query3 = `{ "query": { "range": { "age": { "gte": 2000, "lte": 10000, "boost": 2 } } }, "aggs": { "age_agg": { "terms": { "field": "age", "size": 1000 } } }, "size": 1 }`
	index  = "test_data"
)

func post(name, url, data string) {
	r.Post(name, url, data)
}

func makeTasks(host string) []c.Task {
	path := "/test_data/_search"
	url := host + path
	task1 := c.HttpTask{"Task 1", url, query1, post}
	task2 := c.HttpTask{"Task 2", url, query2, post}
	task3 := c.HttpTask{"Task 3", url, query3, post}
	return []c.Task{task1, task2, task3}
}

func MakeUser(host string) c.User {
	tasks := makeTasks(host)
	return c.WebUser{tasks}
}
