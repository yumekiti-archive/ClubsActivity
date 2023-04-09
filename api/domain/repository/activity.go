package repository

import "api/domain"

type ActivityRepository interface {
	FindByClubID(clubID int) ([]domain.Activity, error)
}
