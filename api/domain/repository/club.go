package repository

import "api/domain"

type ClubRepository interface {
	FindAll() ([]domain.Club, error)
	FindByID(id int) (domain.Club, error)
}
