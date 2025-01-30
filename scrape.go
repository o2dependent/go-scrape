package main

import (
	URL "net/url"
	"os"
	"slices"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/o2dependent/go-scrape/logger"
	"github.com/o2dependent/go-scrape/utils"
)

func scrape(url string, visitedUrls []string) (emails []string, numbers []string, additionalUrls []string) {
	logger.InfoAccent.Println("scraping: " + url)

	visitedUrls = append(visitedUrls, url)

	// filter out path and find base url
	parsedUrl, err := URL.Parse(url)
	if err != nil {
		return
	}
	baseUrl := url
	if parsedUrl.Path == "/" {
		baseUrl = url[:len(url)-1]
	} else if parsedUrl.Path != "" {
		baseUrl = strings.Split(url, parsedUrl.Path)[0]
	}

	c := colly.NewCollector(colly.IgnoreRobotsTxt())
	c.Async = true

	if useJS {
		c.OnResponse(func(r *colly.Response) {
			if err := initWithJavascript(r); err != nil {
				logger.Err.Println(err)
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
		} else if string(href[0]) == "/" && !slices.Contains(slices.Concat(visitedUrls, additionalUrls), baseUrl+href) {
			additionalUrls = append(additionalUrls, baseUrl+href)
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

	c.Visit(url)
	c.Wait()

	return
}

func handleScrape(url string, depth int, visitedUrls []string) ([]string, []string) {
	var f *os.File
	if depth == 1 || !consolidateDepthFiles {
		f = createOutputFile(url)
		defer f.Close()
	}

	emails, numbers, additionalUrls := scrape(url, visitedUrls)

	visitedUrls = append(visitedUrls, url)

	if validateTLD && len(emails) > 0 {
		logger.Info.Println("validating emails by TLD")
		tlds := getTLDs()
		emails = utils.Filter(emails, func(s string) bool {
			split := strings.Split(s, ".")
			tld := strings.ToUpper(split[len(split)-1])
			return slices.Contains(tlds, tld)
		})
	}

	logger.Info.Printf("found %v emails\n", len(emails))
	if collectPhoneNumbers {
		logger.Info.Printf("found %v phone numbers\n", len(numbers))
	}

	// here so that if the process is canceled while not consolidating the first url is saved
	if !consolidateDepthFiles {
		generateOutput(f, emails, numbers)
	}

	if depth < maxDepth {
		if len(additionalUrls) == 0 {
			logger.Warn.Println("no urls found on page")
			return emails, numbers
		}
		logger.Warn.Printf("scraping depth: %v\n", depth+1)
		for _, additionalUrl := range additionalUrls {
			dEmails, dNumbers := handleScrape(additionalUrl, depth+1, visitedUrls)
			if consolidateDepthFiles {
				emails = append(emails, dEmails...)
				numbers = append(numbers, dNumbers...)
			}
		}
	}

	if depth == 1 && consolidateDepthFiles {
		generateOutput(f, emails, numbers)
	}

	return emails, numbers
}
