package in

import "github.com/hillfolk/go-hexagonal-architecture/internal/account/domain"

type GetAccountBalanceQuery interface {
	GetAccountBalance(accountId string) (domain.Money, error)
}
