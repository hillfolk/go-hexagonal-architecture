package service

import (
	"github.com/hillfolk/go-hexagonal-architecture/internal/account/adapter/out/persistance"
	"github.com/hillfolk/go-hexagonal-architecture/internal/account/application/port/in"
	"github.com/hillfolk/go-hexagonal-architecture/internal/account/application/port/out"
	"github.com/pkg/errors"
	"time"
)

type SendMoneyService struct {
	loadAccountPort         out.LoadAccountPort
	accountLockPort         out.AccountLockPort
	updateAccountStatePort  out.UpdateAccountStatePort
	MoneyTransferProperties MoneyTransferProperties
}

func NewSendMoneyService() in.SendMoneyUseCase {
	adapter := persistance.NewAccountPersistenceAdapter()
	return &SendMoneyService{
		loadAccountPort:         adapter,
		accountLockPort:         NewNoOpAccountLock(),
		updateAccountStatePort:  adapter,
		MoneyTransferProperties: NewDefaultMoneyTransferProperties(),
	}
}

func (s *SendMoneyService) SendMoney(cmd in.SendMoneyCommand) error {
	if err := s.checkThreadhold(cmd); err != nil {
		return errors.Wrap(err, "checkThreadhold error")
	}
	baselineDate := time.Now().Add(-10 * 24 * time.Hour)
	sourceAccount, err := s.loadAccountPort.LoadAccount(cmd.SourceAccountId, baselineDate)
	if err != nil {
		return err
	}
	targetAccount, err := s.loadAccountPort.LoadAccount(cmd.TargetAccountId, baselineDate)
	if err != nil {
		return err
	}
	sourceAccountId := sourceAccount.AccountId
	targetAccountId := targetAccount.AccountId

	err = s.accountLockPort.LockAccount(sourceAccountId)
	if err != nil {
		return err
	}
	s.accountLockPort.LockAccount(sourceAccountId)
	if !sourceAccount.Withdraw(cmd.Amount, targetAccountId) {
		if err := s.accountLockPort.ReleaseAccount(sourceAccountId); err != nil {
			return errors.Wrap(err, "error while releasing account to source account")
		}
		return nil
	}

	if err := s.updateAccountStatePort.UpdateActivities(sourceAccount); err != nil {
		return errors.Wrap(err, "error while updating source account activities")
	}
	if err := s.updateAccountStatePort.UpdateActivities(targetAccount); err != nil {
		return errors.Wrap(err, "error while updating target account activities")
	}

	if err := s.accountLockPort.ReleaseAccount(sourceAccountId); err != nil {
		return errors.Wrap(err, "error while releasing account to source account")
	}
	if err := s.accountLockPort.ReleaseAccount(targetAccountId); err != nil {
		return errors.Wrap(err, "error while releasing account to target account")
	}

	return nil
}
func (s *SendMoneyService) checkThreadhold(cmd in.SendMoneyCommand) error {
	if cmd.Amount.IsGreaterThan(s.MoneyTransferProperties.GetMaximumTransferThreshold()) {
		return errors.New("Maximum transfer threshold reached")
	}
	return nil
}
