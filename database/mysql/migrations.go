package mysql

import (
	trainerDomain "api.kyoris.com/core/trainer/domain"
	userDomain "api.kyoris.com/core/user/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.Table("users").AutoMigrate(&userDomain.UserModel{})
	db.Table("trainers").AutoMigrate(&trainerDomain.TrainerModel{})
	db.Table("trainer_users").AutoMigrate(&trainerDomain.TrainerUserModel{})
	db.Table("trainer_tax").AutoMigrate(&trainerDomain.TrainerTax{})
}
