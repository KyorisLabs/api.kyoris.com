package authApplication

import (
	authDomain "api.kyoris.com/core/auth/domain"
	userDomain "api.kyoris.com/core/user/domain"
	passwordUtils "api.kyoris.com/utils/password"
)

// Login implements AuthService.
func (a *authService) Login(email string, password string, IP string) (string, error) {
	user := a.user.FindByEmail(email)

	if user == nil {
		return "", userDomain.ErrNotExistsUser
	}

	// Comprobamos que la contraseña sea correcta
	if !passwordUtils.Validate(password, user.Password) {
		return "", userDomain.ErrNotExistsUser
	}

	// Crear una sesion
	token, err := a.auth.CreateSession(user.Id, IP)

	if err != nil {
		return "", authDomain.ErrSessionCreation
	}

	return token, nil
}
