package trainerDomain

import "time"

// Modelo de base de datos de la tabla Trainers
type Trainer struct {
	Id   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"notNull"`
	Plan int    `gorm:"default:0"`

	// Meta
	CreateAt time.Time `gorm:"autoCreateTime"`
}

type TrainerUserModel struct {
	Id        uint      `gorm:"primaryKey;autoIncrement"`
	TrainerID int       `gorm:"notNull"`
	UserID    int       `gorm:"notNull"`
	RolID     int       `gorm:"notNull"`
	CreateAt  time.Time `gorm:"autoCreateTime"`
}
