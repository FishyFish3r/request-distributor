package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber"
)

func main() {
	currReqsInProcess := 0

	s := fiber.New()

	s.Get("/", func(c *fiber.Ctx) {
		currReqsInProcess++

		defer func() {
			currReqsInProcess--
		}()
		time.Sleep(3 * time.Second)

		c.SendStatus(http.StatusOK)
	})

	s.Get("/reqs", func(c *fiber.Ctx) {
		c.SendString(strconv.Itoa(currReqsInProcess))
	})

	log.Fatal(s.Listen(os.Getenv("PORT")))
}
