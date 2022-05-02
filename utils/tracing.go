package utils

import (
	"fmt"
	"runtime"
)

func GetMethodName() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fmt.Sprintf("%d:%s", frame.Line, frame.Function)
}
