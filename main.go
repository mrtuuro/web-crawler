package main

import (
    "fmt"
    "os"
)

func main() {
    args := os.Args

    if len(args) < 2 {
        fmt.Println("no website provided")
        os.Exit(1)
    } else if len(args) > 2 {
        fmt.Println("too many arguments provided")
        os.Exit(1)
    }

    baseURL := args[1]
    fmt.Printf("starting crawl of: %s...\n", baseURL)

    _, err := getHTML(baseURL)
    if err != nil {
        fmt.Printf("getting HTML: %v", err.Error())
        os.Exit(1)
    }
    pages := make(map[string]int)
    crawlPage(baseURL, baseURL, pages)

    for normalizedURL, v := range pages {
        fmt.Printf("%s - %d\n", normalizedURL, v)
    }

}
