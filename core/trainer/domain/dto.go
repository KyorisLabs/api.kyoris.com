package trainerDomain

type CreateTrainerDTO struct {
	UserID int    `json:"-"`
	Name   string `json:"name"`
}
