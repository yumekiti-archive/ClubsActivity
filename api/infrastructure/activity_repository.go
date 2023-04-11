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
	var activities []domain.Activity
	if err := ar.Conn.Preload("Users").Preload("Club").Where("club_id = ?", clubID).Find(&activities).Error; err != nil {
		return nil, err
	}
	return activities, nil
}
