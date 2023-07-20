package domain

import "time"

type AccountId int64

func NewAccountId(id int64) AccountId {
	return AccountId(id)
}

type Account struct {
	AccountId       AccountId
	BaselineBalance Money
	ActiveWindow    *ActiveWindow
}

func withoutId(baselineBalance Money, window *ActiveWindow) *Account {
	return &Account{
		BaselineBalance: baselineBalance,
		ActiveWindow:    window,
	}
}

func (a *Account) GetAccountId() AccountId {
	return a.AccountId
}

func WithId(id AccountId, baselineBalance Money, window *ActiveWindow) *Account {
	return &Account{
		AccountId:       id,
		BaselineBalance: baselineBalance,
		ActiveWindow:    window,
	}
}

func (a *Account) CalculateBalance() Money {
	return Money{}.Add(a.BaselineBalance, a.ActiveWindow.CalculateBalance(a.AccountId))
}

func (a *Account) Withdraw(amount Money, targetAccountId AccountId) bool {
	if !a.ActiveWindow.CanWithdraw(amount) {
		return false
	}

	var withdrawl = NewActivity(
		a.AccountId,
		a.AccountId,
		targetAccountId,
		time.Now(),
		amount,
	)

	a.ActiveWindow.AddActivity(&withdrawl)
	return true
}

func (a *Account) Deposit(amount Money, sourceAccountId AccountId) bool {
	deposit := NewActivity(
		a.AccountId,
		sourceAccountId,
		a.AccountId,
		time.Now(),
		amount,
	)
	a.ActiveWindow.AddActivity(&deposit)
	return true
}

func (a *Account) GetActivityWindow() *ActiveWindow {
	return a.ActiveWindow
}
