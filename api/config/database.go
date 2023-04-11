package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"api/domain"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 初期化しておく
	if err := db.Migrator().DropTable(
		&domain.Club{},
		&domain.User{},
		&domain.UserClub{},
		&domain.Activity{},
		&domain.ActivityUser{},
	); err != nil {
		panic("failed to drop table")
	}

	if err := db.AutoMigrate(&domain.Club{}); err != nil {
		panic("failed to migrate club")
	}

	if err := db.AutoMigrate(&domain.User{}); err != nil {
		panic("failed to migrate user")
	}

	if err := db.AutoMigrate(&domain.Activity{}); err != nil {
		panic("failed to migrate activity")
	}

	if err := seeder(db); err != nil {
		panic("failed to seed")
	}

	return db
}

func seeder(db *gorm.DB) error {
	clubs := []domain.Club{
		{
			Name:     "サークルA",
			Readme:   "サークルAの説明\nこれはサークルAの説明です。",
			Category: "IT",
		},
		{
			Name:     "サークルB",
			Readme:   "サークルBの説明",
			Category: "Game",
		},
	}

	for _, club := range clubs {
		if err := db.Create(&club).Error; err != nil {
			return err
		}
	}

	users := []domain.User{
		{
			UID:    220001,
			Name:   "ほげ太郎",
			Class:  "IE1A",
			Icon:   "https://source.unsplash.com/random",
			Readme: "こんにちは！\nよろしくお願いします！",
			Clubs:  nil,
		},
		{
			UID:    220002,
			Name:   "ふが太郎",
			Class:  "IE1A",
			Icon:   "https://source.unsplash.com/random",
			Readme: "こんにちは！\nよろしくお願いします！",
			Clubs:  nil,
		},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			return err
		}

		// user 1 はサークルAに所属
		if user.ID == 1 {
			club := domain.Club{}
			if err := db.First(&club, 1).Error; err != nil {
				return err
			}

			if err := db.Model(&user).Association("Clubs").Append(&club); err != nil {
				return err
			}
		}
	}

	activities := []domain.Activity{
		{
			Place:  "サークルAの活動場所",
			Detail: "サークルAの活動詳細\nこれはサークルAの活動詳細です。",
		},
		{
			Place:  "サークルAの活動場所",
			Detail: "サークルA2の活動詳細\nこれはサークルAの活動詳細です。",
		},
	}

	for _, activity := range activities {
		if err := db.Create(&activity).Error; err != nil {
			return err
		}

		if activity.ID == 1 {
			club := domain.Club{}
			if err := db.First(&club, 1).Error; err != nil {
				return err
			}

			if err := db.Model(&activity).Update("club_id", club.ID).Error; err != nil {
				return err
			}

			user := domain.User{}
			if err := db.First(&user, 1).Error; err != nil {
				return err
			}

			if err := db.Model(&activity).Association("Users").Append(&user); err != nil {
				return err
			}
		}
	}

	return nil
}
