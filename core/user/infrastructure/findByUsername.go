package userInfrastructure

import (
	"errors"

	userDomain "api.kyoris.com/core/user/domain"
	"gorm.io/gorm"
)

// FindByEmail implements userDomain.UserRepository.
func (r *userRepository) FindByUsername(username string) *userDomain.User {
	user := userDomain.User{}

	err := r.db.Model(&user).Where("username = ?", username).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	return &user
}
