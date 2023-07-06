package main

import (
	"log"
	"net/http"
	"time"
)

func sendReq(addr string) {
	req, err := http.NewRequest("GET", addr, nil)

	if err != nil {
		log.Fatal(err)
		return
	}

	http.DefaultClient.Do(req)
}

func main() {
	for {
		go sendReq("http://localhost:61337/")

		time.Sleep(100 * time.Millisecond)
	}
}
