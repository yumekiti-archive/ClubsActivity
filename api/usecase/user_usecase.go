package usecase

import (
	"api/domain"
	"api/domain/repository"
)

type UserUsecase interface {
	FindMe() (domain.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: ur,
	}
}

func (uu *userUsecase) FindMe() (domain.User, error) {
	return uu.userRepository.FindMe()
}
