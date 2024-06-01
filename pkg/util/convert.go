package util

import (
	"strconv"
	"strings"
)

func IdStringToInt(id string) int {
	uid, err := strconv.Atoi(id)
	if err != nil {
		return 0
	}
	return uid
}

func EmailToId(mail string) string {
	id := strings.Split(mail, "@")
	return id[0]
}
