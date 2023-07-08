package domain

import "fmt"

type Currency struct {
	Amount   float64
	Currency string
}

func (c *Currency) Add(amount float64) {
	c.Amount += amount
}

func (c *Currency) Sub(amount float64) {
	c.Amount -= amount
}

func (c *Currency) Multiply(amount float64) {
	c.Amount *= amount
}

func (c *Currency) Divide(amount float64) {
	c.Amount /= amount
}

func (c *Currency) IsEqual(amount float64) bool {
	return c.Amount == amount
}

func (c *Currency) IsGreaterThan(amount float64) bool {
	return c.Amount > amount
}

func (c *Currency) String() string {
	return fmt.Sprintf("%.2f %s", c.Amount, c.Currency)
}
