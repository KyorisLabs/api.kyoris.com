package userInfrastructure

import (
	"errors"

	userDomain "api.kyoris.com/core/user/domain"
	"gorm.io/gorm"
)

// FindByEmail implements userDomain.UserRepository.
func (r *userRepository) FindByEmail(email string) *userDomain.UserDTO {
	user := userDomain.UserModel{}

	err := r.db.Model(&user).Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	return &userDomain.UserDTO{
		Id:       user.Id,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Name:     user.Name,
		Gender:   user.Gender,
		Birthday: user.Birthday,
		CreateAt: user.CreateAt,
	}
}
