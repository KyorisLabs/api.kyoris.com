package userDomain

import "time"

// Model de Usuario en la base de datos
type UserModel struct {
	Id       uint      `gorm:"primaryKey;autoIncrement"`
	Username string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
	Email    string    `gorm:"unique;not null"`
	Name     string    `gorm:"not null"`
	Gender   int       `gorm:"not null"`
	Birthday time.Time `gorm:"not null"`
	CreateAt time.Time `gorm:"autoCreateTime"`
}
