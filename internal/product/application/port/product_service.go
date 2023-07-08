package port

type ProductService interface {
	// GetProductType returns the product type
	GetProductType() string
	// GetProductCode returns the product code
	GetProductCode() string
	// GetProductName returns the product name
	GetProductName() string
}
