package Utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func GetProductJob(w http.ResponseWriter, r io.ReadCloser, job *ProductJob) {
	err := json.NewDecoder(r).Decode(job)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func GetSearchJob(w http.ResponseWriter, r io.ReadCloser, job *SearchJob) {
	err := json.NewDecoder(r).Decode(job)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// GetProduct TODO: receives (product *Tags) as args dinamically
func GetTagFile(productTag *Tags, retailerID int, pageType string) {
	// gets from cloud or database (for now its hardcoded)

	productTagFile := fmt.Sprintf("Utils/tmp/retailerTags/%d/%s_tags.json", retailerID, pageType)

	if _, err := os.Stat(productTagFile); os.IsNotExist(err) {
		return

	} else {
		tag, err := getFile(productTagFile)
		if err != nil {
			log.Fatal(err)
			return
		} else {
			_ = json.Unmarshal(tag, productTag)
		}
	}
}

func getFile(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return file, err
}
