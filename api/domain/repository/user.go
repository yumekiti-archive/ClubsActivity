package repository

import "api/domain"

type UserRepository interface {
	FindMe() (domain.User, error)
}
