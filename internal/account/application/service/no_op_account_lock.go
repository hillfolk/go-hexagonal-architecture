package service

import "github.com/hillfolk/go-hexagonal-architecture/internal/account/domain"

type NoOpAccountLock struct {
}

func NewNoOpAccountLock() *NoOpAccountLock {
	return &NoOpAccountLock{}
}

func (l *NoOpAccountLock) LockAccount(accountId domain.AccountId) error {
	return nil
}

func (l *NoOpAccountLock) ReleaseAccount(accountId domain.AccountId) error {
	return nil
}
