# 要件定義

## サイトの目的や概要

目的:

- 学校の部活動を紹介し、生徒たちが参加したいと思えるような情報を提供すること

概要:

このサイトの目的は、学校の部活動を紹介することで、生徒たちが自分に合った部活動を見つけ、より充実した学校生活を送るための情報提供を行うことです。学校の各部活動の情報を掲載し、活動内容や顧問の紹介、部員の声などを通じて、それぞれの特色や魅力を伝えます。また、部活動への入部に興味を持った生徒たちには、部活動への参加方法やスケジュール、練習風景などを紹介し、入部を促します。さらに、掲示板やコメント機能を通じて、生徒たちが意見交換を行い、交流を深める場を提供します。

## ターゲット

- 中高生や大学生などの学生
- 部活動に所属している人、または部活動に興味がある人
- 部活動の情報を探している人
- 部活動の情報を共有したい人
- 部活動のコミュニティを築きたい人

# 開発ガントチャート

```mermaid
gantt

section 要件定義
要件定義 :done, des1, 2023-03-28, 1d
要件仕様書作成 :done, des2, 2023-03-28, 1d

section 開発
設計 :active, des3, 2023-03-28, 1d
実装 : des4, after des3, 5d
プレリリース : des5, after des4, 3d
単体テスト : des6, after des5, 1d
結合テスト : des7, after des6, 1d
システムテスト : des8, after des7, 2d

section 保守・運用
リリース : des9, after des8, 2d
保守・運用 : des10, after des9, 5d
```

# 要件仕様書

## 概要

- 学校内にある部活動について紹介するウェブサイトを作成する
- 部活動に関する情報（部活動名、説明、活動写真、ロゴなど）を掲載する
- ユーザーは登録を行い、投稿やコメントをすることができる
- 投稿には、タイトル、内容、画像の添付が可能である
- ユーザーは部活に出席するか、欠席するかを選べる
- ユーザーが投稿した記事に、他のユーザーからコメントが可能である
- ユーザーは、出席により草を生やすことができる

## ユーザー管理機能

- ユーザー登録機能
  - ユーザー名、パスワード、メールアドレスの入力が必要
  - 所属している部活動は任意入力
- ログイン機能
  - ユーザー名とパスワードでログイン
- ユーザー情報編集機能
  - ユーザー名、パスワード、メールアドレス、所属している部活動の変更が可能
  - パスワード変更時には現在のパスワードの入力が必要
- ユーザー情報閲覧機能
  - 自分のプロフィールと他のユーザーのプロフィールを閲覧可能
- ユーザーアイコン編集機能
  - 画像ファイルのアップロードにより、ユーザーアイコンを変更可能

## 部活動紹介機能

- 部活動一覧表示機能
  - 部活動名と部活動ロゴ画像を表示する
- 部活動詳細表示機能
  - 部活動名、部活動ロゴ画像、部活動の説明を表示する

## 投稿機能

- 投稿作成機能
  - タイトル、投稿内容、添付画像のアップロードが可能
- 投稿閲覧機能
  - 投稿一覧を表示し、各投稿のタイトル、投稿日時、投稿者のユーザー名、お気に入り登録数を表示する
  - 各投稿をクリックすることで、投稿詳細ページを表示し、タイトル、投稿内容、投稿日時、投稿者のユーザー名、お気に入り登録数、添付画像を表示する
- 投稿編集機能
  - タイトル、投稿内容、添付画像の変更が可能
- 投稿削除機能
  - 投稿を削除することができる
- 投稿コメント機能
  - 投稿に対してコメントを投稿することができる
  - コメント本文、投稿日時、コメント者のユーザー名を表示する

## 草機能

- 草表示機能
  - 出席したかを表示する
- 草登録機能
  - ログインした時にデータを登録する

# データ設計

## リレーション

```mermaid
erDiagram
    User ||--o{ Comment : "投稿やコメントをする"
    User ||--o{ Activity : "複数の活動に参加"
    User ||--|{ AffiliatedClub : "複数の部活に所属"
    Comment ||--|{ User : "コメントをする"
    Comment ||--|{ Post : "コメントされる"
    Activity ||--o{ Post : "活動に関する投稿をする"
    Activity ||--o{ AffiliatedClub : "所属する部活動で活動する"
    Post ||--|{ User : "投稿者はユーザーである"
    Post ||--o{ Comment : "コメントが付けられる"
    Post ||--|{ Activity : "投稿は活動に関連する"
    Post ||--|{ AffiliatedClub : "投稿は部活に関連する"
    AffiliatedClub ||--o{ Activity : "部活に関連する活動がある"
    AffiliatedClub ||--|{ User : "ユーザーが所属する部活"
    AffiliatedClub ||--|{ Post : "部活に関連する投稿"

```

## User Data（ユーザーデータ）

| カラム名           | 説明                       | 型        | Unique | Nullable |
| ------------------ | -------------------------- | --------- | ------ | -------- |
| user_id            | ユーザー ID                | Integer   | Yes    | No       |
| user_uid           | ユーザー UID               | String    | Yes    | No       |
| user_name          | ユーザー名                 | String    | No     | No       |
| password           | パスワード                 | String    | No     | No       |
| user_icon          | ユーザーアイコン           | String    | No     | Yes      |
| readme             | 自己紹介                   | String    | No     | Yes      |
| activity_count     | 活動回数                   | Integer   | No     | No       |
| affiliated_club_id | 所属している部活動（複数） | Integer[] | No     | Yes      |
| created_at         | 作成日時                   | DateTime  | No     | No       |
| updated_at         | 更新日時                   | DateTime  | No     | No       |

## Club Data（部活動データ）

| カラム名         | 説明                     | 型       | Unique | Nullable |
| ---------------- | ------------------------ | -------- | ------ | -------- |
| club_id          | 部活動 ID                | Integer  | Yes    | No       |
| club_name        | 部活動名                 | String   | Yes    | No       |
| club_description | 部活動の説明             | String   | No     | Yes      |
| club_logo_image  | 部活動のロゴ画像ファイル | String   | No     | Yes      |
| created_at       | 作成日時                 | DateTime | No     | No       |
| updated_at       | 更新日時                 | DateTime | No     | No       |

## Activity Data（活動データ）

| カラム名         | 説明                    | 型        | Unique | Nullable |
| ---------------- | ----------------------- | --------- | ------ | -------- |
| activity_id      | 活動 ID                 | Integer   | Yes    | No       |
| activity_date    | 活動日                  | Date      | No     | No       |
| activity_place   | 活動場所                | String    | No     | No       |
| activity_detail  | 活動詳細                | Text      | No     | Yes      |
| activity_people  | 活動参加者              | Integer[] | No     | Yes      |
| activity_club_id | 活動が行われた部活動 ID | Integer   | No     | No       |
| created_at       | 作成日時                | DateTime  | No     | No       |
| updated_at       | 更新日時                | DateTime  | No     | No       |

## Post Data（投稿データ）

| カラム名           | 説明                   | 型       | Unique | Nullable |
| ------------------ | ---------------------- | -------- | ------ | -------- |
| post_id            | 投稿 ID                | Integer  | Yes    | No       |
| title              | タイトル               | String   | No     | No       |
| content            | 投稿内容               | Text     | No     | No       |
| poster_user_id     | 投稿者のユーザー ID    | Integer  | No     | No       |
| poster_user_name   | 投稿者のユーザー名     | String   | No     | No       |
| posted_image       | 投稿された画像ファイル | String   | No     | Yes      |
| view_count         | 閲覧数                 | Integer  | No     | No       |
| posted_activity_id | 投稿された活動 ID      | Integer  | No     | No       |
| created_at         | 作成日時               | DateTime | No     | No       |
| updated_at         | 更新日時               | DateTime | No     | No       |

## Comment Data（コメントデータ）

| カラム名            | 説明                       | 型       | Unique | Nullable | default |
| ------------------- | -------------------------- | -------- | ------ | -------- | ------- |
| comment_id          | コメント ID                | integer  | Yes    | No       | No      |
| comment_body        | コメント本文               | text     | No     | No       | No      |
| commenter_user_id   | コメントしたユーザーの ID  | integer  | No     | Yes      | No      |
| commenter_user_name | コメントしたユーザーの名前 | string   | No     | No       | ゲスト  |
| commented_post_id   | コメントされた投稿の ID    | integer  | No     | No       | No      |
| created_at          | 作成日時                   | DateTime | No     | No       | No      |
| updated_at          | 更新日時                   | DateTime | No     | No       | No      |

## Grass Data（草データ）

| カラム名            | 説明                    | 型       | Unique | Nullable |
| ------------------- | ----------------------- | -------- | ------ | -------- |
| grass_id            | 草 ID                   | Integer  | Yes    | No       |
| grass_user_id       | 草を生やしたユーザー ID | Integer  | No     | No       |
| grassed_activity_id | 草を生やされた活動 ID   | Integer  | No     | No       |
| created_at          | 作成日時                | DateTime | No     | No       |
| updated_at          | 更新日時                | DateTime | No     | No       |

# 画面設計

## ホーム画面

- ニュース
- ログイン
- ページネーション
- ログイン時
  - 出席管理
  - プロフィール
  - ログアウト

## IT 画面

- IT 系のニュース
- IT 系の部活リスト
  - 部活名
  - 部活説明
  - 部活ロゴ

## ゲーム画面

- ゲーム系のニュース
- ゲーム系の部活リスト
  - 部活名
  - 部活説明
  - 部活ロゴ

## その他画面

- その他系のニュース
- その他系の部活リスト
  - 部活名
  - 部活説明
  - 部活ロゴ

## プロフィール画面

- プロフィール
- 出席状態の草
- マークダウン

## ログイン画面

- フォーム
- 新規登録のリンク

## 新規登録画面

- フォーム

## 投稿作成画面

- フォーム

## 投稿編集画面

- フォーム

## 投稿詳細画面

- 部活名
- 投稿者
- 投稿日時
- タイトル
- 本文
- 画像
- コメント
- コメントフォーム

## 部活動詳細画面

- 部活名
- 部活説明
- 部活ロゴ
- 部活の投稿一覧
- ページネーション
