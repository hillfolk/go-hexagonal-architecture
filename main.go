package main

import (
	"fmt"
	product "github.com/hillfolk/go-hexagonal-architecture/internal/product/domain"
)

func main() {

	p := &product.Product{
		ProductId:   "1",
		ProductName: "Product 1",
		ProductType: product.ProductTypeEquity,
		Price: product.Currency{
			Amount:   100,
			Currency: "USD",
		},
	}

	fmt.Println("Hello, Hexagonal Architecture!" + p.ToString())

}
