package service

import (
	"github.com/hillfolk/go-hexagonal-architecture/internal/user/domain"
)

type GetUserListService struct {
}

func NewGetUserListService() *GetUserListService {
	return &GetUserListService{}
}

func (s *GetUserListService) GetUserList() ([]domain.User, error) {
	//TODO implement
	return []domain.User{}, nil
}
