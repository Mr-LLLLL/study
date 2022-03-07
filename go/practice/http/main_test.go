package main

import "testing"

func Benchmark_Post(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Post()
	}
}
