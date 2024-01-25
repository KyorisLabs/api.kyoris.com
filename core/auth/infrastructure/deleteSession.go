package authInfrastructure

import authDomain "api.kyoris.com/core/auth/domain"

// DeleteSession implements authDomain.AuthRepository.
func (a *authRepository) DeleteSession(token string) error {
	if err := a.db.Where("token = ?", token).Delete(&authDomain.Session{}).Error; err != nil {
		return authDomain.ErrSessionDeletion
	}

	return nil
}
