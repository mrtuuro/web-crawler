package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    args := os.Args

    if len(args) < 4 {
        fmt.Println("Usage: ./crawler <website> <maxConcurrency> <maxPages>")
        os.Exit(1)
    } else if len(args) > 4 {
        fmt.Println("too many arguments provided")
        os.Exit(1)
    }

    baseURL := args[1]
    maxConcurrency, err := strconv.Atoi(args[2])
    if err != nil {
        fmt.Printf("converting string to int: %v\n", err.Error())
    }

    maxPages, err := strconv.Atoi(args[3])
    if err != nil {
        fmt.Printf("converting string to int: %v\n", err.Error())
    }

    cfg := NewConfig(maxPages, baseURL, maxConcurrency)

    fmt.Printf("starting crawl of: %s...\n", baseURL)

    cfg.wg.Add(1)
    go cfg.crawlPage(baseURL)
    cfg.wg.Wait()

    printReport(cfg.pages, baseURL)

}
