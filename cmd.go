package main

import (
	"errors"
	"os"
	"regexp"
	"strings"

	"github.com/o2dependent/go-scrape/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "scrape-emails",
	Short: "CLI to scrape emails from websites",
	Long:  "CLI utility to scrape emails from provided websites",
	Run: func(cmd *cobra.Command, args []string) {
		site := "https://eolsen.dev"
		writeDir := "output/"

		directoryValid, err := utils.DirectoryExists(writeDir)
		if !directoryValid || err != nil {
			panic(errors.New("directory is invalid"))
		}
		f, err := os.Create(writeDir + strings.ReplaceAll(site, "/", "") + "_emails.txt")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

		emails := scrape(site, emailRegex)

		for _, email := range emails {
			f.WriteString(email + "\n")
		}
	},
}
