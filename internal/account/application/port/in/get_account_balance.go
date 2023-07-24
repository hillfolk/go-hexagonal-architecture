package in

import "github.com/hillfolk/go-hexagonal-architecture/internal/account/domain"

type GetAccountBalanceQuery interface {
	GetAccountBalance(accountId domain.AccountId) (domain.Money, error)
}
