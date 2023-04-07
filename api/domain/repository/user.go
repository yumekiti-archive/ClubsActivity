package repository

import "api/domain"

type UserRepository interface {
	FindMe(id int) (*domain.User, error)
	Store(user *domain.User) (*domain.User, error)
	FindByClubID(id int) ([]*domain.User, error)
}
