package service

import (
	"github.com/hillfolk/go-hexagonal-architecture/internal/account/adapter/out/persistance"
	"github.com/hillfolk/go-hexagonal-architecture/internal/account/application/port/out"
	"github.com/hillfolk/go-hexagonal-architecture/internal/account/domain"
	"time"
)

type GetAccountBalanceService struct {
	loadAccountPort out.LoadAccountPort
}

func NewGetAccountBalanceService() *GetAccountBalanceService {
	return &GetAccountBalanceService{
		loadAccountPort: persistance.NewAccountPersistenceAdapter(),
	}
}

func (s *GetAccountBalanceService) GetAccountBalance(accountId domain.AccountId) (domain.Money, error) {
	account, err := s.loadAccountPort.LoadAccount(accountId, time.Now())
	if err != nil {
		return domain.Money{}, err
	}
	return account.CalculateBalance(), nil
}
