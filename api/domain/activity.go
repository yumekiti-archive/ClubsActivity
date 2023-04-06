/*
## Activity Data（活動データ）

| カラム名         | 説明                    | 型        | Unique | Nullable |
| ---------------- | ----------------------- | --------- | ------ | -------- |
| activity_id      | 活動 ID                 | Integer   | Yes    | No       |
| activity_place   | 活動場所                | String    | No     | No       |
| activity_detail  | 活動詳細                | Text      | No     | Yes      |
| activity_user_id | 活動に参加したユーザー  | Integer[] | No     | Yes      |
| activity_club_id | 活動が行われた部活動 ID | Integer   | No     | No       |
| created_at       | 作成日時                | DateTime  | No     | No       |
| updated_at       | 更新日時                | DateTime  | No     | No       |
*/

package domain

import (
	"time"
)

type Activity struct {
	ActivityID      int
	ActivityPlace   string
	ActivityDetail  string
	ActivityUserID  []int
	ActivityClubID  int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}