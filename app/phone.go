package goscrape

import (
	"regexp"
	"slices"

	"github.com/nyaruka/phonenumbers"
)

func extractAndFormatPhoneNumbers(s string) []string {
	re := regexp.MustCompile(`(\+?\d[\d-. ()]*|\(\d{3}\)[\d-. ]*|\d[\d-. ]*){7,}`)
	matches := re.FindAllString(s, -1)
	var results []string

	for _, match := range matches {
		num, err := phonenumbers.Parse(match, "US")
		if err != nil || !phonenumbers.IsValidNumber(num) {
			continue
		}
		formatted := phonenumbers.Format(num, phonenumbers.NATIONAL)
		// cleaned := cleanNumber(match)
		// formatted := formatNumber(cleaned)
		if formatted != "" && !slices.Contains(results, formatted) {
			results = append(results, formatted)
		}
	}
	return results
}
