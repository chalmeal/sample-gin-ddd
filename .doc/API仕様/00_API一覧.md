# API一覧

## 共通
| API論理名              | API物理名                | URI                    | 備考     |
| ---------------------- | ------------------------ | ---------------------- | -------- |
| [ログイン]()           | Login                    | /login                 |          |
| [アカウント仮登録]()   | TemporaryRegisterAccount | /temporary             |          |
| [アカウント新規登録]() | RegisterAccount          | /                      |          |
| [TODO取得]()           | GetTodo                  | /todo/:id              |          |
| [TODO検索]()           | FindTodo                 | /todo                  |          |
| [TODO登録]()           | RegisterTodo             | /todo                  |          |
| [TODO編集]()           | EditTodo                 | /todo/edit             |          |
| [TODO削除]()           | DeleteTodo               | /todo/delete           | 物理削除 |
| [アカウント取得]()     | GetAccount               | /master/account        |          |
| [アカウント編集]()     | EditAccount              | /master/account/edit   |          |
| [アカウント削除]()     | DeleteAccount            | /master/account/delete | 物理削除 |

## 管理者
| API論理名          | API物理名        | URI                   | 備考 |
| ------------------ | ---------------- | --------------------- | ---- |
| [アカウント検索]() | FindAdminAccount | /admin/master/account |      |