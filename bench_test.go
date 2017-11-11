package calcgo_test

import (
	"testing"

	"github.com/relnod/calcgo"
)

func BenchmarkInterpreter(b *testing.B) {
	for n := 0; n < b.N; n++ {
		calcgo.Interpret("(1 + 2) * 3")
	}
}
