package infrastructure

import (
	"time"

	"api/domain"
	"api/domain/repository"
)

type userRepository struct{}

func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

func (ur *userRepository) FindMe() (domain.User, error) {
	return domain.User{
		UserID:           1,
		UserUID:          "test",
		UserName:         "test",
		UserClass:        "test",
		UserIcon:         "test",
		UserReadme:       "test",
		AffiliatedClubID: []int{1, 2, 3},
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}, nil
}
