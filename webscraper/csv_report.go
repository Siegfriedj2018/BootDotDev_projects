package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func writeCSVReport(pages map[string]PageData, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Could not create file: %w", err)
	}

	defer file.Close()
	wrt := csv.NewWriter(file)
	defer wrt.Flush()

	err = wrt.Write([]string{"page_url", "h1", "first_paragraph", "outgoing_link_urls", "image_urls"})
	if err != nil {
		return fmt.Errorf("Could not write header: %w", err)
	}
	for url, pageInfo := range pages {
		err = wrt.Write([]string{url, pageInfo.H1, pageInfo.FirstParagraph, strings.Join(pageInfo.OutgoingLinks,";"), strings.Join(pageInfo.ImageURLs,";")})
		if err != nil {
			return fmt.Errorf("could not write row to file: %w", err)
		}
	}

	return nil
}