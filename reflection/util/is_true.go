package util

import (
	"strconv"
)

func IsTrue(s string) bool {
	b, _ := strconv.ParseBool(s)
	return b
}
