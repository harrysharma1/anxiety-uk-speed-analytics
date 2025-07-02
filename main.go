package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
)

func read_urls() []string {
	file, err := os.ReadFile("urllist.txt")
	if err != nil {
		panic(err)
	}

	urlList := strings.Split(string(file), "\n")

	return urlList
}

func analyse_page() {
	const endpoint = `https://www.googleapis.com/pagespeedonline/v5/runPagespeed`
	const key = `DO NOT COMMIT KEY`
	urlList := read_urls()

	for i := range urlList {
		fmt.Printf("Running analysis on url: %s...\n", urlList[i])
		fullEndpoint := fmt.Sprintf("%s?url=%s&key=%s", endpoint, urlList[i], key)
		resp, err := http.Get(fullEndpoint)
		if err != nil {
			panic(err)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))
	}
}

func main() {
	fmt.Printf("OS: %s\nArchitecture: %s\n", runtime.GOOS, runtime.GOARCH)
	analyse_page()
}
