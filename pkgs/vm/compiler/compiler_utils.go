package compiler

import (
	"strconv"
)

func SafetyParseInt(raw string) int64 {
	parsed, err := strconv.ParseInt(raw, 0, 64)
	if err != nil {
		panic(err)
	}
	return parsed
}
