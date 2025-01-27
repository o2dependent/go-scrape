package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/o2dependent/go-scrape/utils"
)

var tldListURL = "https://data.iana.org/TLD/tlds-alpha-by-domain.txt"
var tldListTempFileName = "scrapper-tld-list"

func failedGetTLDsMessage() {
	fmt.Println("failed to get TLD list")
	fmt.Println("ignoring validating TLD")
}

func getTLDs() []string {
	if !ignoreTLDCache {
		// check temp
		tmpList, rerun := getTempTLD()
		if !rerun && len(tmpList) > 0 {
			return tmpList
		}
	}

	// fetch tld list
	tldList := []string{}

	req, err := http.NewRequest(http.MethodGet, tldListURL, nil)
	if err != nil {
		failedGetTLDsMessage()
		return tldList
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		failedGetTLDsMessage()
		return tldList
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		failedGetTLDsMessage()
		return tldList
	}

	resBodyStr := string(resBody)
	if resBodyStr == "" {
		failedGetTLDsMessage()
		return tldList
	}

	tldList = strings.Split(resBodyStr, "\n")

	// TODO: store in tmp
	createTempTLD(resBody)

	tldList = tldList[1:]

	return tldList
}

func createTempTLD(body []byte) {
	dir := utils.GetTempDir()
	if dir == "" {
		return
	}

	f, err := os.Create(dir + "/" + tldListTempFileName)
	if err != nil {
		return
	}
	defer f.Close()
	f.Write(body)
}

func getTempTLD() (tldList []string, rerun bool) {
	rerun = true
	dir := utils.GetTempDir()
	if dir == "" {
		return tldList, rerun
	}

	f, err := os.ReadFile(dir + "/" + tldListTempFileName)
	if err != nil {
		return tldList, rerun
	}
	contents := string(f)
	tldList = strings.Split(contents, "\n")

	dateStr := strings.Replace(strings.Split(tldList[0], ",")[1], " Last Updated ", "", -1)
	// dateStr := "Mon Jan 1 15:04:05 2025 MST"
	layout := "Mon Jan 1 15:04:05 2006 MST"
	date, err := time.Parse(layout, dateStr)
	if err == nil && date.Add(time.Hour*27*7).Compare(time.Now()) == -1 {
		return tldList, rerun
	}
	rerun = false

	return tldList[1:], rerun
}
