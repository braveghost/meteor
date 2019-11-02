package istring

import (
	"strconv"
	"strings"
)

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToInt64(num string) (int64, error) {
	return strconv.ParseInt(num, 10, 0)
}