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
	log.Print("Spamer started")

	for {
		go sendReq("http://request-distributor-dist-1:61337/")

		time.Sleep(100 * time.Millisecond)
	}
}
