package main

import (
    "fmt"
    "sync"
)

type ReportElement struct {
    pageCount int
    url       string
}

func printReport(pages map[string]int, baseURL string) {
    wg := &sync.WaitGroup{}
    fmt.Printf("\n=============================\nREPORT for %v\n=============================\n", baseURL)

    var reportElements []ReportElement
    for u, c := range pages {
        var reportElement ReportElement
        reportElement.url = u
        reportElement.pageCount = c
        reportElements = append(reportElements, reportElement)
    }


    wg.Add(1)
    go func() {
        defer wg.Done()
        reportElements = sortPages(reportElements)
    } ()
    wg.Wait()

    for _, e := range reportElements {
        fmt.Printf("Found %d internal links to %s\n", e.pageCount, e.url)
    }

}

func sortPages(reportElements []ReportElement) []ReportElement {
    var isDone = false
    for !isDone {
        isDone = true
        var i = 0
        for i < len(reportElements) - 1 {
            if reportElements[i].pageCount < reportElements[i+1].pageCount {
                reportElements[i], reportElements[i+1] = reportElements[i+1], reportElements[i]
                isDone = false
            }
            i++
        }
    }

    return reportElements
}
