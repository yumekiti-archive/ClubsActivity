/*
## User Data（ユーザーデータ）

| カラム名           | 説明                       | 型        | Unique | Nullable |
| ------------------ | -------------------------- | --------- | ------ | -------- |
| user_id            | ユーザー ID                | Integer   | Yes    | No       |
| user_uid           | ユーザー UID               | String    | Yes    | No       |
| user_name          | ユーザー名                 | String    | No     | No       |
| user_class         | 学年                       | String    | No     | No       |
| user_icon          | ユーザーアイコン           | String    | No     | Yes      |
| user_readme        | 自己紹介                   | String    | No     | Yes      |
| affiliated_club_id | 所属している部活動（複数） | Integer[] | No     | Yes      |
| created_at         | 作成日時                   | DateTime  | No     | No       |
| updated_at         | 更新日時                   | DateTime  | No     | No       |
*/

package domain

import (
	"time"
)

type User struct {
	ID        int `gorm:"primary_key"`
	UID       int `gorm:"unique"`
	Name      string
	Class     string
	Icon      string
	Readme    string
	Clubs     []Club    `gorm:"many2many:user_clubs;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type UserClub struct {
	UserID uint
	ClubID uint
}
