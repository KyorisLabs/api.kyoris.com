package userApplication

import userDomain "api.kyoris.com/core/user/domain"

// GetById implements UserService.
func (s *userService) GetById(id int) *userDomain.User {
	return s.user.FindById(uint(id))
}
