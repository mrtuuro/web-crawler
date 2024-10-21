package main

import "testing"

func TestGetUrlsFromHTML(t *testing.T) {
    tests := []struct {
        name      string
        inputURL  string
        inputBody string
        expected  []string
        tagCount  int
    }{
        {
            name:     "absolute and relative URLs",
            inputURL: "https://blog.boot.dev",
            inputBody: `
            <html>
            <body>
            <a href="/path/one">
            <span>Boot.dev</span>
            </a>
            <a href="https://other.com/path/one">
            <span>Boot.dev</span>
            </a>
            </body>
            </html>
            `,
            expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
            tagCount: 2,
        },
    }

    for i, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {

            actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
            if err != nil {
                t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
                return
            }

            if tc.tagCount != len(actual) {
                t.Errorf("Test %v - %s FAIL: expected URLs: %v, actual: %v", i, tc.name, tc.expected, actual)
            }

        })
    }

}
