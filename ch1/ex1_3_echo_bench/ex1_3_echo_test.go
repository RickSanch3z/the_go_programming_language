
package echo_bench

import (
	"strings"
	"testing"
)

var (
	args_slice = []string{"a0", "a1", "a2", "a3", "a4"}
	sep string = ""
	s string = ""
)

func BenchmarkEchoStandardLoop(b *testing.B) {
	s = ""
	sep = ""
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(args_slice); j++ {
			s += sep + args_slice[j]
			sep = " "
		}
	}
}

func BenchmarkEchoRangeLoop(b *testing.B) {
	s = ""
	sep = ""
	for i := 0; i < b.N; i++ {
		for _, arg := range args_slice {
			s += sep + arg
			sep = " "
		}
	}
}

func BenchmarkEchoStringsJoin(b *testing.B) {
	s = ""
	for i := 0; i < b.N; i++ {
		s = strings.Join(args_slice, " ")
	}
}
