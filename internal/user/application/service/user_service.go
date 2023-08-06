package service

import "github.com/hillfolk/go-hexagonal-architecture/internal/user/domain"

// UserService  여러 인터페이스를 함께 구현 했지만 호출은 모두 인터페이스를 가자 호출한다.
type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUser(userId string) (*domain.User, error) {
	//TODO implement
	return nil, nil
}

func (s *UserService) GetUserList() ([]domain.User, error) {
	//TODO implement
	return []domain.User{}, nil
}

func (s *UserService) CreateUser(user domain.User) (*domain.User, error) {
	//TODO implement
	return nil, nil
}
