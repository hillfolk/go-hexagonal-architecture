package domain

import (
	"strings"
	"time"
)

type ToString interface {
	ToString() string
}
type Product struct {
	ProductId   string
	ProductName string
	ProductType ProductType
	Price       Currency
	Created     time.Time
}

func (p *Product) ToString() string {
	var sb strings.Builder
	sb.WriteString("{")
	sb.WriteString(p.ProductId)
	sb.WriteString(",")
	sb.WriteString(string(p.ProductType))
	sb.WriteString(",")
	sb.WriteString(p.ProductName)
	sb.WriteString(",")
	sb.WriteString(p.Price.String())
	sb.WriteString("}")
	return sb.String()
}
