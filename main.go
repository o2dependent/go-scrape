package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/gocolly/colly/v2"
)

// INIT IDEA
// Scrape all emails on a site
// Maybe hit the sitemap.xml first and if that fail just hit the page and crawl links
// Start with just nabbing the emails from on page

func main() {
	site := "https://eolsen.dev"

	c := colly.NewCollector()

	emails := []string{}

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		href := h.Attr("href")

		if strings.HasPrefix(href, "mailto:") && !slices.Contains(emails, href) {
			// if mailto put in emails
			emails = append(emails, href)
		} else if strings.HasPrefix(href, site) || strings.HasPrefix(href, "/") {
			// TODO this is for later when doing site crawling if no sitemap.xml
			// put into sites to crawl

		}
	})

	// emailRegex, _ := regexp.Compile("/\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}\b/")
	// c.OnHTML("*", func(h *colly.HTMLElement) {
	// 	text := h.Text

	// 	if emailRegex.Match(text) {

	// 	}

	// })

	c.OnScraped(func(r *colly.Response) {
		for i, email := range emails {
			fmt.Println(i, email)
		}

	})

	c.Visit(site)
}
