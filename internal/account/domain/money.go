package domain

type Money struct {
	Amount int64
}

func NewMoney(amount int64) Money {
	return Money{
		Amount: amount,
	}
}

func (Money) Add(a Money, b Money) Money {
	return Money{
		Amount: a.Amount + b.Amount,
	}
}

func (Money) Subtract(a Money, b Money) Money {
	return Money{
		Amount: a.Amount - b.Amount,
	}
}

func (m Money) isPositiveOrZero() bool {
	return m.Amount >= 0
}

func (m Money) IsNegative() bool {
	return m.Amount < 0
}

func (m Money) IsPositive() bool {
	return m.Amount >= 0
}
func (m Money) IsGreaterThan(targetMoney Money) bool {
	return m.Amount > targetMoney.Amount
}

func (m *Money) Minus(targetMoney Money) {
	m.Amount -= targetMoney.Amount
}

func (m *Money) Plus(targetMoney Money) {
	m.Amount += targetMoney.Amount
}

func (m Money) Negate() Money {
	return Money{
		Amount: -m.Amount,
	}
}
