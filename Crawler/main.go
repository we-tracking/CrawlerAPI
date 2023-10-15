package main

import (
	"Crawler/Utils"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/api/product", apiProductHandler)
	http.HandleFunc("/api/search", apiSearchHandler)
	fmt.Println("Server running")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func apiSearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Body == nil {
		http.Error(w, "Body cannot be blank", http.StatusBadRequest)
		return
	}

	var searchJob Utils.SearchJob
	Utils.GetSearchJob(w, r.Body, &searchJob)

	var normalizedProducts []Utils.NormalizedProduct

	// gets the html
	html := getHtml(searchJob.Url)

	// the Tags for the product
	tags := Utils.Tags{}
	Utils.GetTagFile(&tags, 1, "search")

	// Extract raw data into "products" using "tags"
	rawProducts := tags.ExtractSearch(html)

	normalizedProduct := tags.Normalize(rawProducts)

	normalizedProducts = append(normalizedProducts, normalizedProduct...)

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(normalizedProducts)
	if err != nil {
		return
	}
}

func apiProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Body == nil {
		http.Error(w, "Body cannot be blank", http.StatusBadRequest)
		return
	}

	// transforms the input into JOB instance
	// TODO: validate input
	var productJob Utils.ProductJob
	Utils.GetProductJob(w, r.Body, &productJob)

	var normalizedProducts []Utils.NormalizedProduct

	// Iterate in URL -> Iterate in Products
	for _, url := range productJob.Urls {
		fmt.Println(url)

		// gets the html
		html := getHtml(url)

		// the Tags for the product
		tags := Utils.Tags{}
		Utils.GetTagFile(&tags, 1, "product")

		// Extract raw data into "products" using "tags"
		rawProducts := tags.ExtractProduct(html)

		normalizedProduct := tags.Normalize(rawProducts)

		normalizedProducts = append(normalizedProducts, normalizedProduct...)

	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(normalizedProducts)
	if err != nil {
		return
	}

	//// default message for "successful request"
	//// response := map[string]string{"message": "JSON received successfully"}

}

func getHtml(url string) *goquery.Document {
	client := &http.Client{
		Timeout: 6 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err, "The newRequest has finished in an error")
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("HTTP call failed:", err)
		return nil
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	html, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalln(err, "The NewDocumentFromReader has finished in an error")
	}
	return html
}
