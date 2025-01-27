package main

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/o2dependent/go-scrape/utils"
)

func createOutputFile() *os.File {
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

	return f
}

func generateOutput(f *os.File, emails []string) {
	for _, email := range emails {
		f.WriteString(email + "\n")
	}
}
