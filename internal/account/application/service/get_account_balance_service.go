package service

import "github.com/hillfolk/go-hexagonal-architecture/internal/account/domain"

type GetAccountBalanceService interface {
	GetAccountBalance(accountId domain.AccountId) (int64, error)
}
