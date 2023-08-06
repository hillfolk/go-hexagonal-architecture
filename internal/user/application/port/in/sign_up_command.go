package in

import "context"

type SignUpCommand interface {
	SignUp(ctx context.Context, mail string) error
}
