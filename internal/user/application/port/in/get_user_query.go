package in

import "github.com/hillfolk/go-hexagonal-architecture/internal/user/domain"

type GetUserQuery interface {
	GetUser(userId string) (*domain.User, error)
}
