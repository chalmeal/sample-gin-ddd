package util

import (
	"strconv"
	"strings"
)

// String型のIdをIntに変換します。
func IdStringToInt(id string) int {
	uid, err := strconv.Atoi(id)
	if err != nil {
		return 0
	}
	return uid
}

// Eメールから@とドメインを除いた値を返します。
func EmailToId(mail string) string {
	id := strings.Split(mail, "@")
	return id[0]
}
