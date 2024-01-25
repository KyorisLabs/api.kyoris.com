package userDomain

import "time"

// Model de Usuario en la base de datos
type User struct {
	Id       uint      `gorm:"primaryKey;autoIncrement" json:"-"`
	Username string    `gorm:"unique;notNull" json:"username"`
	Password string    `gorm:"notNull" json:"-"`
	Email    string    `gorm:"unique;notNull" json:"email"`
	Name     string    `gorm:"notNull" json:"name"`
	Gender   int       `gorm:"notNull" json:"gender"`
	Birthday time.Time `gorm:"notNull;type:date" json:"birthday"`
	CreateAt time.Time `gorm:"autoCreateTime" json:"create_at"`
}
