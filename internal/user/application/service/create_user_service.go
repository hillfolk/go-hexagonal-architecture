package service

type CreateUserService struct {
}

func NewCreateUserService() *CreateUserService {
	return &CreateUserService{}
}

func (s *CreateUserService) CreateUser() (string, error) {
	//TODO implement
	return "user_id", nil
}
