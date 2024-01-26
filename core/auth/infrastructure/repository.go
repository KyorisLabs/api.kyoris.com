package authInfrastructure

import (
	authDomain "api.kyoris.com/core/auth/domain"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) authDomain.AuthRepository {
	return &authRepository{db}
}
