package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gofiber/fiber"
)

func writeAddrToCfg(port int) {
	cfgPath := filepath.Join("..", "distributor", "servers.cfg")

	data := []byte(fmt.Sprintf("http://localhost:%v\n", port))

	old, _ := ioutil.ReadFile(cfgPath)

	result := append(old, data...)

	ioutil.WriteFile(cfgPath, result, 0644)
}

func main() {
	var port int

	flag.IntVar(&port, "port", 0, "port of service instance")
	flag.Parse()

	writeAddrToCfg(port)

	if port == 0 {
		log.Fatal("-port is null")
		return
	}

	currReqsInProcess := 0

	s := fiber.New()

	s.Get("/", func(c *fiber.Ctx) {
		currReqsInProcess++

		defer func() {
			currReqsInProcess--
		}()

		time.Sleep(5 * time.Second)

		c.SendStatus(http.StatusOK)
	})

	s.Get("/reqs", func(c *fiber.Ctx) {
		c.SendString(strconv.Itoa(currReqsInProcess))
	})

	log.Fatal(s.Listen(port))
}
