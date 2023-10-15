package Utils

import (
	"github.com/PuerkitoBio/goquery"
)

// ExtractSearch will extract the data from the HTML using the Tags
func (productTag Tags) ExtractSearch(html *goquery.Document) []RawProduct {
	var productsBase *goquery.Selection
	var rawProducts []RawProduct

	for _, ProductBaseTag := range productTag.ProductBase {
		productsBase = html.Find(ProductBaseTag.Selector)

		if productsBase.Length() < 1 {
			continue
		}

		for _, node := range productsBase.Nodes {
			element := goquery.NewDocumentFromNode(node)

			rawProducts = append(rawProducts, ApplyRawTags(productTag, element.Children()))
		}
	}
	return rawProducts
}

func (productTag Tags) ExtractProduct(html *goquery.Document) []RawProduct {
	var extractedProduct []RawProduct
	rawProduct := ApplyRawTags(productTag, html.Contents())
	extractedProduct = append(extractedProduct, rawProduct)
	return extractedProduct
}

func ApplyRawTags(productTag Tags, productsBase *goquery.Selection) RawProduct {
	product := RawProduct{}

	NextedTags := productTag.NextedData

	// TODO: Fix CTRL C + CTRL V code (make it less redundant)
	IterateTags(NextedTags.ProductCode, productsBase, &product.ProductCode)
	IterateTags(NextedTags.ProductName, productsBase, &product.ProductName)
	IterateTags(NextedTags.ProductURL, productsBase, &product.ProductURL)
	IterateTags(NextedTags.ProductAvailability, productsBase, &product.InStock)
	IterateTags(NextedTags.ProductFinalPrice, productsBase, &product.FinalPrice)
	IterateTags(NextedTags.ProductFromPrice, productsBase, &product.FromPrice)
	IterateTags(NextedTags.ProductDescription, productsBase, &product.Description)
	IterateTags(NextedTags.ProductImages, productsBase, &product.Images)
	IterateTags(NextedTags.ProductRatingBase, productsBase, &product.RatingBase)
	IterateTags(NextedTags.ProductRatingStars, productsBase, &product.RatingStars)
	IterateTags(NextedTags.ProductRatingCount, productsBase, &product.RatingCount)
	IterateTags(NextedTags.ProductBrand, productsBase, &product.Brand)
	IterateTags(NextedTags.ProductSeller, productsBase, &product.Seller)

	return product
}

func IterateTags(NextedTags []TagNode, productsBase *goquery.Selection, field *string) {
	for i := len(NextedTags) - 1; i >= 0; i-- {
		fieldElement := productsBase.Find(NextedTags[i].Selector)
		fieldCustomAttribute := NextedTags[i].CustomAttribute

		if fieldElement.Length() == 0 {
			continue
		}

		if len(fieldCustomAttribute) > 0 {
			_, customFieldExists := fieldElement.Attr(fieldCustomAttribute)
			if customFieldExists {
				*field, _ = fieldElement.Attr(fieldCustomAttribute)
			} else {
				*field = fieldElement.Text()
			}
		} else {
			*field = fieldElement.Text()
		}
	}
}
