package usecase

import (
	"api/domain"
	"api/domain/repository"
)

type ActivityUsecase interface {
	FindByClubID(clubID int) ([]domain.Activity, error)
}

type activityUsecase struct {
	activityRepository repository.ActivityRepository
}

func NewActivityUsecase(ar repository.ActivityRepository) ActivityUsecase {
	return &activityUsecase{activityRepository: ar}
}

func (au *activityUsecase) FindByClubID(clubID int) ([]domain.Activity, error) {
	activities, err := au.activityRepository.FindByClubID(clubID)
	if err != nil {
		return nil, err
	}

	return activities, nil
}
