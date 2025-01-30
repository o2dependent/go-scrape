package goscrape

import (
	"strings"
)

var (
	output                string
	useJS                 bool
	validateTLD           bool
	timeout               int
	ignoreTLDCache        bool
	collectPhoneNumbers   bool
	fileType              string
	maxDepth              int
	consolidateDepthFiles bool
)

func init() {
	RootCmd.PersistentFlags().BoolVar(&ignoreTLDCache, "ignore-tld-cache", false, "Ignore TLD list cache and refetch TLD list")
	RootCmd.PersistentFlags().BoolVar(&consolidateDepthFiles, "consolidate-depth-files", true, "Combine depth results to parent url output")
	RootCmd.PersistentFlags().StringVarP(&output, "output", "o", "./", "Output folder")
	RootCmd.PersistentFlags().StringVarP(&fileType, "file-type", "f", outputTypes[0], "Output file type ("+strings.Join(outputTypes[:], ", ")+")")
	RootCmd.PersistentFlags().BoolVarP(&useJS, "use-js", "j", false, "Enable Javascript when scraping")
	RootCmd.PersistentFlags().BoolVarP(&collectPhoneNumbers, "phone", "p", false, "Collect phone numbers when scraping")
	RootCmd.PersistentFlags().BoolVarP(&validateTLD, "validate-tld", "v", true, "Validate email address TLD")
	RootCmd.PersistentFlags().IntVarP(&timeout, "timeout", "t", 0, "Page timeout when Javascript Enabled")
	RootCmd.PersistentFlags().IntVarP(&maxDepth, "max-depth", "m", 1, "Max recursive depth page scraping")
}

// add phone numbers scraping
