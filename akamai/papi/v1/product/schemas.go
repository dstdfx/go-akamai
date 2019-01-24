package product

// Product encapsulates information on the product that determines
// a property’s available range of features.
type Product struct {
	// ID is a unique identifier for the product.
	ID string `json:"productId"`

	// Name is a descriptive name for the product.
	Name string `json:"productName"`
}
