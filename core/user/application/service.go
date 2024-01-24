package userApplication

import userDomain "api.kyoris.com/core/user/domain"

type UserService interface {
	Register(user *userDomain.UserDTO) error
	GetById(id int) *userDomain.UserDTO
}

type userService struct {
	user userDomain.UserRepository
}

func NewService(user userDomain.UserRepository) UserService {
	return &userService{user}
}
