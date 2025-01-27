package main

import (
	"slices"
	"strings"

	"github.com/o2dependent/go-scrape/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "scrape-emails",
	Short: "CLI to scrape emails from websites",
	Long:  "CLI utility to scrape emails from provided websites",
	Run: func(cmd *cobra.Command, args []string) {

		// validate cli flags
		if string(output[len(output)-1:]) != "/" {
			output = output + "/"
		}

		f := createOutputFile()
		defer f.Close()

		emails := scrape(url)

		if validateTLD {
			tlds := getTLDs()
			emails = utils.Filter(emails, func(s string) bool {
				split := strings.Split(s, ".")
				tld := strings.ToUpper(split[len(split)-1])
				return slices.Contains(tlds, tld)
			})
		}

		generateOutput(f, emails)
	},
}
