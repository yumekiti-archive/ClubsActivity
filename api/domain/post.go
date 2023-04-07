/*
## Post Data（投稿データ）

| カラム名         | 説明                   | 型       | Unique | Nullable |
| ---------------- | ---------------------- | -------- | ------ | -------- |
| post_id          | 投稿 ID                | Integer  | Yes    | No       |
| title            | タイトル               | String   | No     | No       |
| content          | 投稿内容               | Text     | No     | No       |
| poster_user_name | 投稿者のユーザー名     | String   | No     | No       |
| posted_image     | 投稿された画像ファイル | String   | No     | Yes      |
| view_count       | 閲覧数                 | Integer  | No     | No       |
| poster_user_id   | 投稿者のユーザー ID    | Integer  | No     | No       |
| posted_club_id   | 投稿された部活 ID      | Integer  | No     | No       |
| created_at       | 作成日時               | DateTime | No     | No       |
| updated_at       | 更新日時               | DateTime | No     | No       |
*/

package domain

import (
	"time"
)

type Post struct {
	PostID         int
	Title          string
	Content        string
	PosterUserName string
	PostedImage    string
	ViewCount      int
	PosterUserID   int
	PostedClubID   int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
