package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
)

const FILENAME = "pages_speed.csv"

func isValidURL(url string) bool {
	var pattern = `^(http|https)://[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}(/[a-zA-Z0-9-._~:?#@!$&'()*+,;=]*)*$`
	urlRegex := regexp.MustCompile(pattern)
	return urlRegex.MatchString(url)
}

func readUrls() []string {
	file, err := os.ReadFile("urllist.txt")
	if err != nil {
		panic(err)
	}

	urlList := strings.Split(string(file), "\n")

	return urlList
}

func pPrint(responseStruct PageSpeedResponse) {
	if responseStruct.LoadingExperience.OverallCategory == "SLOW" || responseStruct.LoadingExperience.OverallCategory == "AVERAGE" {
		fmt.Println("Core Web Vitals: Failed")
	} else {
		fmt.Println("Core Web Vitals: Passed")
	}
	fmt.Println("Metrics:")
	fmt.Printf("- First Contentful Paint (FCP): %s (%s)\n", responseStruct.LighthouseResult.Audits["first-contentful-paint"].DisplayValue, responseStruct.LoadingExperience.Metrics["FIRST_CONTENTFUL_PAINT_MS"].Category)
	fmt.Printf("- Largest Contentful Paint (LCP): %s (%s)\n", responseStruct.LighthouseResult.Audits["largest-contentful-paint"].DisplayValue, responseStruct.LoadingExperience.Metrics["LARGEST_CONTENTFUL_PAINT_MS"].Category)
	fmt.Printf("- Cumulative Layout Shift (CLS): %s (%s)\n", responseStruct.LighthouseResult.Audits["cumulative-layout-shift"].DisplayValue, responseStruct.LoadingExperience.Metrics["CUMULATIVE_LAYOUT_SHIFT_SCORE"].Category)
	fmt.Printf("- Total Blocking Time (TBT): %s\n", responseStruct.LighthouseResult.Audits["total-blocking-time"].DisplayValue)
	fmt.Printf("- Speed Index: %s\n", responseStruct.LighthouseResult.Audits["speed-index"].DisplayValue)
	fmt.Printf("Time Stamp (UTC): %s\n", responseStruct.AnalysisUTCTimeStamp)
}

func analysePages() {
	urlList := readUrls()

	for i := range urlList {
		analysePage(i)
	}
}

func analysePage(i int) {
	const endpoint = `https://www.googleapis.com/pagespeedonline/v5/runPagespeed`
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	key := os.Getenv("API_KEY")

	urlList := readUrls()

	fmt.Printf("(%d/%d) Running analysis on url: \"%s\"...\n", i+1, len(urlList), urlList[i])
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

	pPrint(responseStruct)
	// store_csv(responseStruct)
}

func initCsv() {
	file, err := os.Create(FILENAME)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{
		"URL",
		"Core Vitals Assessment",
		"First Contentful Paint (FCP)",
		"Largest Contentful Paint (LCP)",
		"Cumulative Layout Shift (CLS)",
		"Total Blocking Time (TBT)",
		"Speed Index",
		"Timestamp (UTC)",
	}

	errHeader := writer.Write(headers)
	if errHeader != nil {
		panic(errHeader)
	}
}

func storeCsv(responseStruct PageSpeedResponse) {
	file, err := os.OpenFile(FILENAME, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{
		responseStruct.ID,
		responseStruct.LoadingExperience.OverallCategory,
		fmt.Sprintf("%s (%s)", responseStruct.LighthouseResult.Audits["first-contentful-paint"].DisplayValue, responseStruct.LoadingExperience.Metrics["FIRST_CONTENTFUL_PAINT_MS"].Category),
		fmt.Sprintf("%s (%s)", responseStruct.LighthouseResult.Audits["largest-contentful-paint"].DisplayValue, responseStruct.LoadingExperience.Metrics["LARGEST_CONTENTFUL_PAINT_MS"].Category),
		fmt.Sprintf("%s (%s)", responseStruct.LighthouseResult.Audits["cumulative-layout-shift"].DisplayValue, responseStruct.LoadingExperience.Metrics["CUMULATIVE_LAYOUT_SHIFT_SCORE"].Category),
		responseStruct.LighthouseResult.Audits["total-blocking-time"].DisplayValue,
		responseStruct.LighthouseResult.Audits["speed-index"].DisplayValue,
		responseStruct.AnalysisUTCTimeStamp,
	}

	errRecord := writer.Write(record)
	if errRecord != nil {
		panic(errRecord)
	}

	fmt.Printf("✅ CSV record written successfully for %s\n\n", responseStruct.ID)

}
