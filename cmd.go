package main

import (
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

		generateOutput(f, emails)
	},
}
