package in

import "github.com/hillfolk/go-hexagonal-architecture/internal/account/domain"

type SendMoneyCommand struct {
	SourceAccountId domain.AccountId
	TargetAccountId domain.AccountId
	Amount          domain.Money
}

// TODO validate code 필요
func NewSendMoneyCommand(sourceAccountId domain.AccountId, targetAccountId domain.AccountId, amount domain.Money) SendMoneyCommand {
	return SendMoneyCommand{
		SourceAccountId: sourceAccountId,
		TargetAccountId: targetAccountId,
		Amount:          amount,
	}
}
