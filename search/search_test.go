package main

import (
	"math/big"
	"testing"
)

func BenchmarkNextSiec(b *testing.B) {
	M := new(big.Int)
	M.SetString("100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", 16)
	b.ResetTimer()
	for i := 1; i < b.N; i++ {
		nextSiec(M)
	}
}
