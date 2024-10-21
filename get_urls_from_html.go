package main

import (
    "fmt"
    "net/url"
    "strings"

    "golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
    var links []string
    htmlReader := strings.NewReader(htmlBody)
    node, err := html.Parse(htmlReader)
    if err != nil {
        return links, fmt.Errorf("parsing htmlbody: %v", err.Error())
    }

    linkMap := make(map[string]string)
    linkMap = traverseNode(linkMap, node, 0)

    for _, v := range linkMap {
        converted, err := convertToAbsPath(v)
        if err != nil {
            return links, fmt.Errorf("converting relative to absolute: %v", err.Error())
        }

        links = append(links, converted)
    }


    return links, nil

}

func convertToAbsPath(rawURL string) (string, error) {
    parsedURL, err := url.Parse(rawURL)
    if err != nil {
        return "", fmt.Errorf("couldn't parse URL: %w", err)
    }

    fmt.Println(parsedURL)
    return parsedURL.Host, nil
}

func traverseNode(linkMap map[string]string,n *html.Node, depth int) map[string]string {
    // indent := strings.Repeat("  ", depth)

    switch n.Type {
    case html.ElementNode:
        for _, attr := range n.Attr {
            linkMap[attr.Val] = attr.Val
        }
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        traverseNode(linkMap, c, depth+1)
    }
    return linkMap
}
