package in

import "context"

type SignInCommand interface {
	SignIn(ctx context.Context, email, password string) error
}
