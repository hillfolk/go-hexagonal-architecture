package domain

type ProductType string // ProductType is a type of product

const (
	ProductTypeLoan   ProductType = "Loan"
	ProductTypeFund   ProductType = "Fund"
	ProductTypeEquity ProductType = "Equity"
)
