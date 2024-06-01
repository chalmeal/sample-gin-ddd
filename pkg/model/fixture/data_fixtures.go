package fixture

import (
	"sample-gin-ddd/pkg/infrastracture/config"
	"sample-gin-ddd/pkg/infrastracture/security"
	"sample-gin-ddd/pkg/model"
	"time"

	"gorm.io/gorm"
)

type DataFixtures struct {
	db  *gorm.DB
	sec security.Security
}

func NewDataFixtures(db *gorm.DB) *DataFixtures {
	return &DataFixtures{
		db:  db,
		sec: security.NewSecurity(),
	}
}

func (fixture *DataFixtures) DataFixture() {
	fixture.migration()
	fixture.create()
}

func (fixture *DataFixtures) migration() {
	if config.GetEnv("APP_ENV") == "local" {
		fixture.db.Migrator().DropTable(
			model.Accounts{},
			model.Todos{},
		)
	}
	fixture.db.AutoMigrate(
		model.Accounts{},
		model.Todos{},
	)
}

func (fixture *DataFixtures) create() {

	// Accounts
	fixture.db.FirstOrCreate(
		&model.Accounts{},
		model.Accounts{
			AccountID: "taro_yamada",
			Password:  fixture.sec.Hash("Taro_Yamada01"),
			Name:      "山田 太郎",
			Email:     "yamada@example.com",
			AvatorUrl: "https://example.com/image-yamada.jpg",
			Authority: "ADMIN",
			AuthType:  "APP",
		},
	)
	fixture.db.FirstOrCreate(
		&model.Accounts{},
		model.Accounts{
			AccountID: "hanako_tanaka",
			Password:  fixture.sec.Hash("Hanako_Tanaka02"),
			Name:      "田中 花子",
			Email:     "tanaka@example.com",
			AvatorUrl: "https://example.com/image-tanaka.jpg",
			Authority: "ADMIN",
			AuthType:  "APP",
		},
	)
	fixture.db.FirstOrCreate(
		&model.Accounts{},
		model.Accounts{
			AccountID: "shun_suzuki",
			Password:  fixture.sec.Hash("Shun_Suzuki03"),
			Name:      "鈴木 舜",
			Email:     "",
			AvatorUrl: "https://example.com/image-suzuki.jpg",
			Authority: "NORMAL",
			AuthType:  "LINE",
		},
	)
	fixture.db.FirstOrCreate(
		&model.Accounts{},
		model.Accounts{
			AccountID: "yuka_shimada",
			Password:  fixture.sec.Hash("Yuka_Shimada04"),
			Name:      "島田 由香",
			Email:     "",
			AvatorUrl: "https://example.com/image-shimada.jpg",
			Authority: "NORMAL",
			AuthType:  "LINE",
		},
	)

	// Todo
	fixture.db.FirstOrCreate(
		&model.Todos{},
		model.Todos{
			AccountID: "taro_yamada",
			Title:     "打合せ資料の作成",
			Detail:    "6/20での打ち合わせに利用する資料を作成する。",
			Category:  "WORK",
			Status:    "PUBLIC",
			ExpiredAt: time.Date(2024, time.June, 15, 0, 0, 0, 0, time.Now().Local().Location()),
		},
	)
	fixture.db.FirstOrCreate(
		&model.Todos{},
		model.Todos{
			AccountID: "taro_yamada",
			Title:     "スーパー買い出し",
			Detail:    "たまご、にんじん、牛乳",
			Category:  "FAMIRY",
			Status:    "PRIVATE",
			ExpiredAt: time.Time{},
		},
	)
	fixture.db.FirstOrCreate(
		&model.Todos{},
		model.Todos{
			AccountID: "hanako_tanaka",
			Title:     "備品購入",
			Detail:    "USBメモリ不足のため購入",
			Category:  "WORK",
			Status:    "PUBLIC",
			ExpiredAt: time.Time{},
		},
	)
	fixture.db.FirstOrCreate(
		&model.Todos{},
		model.Todos{
			AccountID: "shun_suzuki",
			Title:     "6/10食事会お店探し",
			Detail:    "6/10食事会のお店を探しておく",
			Category:  "",
			Status:    "PRIVATE",
			ExpiredAt: time.Date(2024, time.June, 8, 0, 0, 0, 0, time.Now().Local().Location()),
		},
	)
	fixture.db.FirstOrCreate(
		&model.Todos{},
		model.Todos{
			AccountID: "shun_suzuki",
			Title:     "島田さんPCセットアップサポート",
			Detail:    "島田さんの新しいPCについてセットアップのサポートをする。山田さんにセットアップツールについて問い合わせること。",
			Category:  "WORK",
			Status:    "PUBLIC",
			ExpiredAt: time.Date(2024, time.June, 3, 0, 0, 0, 0, time.Now().Local().Location()),
		},
	)
	fixture.db.FirstOrCreate(
		&model.Todos{},
		model.Todos{
			AccountID: "yuka_shimada",
			Title:     "月末締め支払分精算",
			Detail:    "",
			Category:  "WORK",
			Status:    "PUBLIC",
			ExpiredAt: time.Date(2024, time.June, 5, 0, 0, 0, 0, time.Now().Local().Location()),
		},
	)
	fixture.db.FirstOrCreate(
		&model.Todos{},
		model.Todos{
			AccountID: "yuka_shimada",
			Title:     "こども用備品購入",
			Detail:    "こども用の学校で使う文房具を購入しておく",
			Category:  "FAMIRY",
			Status:    "PUBLIC",
			ExpiredAt: time.Date(2024, time.July, 10, 0, 0, 0, 0, time.Now().Local().Location()),
		},
	)
}
