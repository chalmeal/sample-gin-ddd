# TODO

**todos**
| 論理名       | 物理名     |  主   | 型       | 必須  | 桁数  | 一意  | 備考                         |
| ------------ | ---------- | :---: | -------- | :---: | :---: | :---: | ---------------------------- |
| タスクID     | task_id    |  〇   | uint     |  〇   |       |  〇   |                              |
| アカウントID | account_id |       | varchar  |  〇   |  50   |  〇   |                              |
| タイトル     | title      |       | varchar  |  〇   |  100  |       |                              |
| 説明         | detail     |       | varchar  |       | 1000  |       |                              |
| カテゴリ     | category   |       | varchar  |       |  20   |       | { "WORK", "FAMIRY", "PLAY" } |
| 状態         | status     |       | varchar  |  〇   |  20   |       | { "PUBLIC", "PRIVATE" }      |
| 期限         | expired_at |       | datetime |       |       |       |                              |
| 作成日時     | created_at |       | datetime |  〇   |       |       |                              |
| 更新日時     | updated_at |       | datetime |       |       |       |                              |
| 削除日時     | deleted_at |       | datetime |       |       |       |                              |