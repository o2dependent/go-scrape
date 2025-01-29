package main

import (
	"errors"
	"os"
	"strings"

	"github.com/o2dependent/go-scrape/logger"
	"github.com/o2dependent/go-scrape/utils"
)

var outputTypes = []string{"txt", "csv", "json", "yaml"}

func createOutputFile(url string) *os.File {
	directoryValid, err := utils.DirectoryExists(output)
	if !directoryValid || err != nil {
		logger.Err.Println(errors.New("\"" + output + "\" is an invalid directory"))
		os.Exit(1)
	}

	f, err := os.Create(output + strings.ReplaceAll(url, "/", "") + "." + fileType)
	if err != nil {
		logger.Err.Println("failed to create output file")
		os.Exit(1)
	}

	return f
}

func generateOutput(f *os.File, emails []string, numbers []string) {
	logger.Info.Println("saving data to " + f.Name())
	if fileType == "txt" {
		f.WriteString("Emails\n")
		f.WriteString(strings.Join(emails, "\n"))

		if collectPhoneNumbers {
			f.WriteString("\n\nPhone\n")
			f.WriteString(strings.Join(numbers, "\n"))
		}
	} else if fileType == "csv" {
		contents := ""
		if collectPhoneNumbers {
			contents += "Email;Phone\n"
		} else {
			contents += "Email\n"
		}

		maxLen := max(len(numbers), len(emails))
		for i := 0; i < maxLen; i++ {
			content := ""
			if i < len(emails) {
				content += emails[i]
			}
			if collectPhoneNumbers {
				content += ";"
				if i < len(numbers) {
					content += numbers[i]
				}
			}
			content += "\n"
			contents += content
		}

		f.WriteString(contents)
	} else if fileType == "json" {
		contents := "{"
		contents += "\"Emails\": ["
		for i, email := range emails {
			contents += "\"" + email + "\""
			if i < len(emails)-1 {
				contents += ","
			}
		}
		contents += "]"
		if collectPhoneNumbers {
			contents += ",\"Phones\": ["
			for i, number := range numbers {
				contents += "\"" + number + "\""
				if i < len(numbers)-1 {
					contents += ","
				}
			}
			contents += "]"
		}
		contents += "}"
		f.WriteString(contents)
	} else if fileType == "yaml" {
		contents := ""
		contents += "Emails:\n"
		for _, email := range emails {
			contents += "  - " + email + "\n"
		}

		if collectPhoneNumbers {
			contents += "Phones:\n"
			for _, number := range numbers {
				contents += "  - " + number + "\n"
			}
		}
		f.WriteString(contents)
	}
}
