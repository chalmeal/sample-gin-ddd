package util

import (
	"log"
	"time"
)

const (
	fmt = "2006-01-02T15:04:05"
)

// YYYY-MM-DD hh:mm:ss
func NowDateTime() time.Time {
	l, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Println(err)
		return time.Time{}
	}
	return time.Now().UTC().In(l)
}

func ParseStringTime(str string) time.Time {
	at, err := time.Parse(fmt, str)
	if err != nil {
		log.Println(err)
		return time.Time{}
	}
	return at
}
