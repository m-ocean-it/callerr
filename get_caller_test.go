package callerr

import (
	"errors"
	"testing"
)

var errSome = errors.New("oops")

func BenchmarkNew(b *testing.B) {
	for range b.N {
		err := New("oops")
		_ = err.Error()
	}
}

func BenchmarkWrap(b *testing.B) {
	for range b.N {
		err := Wrap(errSome)
		_ = err.Error()
	}
}
