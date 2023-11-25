package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func getProtocol(target string) string {
    if strings.Contains(target, ":443") {
        return "https"
    } else {
        return "http"
    }
}

func main() {
    app := fiber.New()

    app.All("*", func (c *fiber.Ctx) error {
		// headers := c.Request().Header

		// Iterate over all headers and print them
		// headers.VisitAll(func(key, value []byte) {
		// 	c.Response().Header.Add(string(key), string(value))
		// })
		url := c.Request().URI()
		println(string(url.FullURI()))
		target := string(url.LastPathSegment())
		protocol := getProtocol(target)
		formatted := fmt.Sprintf("%s://%s", protocol, target)

		println(formatted)

		target_response, err := http.Get(formatted)

		if err != nil {
			return c.SendStatus(400)
		}

		// Iterate over all headers and print them
		for key, values := range target_response.Header {
			for _, value := range values {
				c.Response().Header.Add(string(key), string(value))
			}
		}

		defer target_response.Body.Close()
		c.Status(target_response.StatusCode)
		// c.Response().Header.Add()
		body, _ := io.ReadAll(target_response.Body)
        return c.Send(body)
    })

    log.Fatal(app.Listen(":3000"))
}
