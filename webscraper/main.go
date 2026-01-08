package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"sync"
	"strconv"
)

func main() {
	if len(os.Args[1:]) < 1 {
		log.Fatalln("no website provided")
	} else if len(os.Args[1:]) > 4 {
		log.Fatalln("too many arguements")
	}

	pd := make(map[string]PageData)
	baseURL := os.Args[1]
	parseBaseUrl, err := url.Parse(baseURL)
	if err != nil {
		log.Printf("Error parsing base url: %v", err)
		return
	}

	numWorkers, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalln("Invalid number entered for max concurrency")
	}

	maxPage, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatalln("Invalid number entered for max pages")
	}

	isPageLimit := false

	cfg := config{
		pages:      	pd,
		baseUrl:    	parseBaseUrl,
		mu:         	&sync.Mutex{},
		conControl: 	make(chan struct{}, numWorkers),
		wg:         	&sync.WaitGroup{},
		maxPages: 		maxPage,
		isPageLimit: 	isPageLimit,
	}

	fmt.Printf("starting crawl of: %s\n", baseURL)

	// Using new 1.25 waitgroup.go letting go manage goroutines
	// wg.done() is called automatically when the func() returns
	for range numWorkers {
		cfg.wg.Go(func() {
			if !isPageLimit {
				cfg.safeCrawl(parseBaseUrl.String())
			}
		})
	}

	cfg.wg.Wait()

	fmt.Printf("Crawled: %d pages...\n", len(cfg.pages))
	for page := range cfg.pages {
		fmt.Printf("Pages Crawled: %s\n", page)
	}

	err = writeCSVReport(cfg.pages, "report.csv")
	if err != nil {
		log.Printf("Could save page data: %v", err)
	}
}
