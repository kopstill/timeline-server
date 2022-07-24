package service

type UserService interface {
	Login() error
	Logout() error
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (userService) Login() error {
	return nil
}

func (userService) Logout() error {
	return nil
}
