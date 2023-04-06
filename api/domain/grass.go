/*
## Grass Data（草データ）

| カラム名            | 説明                    | 型       | Unique | Nullable |
| ------------------- | ----------------------- | -------- | ------ | -------- |
| grass_id            | 草 ID                   | Integer  | Yes    | No       |
| grass_user_id       | 草を生やしたユーザー ID | Integer  | No     | No       |
| grassed_activity_id | 草を生やされた活動 ID   | Integer  | No     | No       |
| created_at          | 作成日時                | DateTime | No     | No       |
| updated_at          | 更新日時                | DateTime | No     | No       |
*/

package domain

import (
	"time"
)

type Grass struct {
	GrassID            int
	GrassUserID        int
	GrassedActivityID  int
	CreatedAt          time.Time
	UpdatedAt          time.Time
}