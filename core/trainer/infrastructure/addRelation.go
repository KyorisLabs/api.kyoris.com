package trainerInfrastructure

import trainerDomain "api.kyoris.com/core/trainer/domain"

func (r trainerRepository) AddRelation(trainerID int, userID int, rolID int) error {
	if err := r.db.Model(&trainerDomain.TrainerUserModel{}).Create(&trainerDomain.TrainerUserModel{TrainerID: trainerID, UserID: userID, RolID: rolID}).Error; err != nil {
		return err
	}

	return nil
}
