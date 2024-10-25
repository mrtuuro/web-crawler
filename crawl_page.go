package main

import (
    "fmt"
    "net/url"
    "sync"
)

type config struct {
    maxPages           int
    pages              map[string]int
    baseURL            *url.URL
    mu                 *sync.Mutex
    concurrencyControl chan struct{}
    wg                 *sync.WaitGroup
}

func NewConfig(maxPages int, strURL string, maxConcurrency int) *config {
    baseURL, err := url.Parse(strURL)
    if err != nil {
        fmt.Printf("parsing url: %v\n", err.Error())
        baseURL = &url.URL{}
    }
    return &config{
        maxPages:           maxPages,
        pages:              make(map[string]int),
        baseURL:            baseURL,
        mu:                 &sync.Mutex{},
        concurrencyControl: make(chan struct{}, maxConcurrency),
        wg:                 &sync.WaitGroup{},
    }
}

func (c *config) crawlPage(rawCurrentURL string) {

    c.concurrencyControl <- struct{}{}
    defer func() {
        <-c.concurrencyControl
        c.wg.Done()
    }()

    c.mu.Lock()
    if len(c.pages) >= c.maxPages {
        c.mu.Unlock()
        return
    }
    c.mu.Unlock()

    currentURL, err := url.Parse(rawCurrentURL)
    if err != nil {
        fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
        return
    }

    // skip other websites
    if currentURL.Hostname() != c.baseURL.Hostname() {
        return
    }

    normalizedURL, err := normalizeURL(rawCurrentURL)
    if err != nil {
        fmt.Printf("Error - normalizedURL: %v", err)
        return
    }

    isFirst := c.addPageVisit(normalizedURL)
    if !isFirst {
        return
    }

    fmt.Printf("crawling %s\n", rawCurrentURL)

    htmlBody, err := getHTML(rawCurrentURL)
    if err != nil {
        fmt.Printf("Error - getHTML: %v", err)
        return
    }

    nextURLs, err := getURLsFromHTML(htmlBody, c.baseURL)
    if err != nil {
        fmt.Printf("Error - getURLsFromHTML: %v", err)
        return
    }

    for _, nextURL := range nextURLs {
        c.wg.Add(1)
        go c.crawlPage(nextURL)
    }

}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
    cfg.mu.Lock()
    defer cfg.mu.Unlock()

    if _, visited := cfg.pages[normalizedURL]; visited {
        cfg.pages[normalizedURL]++
        return false
    }

    cfg.pages[normalizedURL] = 1
    return true

}
