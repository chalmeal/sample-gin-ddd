package util

import (
	"log"
	"os"
	"sample-gin-ddd/pkg/infrastracture/config"
	"time"
)

type DateTime struct {
	Year   int
	Month  int
	Day    int
	Hour   int
	Minute int
}

// ロケーションを返します。
func location() *time.Location {
	location, err := time.LoadLocation(os.Getenv("LOCATION"))
	if err != nil {
		log.Println(err)
		return nil
	}

	return location
}

// 現在日時を返します。
func NowDateTime() *time.Time {
	at := time.Now().In(location())
	return &at
}

// リクエストされた日時を*time.Timeに変換します。
// フォーマットとして"2006-01-02 15:04:05"を期待します。
func ParseStringTime(date string) *time.Time {
	at, err := time.ParseInLocation(config.DEFAULT_DATETIME_FORMAT, date, location())
	if err != nil {
		log.Println(err)
		return nil
	}

	return &at
}

// DateTimeの構造体で指定した値を基に*time.Timeに変換します。
func ParseDateTime(date *DateTime) *time.Time {
	at := time.Date(date.Year, time.Month(date.Month), date.Day, date.Hour, date.Minute, 0, 0, location())
	return &at
}
