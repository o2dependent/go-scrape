package main

import (
	"strings"
)

var (
	output              string
	useJS               bool
	validateTLD         bool
	timeout             int
	ignoreTLDCache      bool
	collectPhoneNumbers bool
	fileType            string
)

func init() {
	rootCmd.PersistentFlags().BoolVar(&ignoreTLDCache, "ignore-tld-cache", false, "Ignore TLD list cache and refetch TLD list")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "./", "Output folder")
	rootCmd.PersistentFlags().StringVarP(&fileType, "file-type", "f", outputTypes[0], "Output file type ("+strings.Join(outputTypes[:], ", ")+")")
	rootCmd.PersistentFlags().BoolVarP(&useJS, "use-js", "j", false, "Enable Javascript when scraping")
	rootCmd.PersistentFlags().BoolVarP(&collectPhoneNumbers, "phone", "p", false, "Collect phone numbers when scraping")
	rootCmd.PersistentFlags().BoolVarP(&validateTLD, "validate-tld", "v", true, "Validate email address TLD")
	rootCmd.PersistentFlags().IntVarP(&timeout, "timeout", "t", 0, "Page timeout when Javascript Enabled")
}

// add phone numbers scraping
