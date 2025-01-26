package main

import (
	"log"
	"os"
)

// INIT IDEA
// Scrape all emails on a site
// Maybe hit the sitemap.xml first and if that fail just hit the page and crawl links
// Start with just nabbing the emails from on page

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
