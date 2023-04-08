package infrastructure

import (
	"gorm.io/gorm"
	"time"

	"api/domain"
	"api/domain/repository"
)

type clubRepository struct {
	Conn *gorm.DB
}

func NewClubRepository(conn *gorm.DB) repository.ClubRepository {
	return &clubRepository{Conn: conn}
}

func (cr *clubRepository) FindAll() ([]domain.Club, error) {
	return []domain.Club{
		{
			ID:          1,
			Name:        "test",
			Readme: "test",
			Category:    "test",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			Name:        "test",
			Readme: "test",
			Category:    "test",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          3,
			Name:        "test",
			Readme: "test",
			Category:    "test",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}, nil
}

func (cr *clubRepository) FindByID(id int) (domain.Club, error) {
	return domain.Club{
		ID:          1,
		Name:        "test",
		Readme: "test",
		Category:    "test",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
