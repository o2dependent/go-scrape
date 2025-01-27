package main

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/o2dependent/go-scrape/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "scrape-emails",
	Short: "CLI to scrape emails from websites",
	Long:  "CLI utility to scrape emails from provided websites",
	Run: func(cmd *cobra.Command, args []string) {

		directoryValid, err := utils.DirectoryExists(output)
		if !directoryValid || err != nil {
			log.Println(errors.New("directory is invalid"))
			os.Exit(1)
		}
		f, err := os.Create(output + strings.ReplaceAll(url, "/", "") + "_emails.txt")
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		defer f.Close()

		emails := scrape(url)

		for _, email := range emails {
			f.WriteString(email + "\n")
		}
	},
}
