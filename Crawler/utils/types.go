package Utils

// Structs that will build the JSON

// OUTPUT:

// RawProduct is the raw data EXTRACTED using Tags
type RawProduct struct {
	ProductCode string
	ProductName string
	ProductURL  string

	InStock     string
	FinalPrice  string
	FromPrice   string
	Description string
	Images      string

	RatingBase  string
	RatingStars string
	RatingCount string

	Brand  string
	Seller string
}

// NormalizedProduct is the RawProduct after being normalized
type NormalizedProduct struct {
	ProductCode string
	ProductName string
	ProductURL  string

	InStock     bool
	FinalPrice  float32
	FromPrice   float32
	Description string
	Images      string

	// TODO: MAKE IT INTEGER
	RatingBase  string
	RatingStars string
	RatingCount string

	Brand  string
	Seller string
}

// INPUT:

// SearchJob will be the INPUT for Search analysis
type SearchJob struct {
	RetailerId          int      `json:"retailer-id"`
	ProductsIdentifiers []string `json:"products"`
	Url                 string   `json:"url"`
	MaxPages            int      `json:"max-pages,omitempty"`
}

// ProductJob will be the INPUT for Product analysis
type ProductJob struct {
	ProductsIdentifiers []string `json:"products"`
	Urls                []string `json:"urls"`
}

type Tags struct {
	ProductBase []TagNode `json:"ProductBase"`
	NextedData  struct {
		ProductName         []TagNode `json:"productName,omitempty"`
		ProductCode         []TagNode `json:"productCode,omitempty"`
		ProductURL          []TagNode `json:"productURL,omitempty"`
		ProductAvailability []TagNode `json:"productAvailability,omitempty"`
		ProductDescription  []TagNode `json:"ProductDescription,omitempty"`
		ProductFinalPrice   []TagNode `json:"productFinalPrice,omitempty"`
		ProductFromPrice    []TagNode `json:"productFromPrice,omitempty"`
		ProductImages       []TagNode `json:"productImages,omitempty"`
		ProductRatingBase   []TagNode `json:"productRatingBase,omitempty"`
		ProductRatingStars  []TagNode `json:"productRatingStars,omitempty"`
		ProductRatingCount  []TagNode `json:"productRatingCount,omitempty"`
		ProductBrand        []TagNode `json:"ProductBrand,omitempty"`
		ProductSeller       []TagNode `json:"ProductSeller,omitempty"`
	} `json:"NestedData"`
}

type TagNode struct {
	Selector        string `json:"selector"`
	CustomAttribute string `json:"customAttribute,omitempty"`
}
