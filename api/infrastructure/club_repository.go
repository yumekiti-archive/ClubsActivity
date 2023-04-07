package infrastructure

import (
	"time"

	"api/domain"
	"api/domain/repository"
)

type clubRepository struct{}

func NewClubRepository() repository.ClubRepository {
	return &clubRepository{}
}

func (cr *clubRepository) FindAll() ([]domain.Club, error) {
	return []domain.Club{
		{
			ClubID:          1,
			ClubName:        "test",
			ClubDescription: "test",
			ClubCategory:    "test",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
		{
			ClubID:          2,
			ClubName:        "test",
			ClubDescription: "test",
			ClubCategory:    "test",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
		{
			ClubID:          3,
			ClubName:        "test",
			ClubDescription: "test",
			ClubCategory:    "test",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
	}, nil
}

func (cr *clubRepository) FindByID(id int) (domain.Club, error) {
	return domain.Club{
		ClubID:          1,
		ClubName:        "test",
		ClubDescription: "test",
		ClubCategory:    "test",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}, nil
}
