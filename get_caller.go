package werror

import (
	"runtime"
)

func getCaller() string {
	const depth = 32 // Like in github.com/pkg/errors
	pc := [depth]uintptr{}
	n := runtime.Callers(3, pc[:])
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	return getAfterLastSlash(frame.Function)
}

func getAfterLastSlash(s string) string {
	var from int
	for i, char := range s {
		if char == '/' {
			from = i + 1
		}
	}

	return s[from:]
}
