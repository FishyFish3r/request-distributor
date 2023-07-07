package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
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
	bestServer := -1

	for id, addr := range servs {
		load := getServerLoad(addr)

		if load >= 0 {
			if load < bestLoad {
				bestLoad = load
				bestServer = id
			}
		}
	}

	return bestServer
}

func serverLive(addr string) bool {
	return getServerLoad(addr) >= 0
}

func sortServers(l []string) []string {
	servs := make([]string, 0)
	for _, addr := range l {
		if serverLive(addr) {
			servs = append(servs, addr)
		}
	}

	return servs
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

func main() {
	s := fiber.New()

	logfile, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		log.Print(err)
		return
	}

	defer logfile.Close()

	logrus.SetOutput(logfile)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	s.Get("/", func(c *fiber.Ctx) error {
		servs := getAddrsFromFile("servsers.cfg")

		servs = sortServers(servs)

		if len(servs) > 0 {
			serverId := getBestServer(servs)

			if serverId >= 0 && serverId < len(servs) {
				err := sendToServer(servs[serverId])
				if err != nil {
					log.Printf("Server closed: %v", servs[serverId])
					logrus.Info(fmt.Sprintf("Server closed: %v", servs[serverId]))
					return c.SendStatus(http.StatusInternalServerError)
				} else {
					log.Printf("Server %v ok load: %v", servs[serverId], getServerLoad(servs[serverId]))
					logrus.Info(fmt.Sprintf("Server %v ok load: %v", servs[serverId], getServerLoad(servs[serverId])))
					return c.SendStatus(http.StatusOK)
				}
			}
		} else {
			log.Print("Servers down")
			logrus.Info("Servers down")
		}

		return nil
	})

	log.Fatal(s.Listen(":61337"))
}
