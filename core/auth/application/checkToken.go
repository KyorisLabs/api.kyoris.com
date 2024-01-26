package authApplication

import (
	authDomain "api.kyoris.com/core/auth/domain"
	tokenUtils "api.kyoris.com/utils/token"
)

// CheckToken implements AuthService.
func (s *authService) CheckToken(token string) (int, error) {
	token, ok := tokenUtils.ValidateToken(token)

	if !ok {
		return 0, authDomain.ErrInvalidToken
	}

	// Comprobamos si existe el token y validamo el token.
	model, err := s.auth.FindSession(token)

	return int(model.UserID), err
}
