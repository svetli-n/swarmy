//Package ptest contains the performance test specific code.
//One needs to define tasks in func makeTasks. Each tasks makes a request
//to a urls with certain data which should come from PostData
package ptest

import (
	"bufio"
	"log"
	"os"

	c "github.com/svetli-n/swarmy/core"
)

func makeTasks(host string, postData *chan string) []c.Task {
	path := "/test_data/_search"
	url := host + path
	task1 := c.HttpTask{"Task 1", url, postData, 1000}
	return []c.Task{task1}
}

func MakeUser(host string, lines *chan string) c.User {
	tasks := makeTasks(host, lines)
	return c.WebUser{tasks}
}

func PostData(postData *chan string) {
	file, err := os.Open("/home/svetlin/workspace/go/src/github.com/svetli-n/swarmy/ptest/data/query3.txt")
	if err != nil {
		log.Fatal("File error", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		*postData <- scanner.Text()
	}
	close(*postData)
	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanning file", err)
	}
}
