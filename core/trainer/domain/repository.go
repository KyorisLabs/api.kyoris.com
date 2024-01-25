package trainerDomain

type TrainerRepository interface {
	Create(trainer *CreateTrainerDTO) error
	AddRelation(trainerID int, userID int, rolID int) error
}
