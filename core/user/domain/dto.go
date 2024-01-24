package userDomain

import "time"

// Modelo de datos
type UserDTO struct {
	Id       uint
	Username string    `json:"username"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Name     string    `json:"name"`
	Gender   int       `json:"gender"`
	Birthday time.Time `json:"birthday"`
	CreateAt time.Time `json:"create_at"`
}
