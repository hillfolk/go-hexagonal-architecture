package service

import "github.com/hillfolk/go-hexagonal-architecture/internal/account/domain"

type MoneyTransferProperties interface {
	GetMaximumTransferThreshold() domain.Money
}
