package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const http_prefix = "http://"

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, http_prefix) {
			url = http_prefix + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}

		if _, err = io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		}

		fmt.Printf("Status Code: %d\n", resp.StatusCode)
	}
}
