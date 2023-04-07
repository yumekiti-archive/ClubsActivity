/*
## Comment Data（コメントデータ）

| カラム名            | 説明                       | 型       | Unique | Nullable | default |
| ------------------- | -------------------------- | -------- | ------ | -------- | ------- |
| comment_id          | コメント ID                | integer  | Yes    | No       | No      |
| comment_body        | コメント本文               | text     | No     | No       | No      |
| commenter_user_name | コメントしたユーザーの名前 | string   | No     | No       | ゲスト  |
| commenter_user_id   | コメントしたユーザーの ID  | integer  | No     | Yes      | No      |
| commented_post_id   | コメントされた投稿の ID    | integer  | No     | No       | No      |
| created_at          | 作成日時                   | DateTime | No     | No       | No      |
| updated_at          | 更新日時                   | DateTime | No     | No       | No      |
*/

package domain

import (
	"time"
)

type Comment struct {
	CommentID         int
	CommentBody       string
	CommenterUserName string
	CommenterUserID   int
	CommentedPostID   int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
