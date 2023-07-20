package service

import "github.com/hillfolk/go-hexagonal-architecture/internal/account/domain"

type MoneyTransferProperties interface {
	GetMaximumTransferThreshold() domain.Money
}

type DefaultMoneyTransferProperties struct {
	maximumTransferThreshold domain.Money
}

func (d DefaultMoneyTransferProperties) GetMaximumTransferThreshold() domain.Money {
	return d.maximumTransferThreshold
}

func NewDefaultMoneyTransferProperties() *DefaultMoneyTransferProperties {
	return &DefaultMoneyTransferProperties{
		maximumTransferThreshold: domain.Money{
			Amount: 1000000,
		},
	}
}
