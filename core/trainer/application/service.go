package trainerApplication

import (
	sharedDomain "api.kyoris.com/core/_shared/domain"
	trainerDomain "api.kyoris.com/core/trainer/domain"
)

type TrainerService interface {
	Create(payload *trainerDomain.CreateTrainerDTO) error
}

type trainerService struct {
	trainer trainerDomain.TrainerRepository
}

func NewTrainerService(r *sharedDomain.Repositories) TrainerService {
	return &trainerService{
		trainer: r.TrainerRepository,
	}
}
