package main

var (
	url    string
	output string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&url, "url", "u", "https://eolsen.dev", "URL to scrape")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "./", "Output folder")
}

// add phone numbers scraping
