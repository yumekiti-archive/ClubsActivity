package usecase

import (
	"api/domain"
	"api/domain/repository"
)

type ClubUsecase interface {
	FindAll() ([]domain.Club, error)
	FindByID(id int) (domain.Club, error)
}

type clubUsecase struct {
	clubRepository repository.ClubRepository
}

func NewClubUsecase(cr repository.ClubRepository) ClubUsecase {
	return &clubUsecase{
		clubRepository: cr,
	}
}

func (cu *clubUsecase) FindAll() ([]domain.Club, error) {
	return cu.clubRepository.FindAll()
}

func (cu *clubUsecase) FindByID(id int) (domain.Club, error) {
	return cu.clubRepository.FindByID(id)
}
