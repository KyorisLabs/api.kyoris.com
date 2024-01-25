package authApplication

import (
	userDomain "api.kyoris.com/core/user/domain"
	passwordUtils "api.kyoris.com/utils/password"
)

// Register implements AuthService.
func (a *authService) Register(user *userDomain.User, ip string) (string, error) {
	// Encriptamos la contraseña
	passHash, err := passwordUtils.Encrypt(user.Password)
	if err != nil {
		return "", userDomain.ErrEncryptPassword
	}

	// Asignamos la contraseña encriptada al usuario
	user.Password = passHash

	// Creamos el usuario
	err = a.user.Create(user)
	if err != nil {
		return "", err
	}

	// Creamos la sesion
	token, err := a.auth.CreateSession(user.Id, ip)

	return token, err
}
