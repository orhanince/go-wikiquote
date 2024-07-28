package handler

import (
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/orhanince/go-wikiquote/src/model"
)

func Home(c *fiber.Ctx) error {
	err := c.Status(fiber.StatusOK).JSON(model.Response{
		Status:  true,
		Message: "Success",
		Data: map[string]interface{}{
			"name": "Home",
		}})

	return err
}

func Check(c *fiber.Ctx) error {
	err := c.Status(fiber.StatusOK).JSON(model.Response{
		Status:  true,
		Message: "Success",
		Data: map[string]interface{}{
			"name": "User",
		}})

	return err
}

func Test(c *fiber.Ctx) error {
	err := c.Status(fiber.StatusOK).JSON(model.Response{
		Status:  true,
		Message: "Success",
		Data: map[string]interface{}{
			"name": "User Test!",
		}})

	return err
}

func Scrape(c *fiber.Ctx) error {
	co := colly.NewCollector()
	url := "https://tr.wikiquote.org/wiki/Albert_Einstein"

	// Set up callbacks to handle scraping events
	co.OnHTML(".mw-content-ltr", func(e *colly.HTMLElement) {
		// Extract data from HTML elements
		ul := e.ChildText("ul>li")

		ul = strings.TrimSpace(ul)
	})

	// Visit the URL and start scraping
	co_err := co.Visit(url)
	if co_err != nil {
		log.Fatal(co_err)
	}

	err := c.Status(fiber.StatusOK).JSON(model.Response{
		Status:  true,
		Message: "Success",
		Data: map[string]interface{}{
			"name": "User Scrape!",
		}})

	return err
}
