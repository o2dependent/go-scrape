package main

import (
	"errors"
	"fmt"
	URL "net/url"
	"slices"
	"strings"

	"github.com/o2dependent/go-scrape/logger"
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
			_, err := URL.Parse(a)
			if err != nil {
				logger.Err.Println(a + " is an invalid website and will be skipped")
			}
			return err == nil
		})
		if len(urls) < 1 {
			return errors.New("requires at least one valid website")
		}

		if len(urls) < len(args) {
			logger.Warn.Println("\"http://\" or \"https://\" is required before website url to work")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// validate cli flags
		// output
		if string(output[len(output)-1:]) != "/" {
			output = output + "/"
		}

		// outputTypes
		fileType = strings.ToLower(fileType)
		if !slices.Contains(outputTypes, fileType) {
			fmt.Println(fileType + " is not a valid file type")
			fmt.Println("defaulting to " + outputTypes[0])
			fileType = outputTypes[0]
		}

		for _, url := range urls {
			handleScrape(url, 1, []string{})
		}
	},
}
