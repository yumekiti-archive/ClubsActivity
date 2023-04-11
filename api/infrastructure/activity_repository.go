package infrastructure

import (
	"gorm.io/gorm"

	"api/domain"
	"api/domain/repository"
)

type activityRepository struct {
	Conn *gorm.DB
}

func NewActivityRepository(conn *gorm.DB) repository.ActivityRepository {
	return &activityRepository{Conn: conn}
}

func (ar *activityRepository) FindByClubID(clubID int) ([]domain.Activity, error) {
	activities := []domain.Activity{}
	if err := ar.Conn.
		Joins("JOIN clubs ON activities.club_id = clubs.id").
		Where("clubs.id = ?", clubID).
		Preload("Users").
		Find(&activities).
		Error; err != nil {
		return activities, err
	}
	return activities, nil
}
