package userInfrastructure

import (
	userDomain "api.kyoris.com/core/user/domain"
)

// Create implements userDomain.UserRepository.
func (r *userRepository) Create(user *userDomain.User) error {

	// Comprobamos si existe ya el email
	userModel := r.FindByEmail(user.Email)

	if userModel != nil {
		return userDomain.ErrEmailAlreadyExists
	}

	// Comprobamos si existe ya el username
	userModel = r.FindByUsername(user.Username)

	if userModel != nil {
		return userDomain.ErrUsernameAlreadyExists
	}

	// Creamos la consulta para crear el usuario
	err := r.db.Model(&userDomain.User{}).Create(&user).Error

	if err != nil {
		return userDomain.ErrCreateUser
	}

	return nil
}
