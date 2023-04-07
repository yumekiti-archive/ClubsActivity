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
		UserID:     1,
		UserUID:    "test",
		UserName:   "test",
		UserClass:  "test",
		UserIcon:   "test",
		UserReadme: "test",
		AffiliatedClubs: []domain.Club{
			domain.Club{
				ClubID:          1,
				ClubName:        "test",
				ClubDescription: "test",
				ClubCategory:    "test",
				CreatedAt:       time.Now(),
				UpdatedAt:       time.Now(),
			},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
