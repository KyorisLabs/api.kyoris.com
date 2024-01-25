package authInfrastructure

import (
	authDomain "api.kyoris.com/core/auth/domain"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

// DeleteSession implements authDomain.AuthRepository.
func (*authRepository) DeleteSession(token string) error {
	panic("unimplemented")
}

// FindSession implements authDomain.AuthRepository.
func (*authRepository) FindSession(token string) (*authDomain.Session, error) {
	panic("unimplemented")
}

func NewMySQLRepository(db *gorm.DB) authDomain.AuthRepository {
	return &authRepository{db}
}
