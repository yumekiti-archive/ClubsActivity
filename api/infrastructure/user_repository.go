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
	var affiliatedClubs []domain.Club
	affiliatedClubs = append(affiliatedClubs, domain.Club{
		ClubID:          1,
		ClubName:        "test",
		ClubDescription: "test",
		ClubCategory:    "test",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	})

	return domain.User{
		UserID:     1,
		UserUID:    "test",
		UserName:   "test",
		UserClass:  "test",
		UserIcon:   "test",
		UserReadme: "test",
		AffiliatedClubs: affiliatedClubs,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
