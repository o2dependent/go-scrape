package main

import (
	"log"
	"os"

	goscrape "github.com/o2dependent/goscrape/app"
)

// INIT IDEA
// Scrape all emails on a site
// Maybe hit the sitemap.xml first and if that fail just hit the page and crawl links
// Start with just nabbing the emails from on page

func main() {
	if err := goscrape.RootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
