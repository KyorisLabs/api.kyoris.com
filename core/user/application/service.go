package userApplication

import (
	sharedDomain "api.kyoris.com/core/_shared/domain"
	userDomain "api.kyoris.com/core/user/domain"
)

type UserService interface {
	Register(user *userDomain.User) error
	GetById(id int) *userDomain.User
}

type userService struct {
	user userDomain.UserRepository
}

func NewService(r *sharedDomain.Repositories) UserService {
	return &userService{
		user: r.UserRepository,
	}
}
