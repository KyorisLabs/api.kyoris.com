package userInfrastructure

import (
	userDomain "api.kyoris.com/core/user/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) userDomain.UserRepository {
	return &userRepository{db}
}
