package Utils

import (
	"regexp"
	"strconv"
	"strings"
)

// Normalize gets a RAW data and returns a NORMALIZED data
func (productTag Tags) Normalize(products []RawProduct) []NormalizedProduct {
	var normalizedProducts []NormalizedProduct

	for _, product := range products {
		normalizedProduct := NormalizedProduct{}

		// TODO: make it less redundant
		normalizedProduct.ProductCode = normalizeString(product.ProductCode)
		normalizedProduct.ProductName = normalizeString(product.ProductName)
		normalizedProduct.ProductURL = normalizeString(product.ProductURL)
		normalizedProduct.Description = normalizeString(product.Description)
		normalizedProduct.Brand = normalizeString(product.Brand)
		normalizedProduct.Seller = normalizeString(product.Seller)

		normalizedProduct.InStock = normalizeAvailability(product.InStock)

		normalizedProduct.FinalPrice = normalizePrice(product.FinalPrice)
		normalizedProduct.FromPrice = normalizePrice(product.FromPrice)

		// TODO: Custom normalizer for images
		normalizedProduct.Images = normalizeString(product.Images)

		// TODO: custom normalizer for rating
		normalizedProduct.RatingBase = product.RatingBase
		normalizedProduct.RatingStars = product.RatingStars
		normalizedProduct.RatingCount = product.RatingCount

		normalizedProducts = append(normalizedProducts, normalizedProduct)
	}

	return normalizedProducts
}

// Normalizes:

func normalizeString(content string) string {
	return strings.TrimSpace(content)
}

func normalizePrice(content string) float32 {
	var floatContent float32
	stringContent := strings.TrimSpace(content)

	regexMask := regexp.MustCompile(`[+-]?([0-9]*[.])?[0-9]+`)
	contentFiltered := regexMask.FindString(stringContent)

	if s, err := strconv.ParseFloat(contentFiltered, 32); err == nil {
		floatContent = float32(s)
	}

	return floatContent
}

//func normalizeRating(content string)  {
//	// TODO
//	//return content
//}

// TODO: make it return "bool" value (true/false) based on collected keyword
func normalizeAvailability(content string) bool {
	availableKeywords := []string{"instock", "in stock", "in-stock", "available"}

	isAvailable := false

	for _, keyword := range availableKeywords {
		if strings.TrimSpace(strings.ToLower(content)) == keyword {
			isAvailable = true
		}
	}

	return isAvailable
}
