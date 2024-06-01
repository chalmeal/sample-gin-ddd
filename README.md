**sample-gin-ddd**

## はじめに

本リポジトリはGo(Gin)のDDDサンプル実装です。主に学習用として実装しました。本来SSOなどまで実装する予定でしたが、時間の都合上ベーシック認証によるログインと簡単なREST-APIのみ実装しています。

## パッケージ構成

各パッケージの役割を記載します。パッケージ内の詳細なファイルなどは直接リポジトリ内を確認してください。
```
├── .docker
|　　　└── docker-compose.yml
├── .vscode
├── pkg
|　　　├── controller
|　　　|　　  ├── authority
|　　　|　　  |　　　└── authority.go
|　　　|　　  ├── middleware
|　　　|　　  |　　　└── middleware.go
|　　　|　　  ├── response
|　　　|　　  |　　　└── response.go
|　　　|　　  ├── admin_master_controller.go
|　　　|　　  ├── app_controller.go
|　　　|　　  ├── master_controller.go
|　　　|　　  ├── routes.go
|　　　|　　  └── todo_controller.go
|　　　├── error
|　　　|　　  └── errors.go
|　　　├── infrastracture
|　　　|　　  ├── config
|　　　|　　  |　　　├── .env
|　　　|　　  |　　　├── config.go
|　　　|　　  |　　　└── const.go
|　　　|　　  ├── db
|　　　|　　  |　　　├── db.go
|　　　|　　  |　　　├── orm.go
|　　　|　　  |　　　└── tx.go
|　　　|　　  ├── repository
|　　　|　　  |　　　├── account_repository.go
|　　　|　　  |　　　└── todo_repository.go
|　　　|　　  └── security
|　　　|　　   　　　└── security.go
|　　　├── model
|　　　| 　　 ├── fixture
|　　　| 　　 |　　　└── data_fixtures.go
|　　　| 　　 ├── accounts.go
|　　　| 　　 └── todos.go
|　　　├── test
|　　　| 　　 ├── master_test.go
|　　　|　　  └── todo_test.go
|　　　├── usecase
|　　　| 　　 ├── dto
|　　　| 　　 |　　　├── admin_master_dto.go
|　　　| 　　 |　　　├── app_dto.go
|　　　| 　　 |　　　├── dto.go
|　　　| 　　 |　　　├── master_dto.go
|　　　| 　　 |　　　└── todo_dto.go
|　　　| 　　 ├── support
|　　　| 　　 |　　　├── jwt.go
|　　　| 　　 |　　　└── mail.go
|　　　| 　　 ├── admin_master_service.go
|　　　| 　　 ├── app_service.go
|　　　| 　　 ├── master_service.go
|　　　|  　　└── todo_service.go
|　　　└── util
|　　　  　　 ├── convert.go
|　　　  　　 └── time.go
├── go.mod
├── go.sum
└── main.go
```

## API仕様
厳密に細かい仕様は期待していません。詳しくは[ドキュメント](.doc)を参考にしてください。

## セットアップ
### DB
* DBの環境は以下を想定します。
  * MySQL
  * GORM
* create schemaのみ行う必要があります。
* [.env](pkg/infrastracture/config/.env)に対してDB接続情報を定義してください。
* テーブルはGORMが提供するAutoMigrateを利用します。
  * 各テーブルはアプリケーション起動時に作成されます。
* DataFixtures
  * [fixture](pkg/model/fixture/data_fixtures.go)によって、自動的にテーブルの生成からテストデータの生成がされます。
  * 次回アプリケーション起動時にデータを全消去し、フィクスチャに設定しているデータに初期化します。全消去されたくない場合は[.env](pkg/infrastracture/config/.env)のAPP_ENVを編集してください
  
### Docker
* アカウント仮登録時のメール送信にMailHogを利用しています。Dockerの導入を行ってください。
* Dockerの起動後、メールは http://localhost:8025/ に送信されます。
* 送信されたメールアドレスにURLが付与されますが、URLはモックであり利用できません。
   
## アプリケーションスタート
アプリケーションのスタートはデバッガを推奨しています。

Run and Debugの`Run sample-gin-ddd`から実行してください。

## License
MIT