package main

var (
	url         string
	output      string
	useJS       bool
	validateTLD bool
	timeout     int
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&url, "url", "u", "https://eolsen.dev", "URL to scrape")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "./", "Output folder")
	rootCmd.PersistentFlags().BoolVarP(&useJS, "use-js", "j", false, "Enable Javascript when scraping")
	rootCmd.PersistentFlags().BoolVarP(&validateTLD, "validate-tld", "v", false, "Validate email address TLD")
	rootCmd.PersistentFlags().IntVarP(&timeout, "timeout", "t", 0, "Page timeout when Javascript Enabled")
}

// add phone numbers scraping
