package service

import (
	"github.com/hillfolk/go-hexagonal-architecture/internal/user/adapter/out/persistance"
	"github.com/hillfolk/go-hexagonal-architecture/internal/user/application/port/out"
	"github.com/hillfolk/go-hexagonal-architecture/internal/user/domain"
)

type GetUserService struct {
	getUserService  out.GetUserPort
	findUserService out.FindUsersPort
}

func NewGetUserService() *GetUserService {
	adapter := persistance.NewUserPersistenceAdapter()
	return &GetUserService{
		getUserService:  adapter,
		findUserService: adapter,
	}
}

func (s *GetUserService) GetUser(userId string) (*domain.User, error) {
	user, err := s.getUserService.GetUser(userId)
	return user, err
}

func (s *GetUserService) GetUserList() ([]*domain.User, error) {
	users, err := s.findUserService.GetUserList()
	return users, err
}
