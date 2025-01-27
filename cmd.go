package main

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/o2dependent/go-scrape/utils"
	"github.com/spf13/cobra"
)

var urls []string

var rootCmd = &cobra.Command{
	Use:   "go-scrape [...websites]",
	Short: "CLI to scrape emails from websites",
	Long:  "CLI utility to scrape emails from provided websites",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one website")
		}
		urls = utils.Filter(args, func(a string) bool {
			valid := utils.WebsiteRegex.MatchString(a)
			if !valid {
				fmt.Println(a + " is an invalid website and will be skipped")
			}
			return valid
		})
		if len(urls) < 1 {
			return errors.New("requires at least one valid website")
		}
		if len(urls) < len(args) {
			fmt.Println("\"http://\" or \"https://\" is required before website url to work")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// validate cli flags
		if string(output[len(output)-1:]) != "/" {
			output = output + "/"
		}

		for _, url := range urls {
			f := createOutputFile(url)
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
		}
	},
}
