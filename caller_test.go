package err

import (
	"testing"
)

func Benchmark_getCaller(b *testing.B) {
	for range b.N {
		_ = getCaller()
	}
}
