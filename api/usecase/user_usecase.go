package usecase

import (
	"api/domain"
	"api/domain/repository"
)

type UserUsecase interface {
	FindMe(uid int) (*domain.User, error)
	Store(user *domain.User) (*domain.User, error)
	FindByClubID(id int) ([]*domain.User, error)
	Update(user *domain.User) (*domain.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: ur,
	}
}

func (uu *userUsecase) FindMe(uid int) (*domain.User, error) {
	user, err := uu.userRepository.FindMe(uid)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (uu *userUsecase) Store(user *domain.User) (*domain.User, error) {
	user, err := uu.userRepository.Store(user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (uu *userUsecase) FindByClubID(id int) ([]*domain.User, error) {
	users, err := uu.userRepository.FindByClubID(id)
	if err != nil {
		return users, err
	}
	return users, nil
}

func (uu *userUsecase) Update(user *domain.User) (*domain.User, error) {
	user, err := uu.userRepository.Update(user)
	if err != nil {
		return user, err
	}
	return user, nil
}
