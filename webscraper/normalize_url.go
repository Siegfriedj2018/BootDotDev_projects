package main

import (
	"fmt"
	"net/url"
	"strings"
	// "strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsedUrl, err := url.Parse(strings.ToLower(rawURL))
	if err != nil {
		return "", fmt.Errorf("Error, invalid url")
	}

	fmt.Printf("Url: Host:%s Path:%s\n", parsedUrl.Host, parsedUrl.Path)
	
	if parsedUrl.Host == "" && parsedUrl.Path != "" {
		reparsedUrl := strings.Split(parsedUrl.Path, "/")
		
		if len(reparsedUrl) == 1 {
			return "", fmt.Errorf("Error, invalid url")
		}
		
		return fmt.Sprintf("%s/%s", reparsedUrl[0], reparsedUrl[1]), nil
	}
	

	normalizedURL := fmt.Sprintf("%s%s", parsedUrl.Host, strings.TrimSuffix(parsedUrl.Path, "/"))
	return normalizedURL, nil
}