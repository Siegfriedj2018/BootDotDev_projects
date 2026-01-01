package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type PageData struct {
	URL 						string
	H1							string
	FirstParagraph 	string
	OutgoingLinks		[]string
	ImageURLs				[]string
}

func getH1FromHTML(html string) string {
	if html == "" {
		return ""
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}

	return doc.Find("h1").First().Text()
}

func getFirstParagraphFromHTML(html string) string {
	if html == "" {
		return ""
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return "Error, parsing html"
	}

	foundText := doc.Find("main p").First()

	if foundText.Length() == 0 {
		foundText = doc.Find("p").First()
	}
	return foundText.Text()
}

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	foundUrls := make([]string, 0)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, fmt.Errorf("Cant parse html, %w\n", err)
	}

	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
    // For each '<a href>' it finds, it will run this function.
		href, ok := s.Attr("href")
		if !ok {
			return
		}

		newUrl := strings.TrimSpace(href)
		if href == "" {
			return
		}


		parsedUrl, err := url.Parse(newUrl)
		if err != nil {
			return
		}

		if !parsedUrl.IsAbs() {
			parsedUrl = baseURL.JoinPath(parsedUrl.String())
		}

		foundUrls = append(foundUrls, parsedUrl.String())
  })

	return foundUrls, nil
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	foundImgs := make([]string, 0)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, fmt.Errorf("Cant parse html, %w\n", err)
	}

	doc.Find("img").Each(func(_ int, s *goquery.Selection) {
    // For each '<img src>' it finds, it will run this function.
		href, ok := s.Attr("src")
		if !ok {
			return
		}

		newUrl := strings.TrimSpace(href)
		if href == "" {
			return
		}
		parsedUrl, err := url.Parse(newUrl)
		if err != nil {
			return
		}
		
		if !parsedUrl.IsAbs() {
			parsedUrl = baseURL.JoinPath(parsedUrl.String())
		}

		foundImgs = append(foundImgs, parsedUrl.String())
  })

	return foundImgs, nil
}

func extractPageData(html, pageURL string) PageData {
	if html == "" {
		log.Println("Invalid html")
		return PageData{}
	}

	currentUrl, err := url.Parse(pageURL)
	if err != nil {
		log.Printf("Invalid Url, %v", err)
		return PageData{}
	}

	foundH1 := getH1FromHTML(html)
	firstPara := getFirstParagraphFromHTML(html)
	outLinks, err := getURLsFromHTML(html, currentUrl)
	if err != nil {
		log.Fatalf("Error parsing html links, %v\n", err)
	}

	imageLinks, err := getImagesFromHTML(html, currentUrl)
	if err != nil {
		log.Fatalf("Error parsing image links, %v\n", err)
	}

	return PageData{
		URL: 						currentUrl.String(),
		H1: 						foundH1,
		FirstParagraph: firstPara,
		OutgoingLinks: 	outLinks,
		ImageURLs: 			imageLinks,
	}
}