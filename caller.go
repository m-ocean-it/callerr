package err

import (
	"runtime"
	"runtime/debug"
	"strings"
)

var moduleName string

func init() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		panic("Failed to read build info")
	}

	moduleName = bi.Main.Path
}

func getCaller() string {
	const depth = 32 // Like in github.com/pkg/errors
	pc := [depth]uintptr{}
	n := runtime.Callers(3, pc[:])
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	return strings.TrimPrefix(frame.Function, moduleName+"/")
}
