package main

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/o2dependent/go-scrape/utils"
)

func scrape(site string) ([]string, []string) {
	c := colly.NewCollector(colly.IgnoreRobotsTxt())
	c.Async = true

	emails := []string{}
	numbers := []string{}

	if useJS {
		c.OnResponse(func(r *colly.Response) {
			if err := initWithJavascript(r); err != nil {
				log.Println(err)
				return
			}
		})
	}

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

		matches := utils.EmailRegex.FindAllString(text, -1)

		if len(matches) != 0 {
			for _, match := range matches {
				if !slices.Contains(emails, match) {
					emails = append(emails, match)
				}
			}
		}
		if collectPhoneNumbers {
			extractedNumbers := extractAndFormatPhoneNumbers(text)
			if len(extractedNumbers) > 0 {
				for _, num := range extractedNumbers {
					if !slices.Contains(numbers, num) {
						numbers = append(numbers, num)
					}
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

	for _, num := range numbers {
		fmt.Println(num)
	}

	return emails, numbers
}
