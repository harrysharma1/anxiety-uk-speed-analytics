package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func read_urls() []string {
	file, err := os.ReadFile("urllist.txt")
	if err != nil {
		panic(err)
	}

	urlList := strings.Split(string(file), "\n")

	return urlList
}

func p_print(responseStruct PageSpeedResponse) {
	if responseStruct.LoadingExperience.OverallCategory == "SLOW" || responseStruct.LoadingExperience.OverallCategory == "AVERAGE" {
		fmt.Println("Core Web Vitals: Failed")
	} else {
		fmt.Println("Core Web Vitals: Passed")
	}
	fmt.Println("Metrics:")
	fmt.Printf("- First Contentful Paint (FCB): %s (%s)\n", responseStruct.LighthouseResult.Audits["first-contentful-paint"].DisplayValue, responseStruct.LoadingExperience.Metrics["FIRST_CONTENTFUL_PAINT_MS"].Category)
	fmt.Printf("- Largest Contentful Paint (LCP): %s (%s)\n", responseStruct.LighthouseResult.Audits["largest-contentful-paint"].DisplayValue, responseStruct.LoadingExperience.Metrics["LARGEST_CONTENTFUL_PAINT_MS"].Category)
	fmt.Printf("- Cumulative Layout Shift (CLS): %s (%s)\n", responseStruct.LighthouseResult.Audits["cumulative-layout-shift"].DisplayValue, responseStruct.LoadingExperience.Metrics["CUMULATIVE_LAYOUT_SHIFT_SCORE"].Category)
	fmt.Printf("- Total Blocking Time (TBT): %s\n", responseStruct.LighthouseResult.Audits["total-blocking-time"].DisplayValue)
	fmt.Printf("- Speed Index: %s\n", responseStruct.LighthouseResult.Audits["speed-index"].DisplayValue)
	fmt.Printf("Time Stamp (UTC): %s\n\n", responseStruct.AnalysisUTCTimeStamp)
}

func analyse_pages() {
	urlList := read_urls()

	for i := range urlList {
		analyse_page(i)
	}
}

func analyse_page(i int) {
	const endpoint = `https://www.googleapis.com/pagespeedonline/v5/runPagespeed`
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	key := os.Getenv("API_KEY")

	urlList := read_urls()

	fmt.Printf("Running analysis on url: \"%s\"...\n", urlList[i])
	fullEndpoint := fmt.Sprintf("%s?url=%s&key=%s", endpoint, urlList[i], key)
	resp, err := http.Get(fullEndpoint)
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var responseStruct PageSpeedResponse
	errDecode := json.Unmarshal(body, &responseStruct)
	if errDecode != nil {
		panic(errDecode)
	}

	p_print(responseStruct)
}

func main() {
	// fmt.Printf("OS: %s\nArchitecture: %s\n", runtime.GOOS, runtime.GOARCH)
	analyse_pages()
}
