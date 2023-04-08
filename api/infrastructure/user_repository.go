package infrastructure

import (
	"gorm.io/gorm"

	"api/domain"
	"api/domain/repository"
)

type userRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) repository.UserRepository {
	return &userRepository{Conn: conn}
}

func (ur *userRepository) FindMe(uid int) (*domain.User, error) {
	user := &domain.User{}
	if err := ur.Conn.Where("uid = ?", uid).Preload("Clubs").First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (ur *userRepository) Store(user *domain.User) (*domain.User, error) {
	if err := ur.Conn.Create(user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (ur *userRepository) FindByClubID(id int) ([]*domain.User, error) {
	users := []*domain.User{}
	if err := ur.Conn.Where("club_id = ?", id).Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}
