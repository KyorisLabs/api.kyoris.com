package authApplication

import (
	sharedDomain "api.kyoris.com/core/_shared/domain"
	authDomain "api.kyoris.com/core/auth/domain"
	userDomain "api.kyoris.com/core/user/domain"
)

type AuthService interface {
	Register(user *userDomain.User, IP string) (string, error)
	Login(email string, password string, IP string) (string, error)
	Logout(token string) error
	CheckToken(token string) (int, error)
}

type authService struct {
	auth authDomain.AuthRepository
	user userDomain.UserRepository
}

// CheckToken implements AuthService.
func (*authService) CheckToken(token string) (int, error) {
	panic("unimplemented")
}

func NewAuthService(r *sharedDomain.Repositories) AuthService {
	return &authService{
		auth: r.AuthRepository,
		user: r.UserRepository,
	}
}
