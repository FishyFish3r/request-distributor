package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber"
)

func getServerLoad(addr string) int {
	req, err := http.NewRequest("GET", fmt.Sprintf("%v/reqs", addr), nil)

	if err != nil {
		return -1
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return -1
	}

	str, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return -1
	}

	load, _ := strconv.Atoi(string(str))

	return load
}

func getAddrsFromFile(filename string) []string {
	data, _ := ioutil.ReadFile("servers.cfg")

	addrs := strings.Split(string(data), "\n")

	list := make([]string, 0)

	for _, addr := range addrs {
		addr := strings.TrimSpace(addr)

		if addr != "" {
			list = append(list, addr)
		}
	}

	return list
}

func getBestServer(servs []string) int {
	bestLoad := math.MaxInt64
	bestServer := 0

	for id, addr := range servs {
		load := getServerLoad(addr)

		if load >= 0 {
			if load < bestLoad {
				log.Printf("Server %v load: %v", addr, load)
				bestLoad = load
				bestServer = id
			}
		}
	}

	return bestServer
}

func sendToServer(addr string) error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%v/", addr), nil)

	if err != nil {
		return err
	}

	_, err = http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	return nil
}

func deleteServerFromList(servs *[]string, id int) {
	if id < len(*servs) && id >= 0 {
		*servs = append((*servs)[:id], (*servs)[id+1:]...)
	}
}

func main() {
	s := fiber.New()

	s.Get("/", func(c *fiber.Ctx) {
		servs := getAddrsFromFile("servsers.cfg")

		if len(servs) > 0 {
			serverId := getBestServer(servs)

			if serverId >= 0 && serverId < len(servs) {
				err := sendToServer(servs[serverId])
				if err != nil {
					log.Printf("Server closed: %v", servs[serverId])
				} else {
					c.SendStatus(http.StatusOK)
				}
			}
		}
	})

	log.Fatal(s.Listen(61337))
}
