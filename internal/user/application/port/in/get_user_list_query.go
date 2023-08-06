package in

import (
	"github.com/hillfolk/go-hexagonal-architecture/internal/user/domain"
)

type GetUserListQuery interface {
	GetUserList() ([]*domain.User, error)
}
