package service

import "github.com/hillfolk/go-hexagonal-architecture/internal/account/domain"

type NoOpAccountLock interface {
	LockAccount(accountId domain.AccountId) error
	ReleaseAccount(accountId domain.AccountId) error
}
