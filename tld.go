package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

var tldListURL = "https://data.iana.org/TLD/tlds-alpha-by-domain.txt"

func failedGetTLDsMessage() {
	fmt.Println("failed to get TLD list")
	fmt.Println("ignoring validating TLD")
}

func getTLDs() []string {
	tldList := []string{}

	// TODO: check tmp

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
	tldList = tldList[1:]

	return tldList
}
