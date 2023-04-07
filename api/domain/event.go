/*
## Event Data（イベントデータ）

| カラム名       | 説明                   | 型       | Unique | Nullable |
| -------------- | ---------------------- | -------- | ------ | -------- |
| event_id       | イベント ID            | Integer  | Yes    | No       |
| event_name     | イベント名             | String   | No     | No       |
| event_detail   | イベント詳細           | Text     | No     | Yes      |
| event_place    | イベント場所           | String   | No     | No       |
| event_start_at | イベント開始日時       | DateTime | No     | No       |
| event_end_at   | イベント終了日時       | DateTime | No     | No       |
| event_image    | イベントの画像ファイル | String   | No     | Yes      |
| event_club_id  | イベントが行われた部活 | Integer  | No     | No       |
| created_at     | 作成日時               | DateTime | No     | No       |
| updated_at     | 更新日時               | DateTime | No     | No       |
*/

package domain

import (
	"time"
)

type Event struct {
	EventID      int
	EventName    string
	EventDetail  string
	EventPlace   string
	EventStartAt time.Time
	EventEndAt   time.Time
	EventImage   string
	EventClubID  int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
