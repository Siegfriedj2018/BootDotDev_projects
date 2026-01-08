package main

import (
	"net/url"
	"sync"
)

type config struct {
	pages					map[string]PageData
	baseUrl				*url.URL
	mu 						*sync.Mutex
	conControl		chan struct{}
	wg						*sync.WaitGroup
	maxPages			int
	isPageLimit		bool
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	_, isFirst = cfg.pages[normalizedURL]
	if isFirst {
		return false
	}

	if len(cfg.pages) >= cfg.maxPages {
		return false
	}

	cfg.pages[normalizedURL] = PageData{URL: normalizedURL}
	return true
}

func (cfg *config) setPageData(normalizedURL string, data PageData) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	
	cfg.pages[normalizedURL] = data
}

func (cfg *config) maxPageCheck() bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if len(cfg.pages) >= cfg.maxPages {
		return true
	}

	return false
}