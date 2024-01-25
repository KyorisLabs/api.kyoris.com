package trainerInfrastructure

import (
	sharedDomain "api.kyoris.com/core/_shared/domain"
	trainerDomain "api.kyoris.com/core/trainer/domain"
)

func (r *trainerRepository) Create(data *trainerDomain.CreateTrainerDTO) error {
	if data.Name == "" {
		return trainerDomain.ErrTrainerNameIsEmpty
	}

	trainer := &trainerDomain.TrainerModel{
		Name: data.Name,
	}

	// Crear trainer
	if err := r.db.Model(&trainerDomain.CreateTrainerDTO{}).Create(&trainer).Error; err != nil {
		return sharedDomain.ErrInternalServerError
	}

	// Insertar relación
	//! ROL ID 1 == ADMIN
	if err := r.AddRelation(int(trainer.Id), data.UserID, 1); err != nil {
		return err
	}

	return nil
}
