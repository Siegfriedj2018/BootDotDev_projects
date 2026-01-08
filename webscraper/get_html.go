package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	client := &http.Client{}
	
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		log.Fatalf("Get request failed: %v\n", err)
	}

	req.Header.Set("User-Agent", "BootStudentCrawler/1.0")
	
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error getting response, %w\n", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return "", fmt.Errorf("status code was not 200: %v\n", res.StatusCode)
	}
	if !strings.Contains(res.Header.Get("Content-Type"), "text/html") {
		return "", fmt.Errorf("Invalid content-type: %v", res.Header.Get("Content-Type"))
	}

	webpage, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("Could not read webpage: %w\n", err)
	}

	return string(webpage), nil
}
