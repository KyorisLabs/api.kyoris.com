package trainerInfrastructure

import (
	trainerDomain "api.kyoris.com/core/trainer/domain"
	"gorm.io/gorm"
)

type trainerRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) trainerDomain.TrainerRepository {
	return &trainerRepository{db}
}
