package main

import (
    "fmt"
    "net/url"
    "strings"

    "golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {

    htmlReader := strings.NewReader(htmlBody)
    doc, err := html.Parse(htmlReader)
    if err != nil {
        return nil, fmt.Errorf("could not parse HTML: %v", err.Error())
    }

    var urls []string
    var traverseNodes func(*html.Node)

    traverseNodes = func(node *html.Node) {
        if node.Type == html.ElementNode && node.Data == "a" {
            for _, anchor := range node.Attr {
                if anchor.Key == "href" {
                    href, err := url.Parse(anchor.Val)
                    if err != nil {
                        continue
                    }

                    resolvedURL := baseURL.ResolveReference(href)
                    urls = append(urls, resolvedURL.String())
                }
            }
        }

        for child := node.FirstChild; child != nil; child = child.NextSibling {
            traverseNodes(child)
        }
    }
    traverseNodes(doc)

    return urls, nil

}
