package main

import (
	"os"
	"regexp"
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
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	c := colly.NewCollector(colly.IgnoreRobotsTxt())

	emails := []string{}

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		href := h.Attr("href")

		if strings.HasPrefix(href, "mailto:") {
			// if mailto put in emails
			email, _ := strings.CutPrefix(href, "mailto:")

			if !slices.Contains(emails, email) {
				emails = append(emails, email)
			}
		} else if strings.HasPrefix(href, site) || strings.HasPrefix(href, "/") {
			// TODO this is for later when doing site crawling if no sitemap.xml
			// put into sites to crawl

		}
	})

	f, err := os.Create(strings.ReplaceAll(site, "/", "") + "_emails.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	c.OnHTML("*", func(h *colly.HTMLElement) {
		text := h.Text

		matches := emailRegex.FindAllString(text, -1)

		if len(matches) != 0 {
			for _, match := range matches {
				if !slices.Contains(emails, match) {
					emails = append(emails, match)
				}
			}
		}

	})

	c.OnScraped(func(r *colly.Response) {
		for _, email := range emails {
			f.WriteString(email + "\n")
		}

	})

	c.Visit(site)
	c.Wait()
}
