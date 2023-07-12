package service

import "github.com/hillfolk/go-hexagonal-architecture/internal/account/application/port/in"

type SendMoneyService struct {
}

func NewSendMoneyService() in.SendMoneyUseCase {
	return &SendMoneyService{}
}

func (s *SendMoneyService) SendMoney(cmd in.SendMoneyCommand) error {
	//TODO implement
	return nil
}
