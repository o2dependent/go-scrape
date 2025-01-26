package main

import (
	"regexp"
	"slices"
	"strings"

	"github.com/gocolly/colly/v2"
)

func scrape(site string, emailRegex *regexp.Regexp) []string {
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
		}
	})

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

	// c.OnScraped(func(r *colly.Response) {
	// 	for _, email := range emails {
	// 		f.WriteString(email + "\n")
	// 	}
	// })

	c.Visit(site)
	c.Wait()

	return emails
}
