package pool

import (
	_ "net/http/pprof"
	"testing"
)

func BenchmarkPoolNew(b *testing.B) {
	for n := 0; n < b.N; n++ {
		poolNew()
	}
}

func BenchmarkPoolNew1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		normalNew()
	}
}

func BenchmarkPoolNewByte(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PoolNewByte()
	}
}
