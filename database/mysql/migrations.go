package mysql

import (
	authDomain "api.kyoris.com/core/auth/domain"
	trainerDomain "api.kyoris.com/core/trainer/domain"
	userDomain "api.kyoris.com/core/user/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.Table("users").AutoMigrate(&userDomain.User{})
	db.Table("sessions").AutoMigrate(&authDomain.Session{})
	db.Table("trainers").AutoMigrate(&trainerDomain.TrainerModel{})
	db.Table("trainer_users").AutoMigrate(&trainerDomain.TrainerUserModel{})
	db.Table("trainer_tax").AutoMigrate(&trainerDomain.TrainerTax{})
}
