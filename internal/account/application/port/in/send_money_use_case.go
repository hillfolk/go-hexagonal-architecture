package in

type SendMoneyUseCase interface {
	SendMoney(cmd SendMoneyCommand) error
}
