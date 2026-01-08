package main

import (
	"log"
	"net/url"
)


func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.conControl <- struct{}{}
	defer func() {
		<-cfg.conControl
	}()
	
	var pageData PageData
	cfg.isPageLimit = cfg.maxPageCheck()

	if cfg.isPageLimit {
		log.Println("Max pages reached, returning...")
		return
	} else {
		log.Printf("Parsing urls %v...\n", rawCurrentURL)
		currentUrl, err := url.Parse(rawCurrentURL)
		if err != nil {
			log.Printf("Error parsing current url: %v\n", err)
			return
		}

		if currentUrl.Hostname() != cfg.baseUrl.Hostname()	 {
			log.Printf("Invalid url or Url out of scope: %v\n", rawCurrentURL)
			return
		}

		log.Println("Normalizing url...")
		normalized, err := normalizeURL(rawCurrentURL)
		if err != nil {
			log.Printf("Error normalizing url: %v\n", err)
			return
		}
		
	
		log.Printf("Done! Crawling %v...\n", normalized)
		// NOTE: You are passing rawCurrentURL here. This treats variations of the same URL as unique pages, causing duplicates and filling up the maxPages limit with redundant content. Use the normalized URL here.
		ok := cfg.addPageVisit(normalized)
		if !ok {
			log.Printf("Already visited %v\n", normalized)
			return
		}

		log.Printf("Extracting html from %v...\n", normalized)
		html, err := getHTML(rawCurrentURL)
		if err != nil {
			log.Printf("Could not extract html: %v\n", err)
			return
		}

		log.Printf("Done! Extracting Data from %v...\n", normalized)
		pageData = extractPageData(html, rawCurrentURL)
		log.Println("Adding page data...")
		// NOTE: Ensure you use the normalized URL here as well so the map key matches the one used in addPageVisit.
		cfg.setPageData(normalized, pageData)
	}

	log.Println("Going to next url...")
	for _, nextUrl := range pageData.OutgoingLinks {
		cfg.wg.Go(func() {
			if !cfg.isPageLimit {
				cfg.safeCrawl(nextUrl)
			}
		})
	}
}

func (cfg *config) safeCrawl(url string) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
		}
	}()
	cfg.crawlPage(url)
}