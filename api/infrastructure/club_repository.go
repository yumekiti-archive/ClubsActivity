package infrastructure

import (
	"gorm.io/gorm"

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
	var clubs []domain.Club
	if err := cr.Conn.Find(&clubs).Error; err != nil {
		return nil, err
	}

	return clubs, nil
}

func (cr *clubRepository) FindByID(id int) (domain.Club, error) {
	var club domain.Club
	if err := cr.Conn.First(&club, id).Error; err != nil {
		return domain.Club{}, err
	}

	return club, nil
}
