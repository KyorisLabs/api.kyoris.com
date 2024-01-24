package userInfrastructure

import userDomain "api.kyoris.com/core/user/domain"

// Create implements userDomain.UserRepository.
func (r *userRepository) Create(user *userDomain.UserDTO) error {

	// Comprobamos si existe ya el email
	userDTO := r.FindByEmail(user.Email)

	if userDTO != nil {
		return userDomain.ErrUserAlreadyExists
	}

	// Comprobamos si existe ya el username
	userDTO = r.FindByUsername(user.Username)

	if userDTO != nil {
		return userDomain.ErrUsernameAlreadyExists
	}

	// Creamos la consulta para crear el usuario
	err := r.db.Model(&userDomain.UserModel{}).Create(userDomain.UserModel{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Name:     user.Name,
		Gender:   user.Gender,
		Birthday: user.Birthday,
	}).Error

	if err != nil {
		return userDomain.ErrCreateUser
	}

	return nil
}
