package requests

import (
	"bytes"
	"log"
	"net/http"
)

func Post(name, url, data string) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		log.Fatal("Request error: ", err)
	}
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	client := &http.Client{}
	resp, err := client.Do(req)
	_ = resp
	if err != nil {
		log.Fatal("Request error: ", err)
	}
}

func Get(name, url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Request error: ", err)
	}
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	client := &http.Client{}
	resp, err := client.Do(req)
	_ = resp
	if err != nil {
		log.Fatal("Request error: ", err)
	}
}
