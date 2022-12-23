package logger

import "fmt"

func Warn(args ...interface{}) {
	println(fmt.Sprintf("WARN: %v", args))
}
