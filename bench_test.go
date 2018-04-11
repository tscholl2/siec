package main

import (
	"crypto/elliptic"
	"math/big"
	"testing"
)

func BenchmarkDouble(b *testing.B) {
	curve := SIEC255()
	x, _ := new(big.Int).SetString("6784692728748995825599862402855483522016546426567910438357042338075027826575", 10)
	y, _ := new(big.Int).SetString("14982863109320699114866362806305859444453206692004135551371801829915686450358", 10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curve.Double(x, y)
	}
}

func BenchmarkAdd(b *testing.B) {
	curve := SIEC255()
	x1, _ := new(big.Int).SetString("5", 10)
	y1, _ := new(big.Int).SetString("12", 10)
	x2, _ := new(big.Int).SetString("6784692728748995825599862402855483522016546426567910438357042338075027826575", 10)
	y2, _ := new(big.Int).SetString("14982863109320699114866362806305859444453206692004135551371801829915686450358", 10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curve.Add(x1, y1, x2, y2)
	}
}

func BenchmarkScale(b *testing.B) {
	curve := SIEC255()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curve.ScalarBaseMult([]byte{0x40, 0x0, 0x41})
	}
}

func BenchmarkScaleP256(b *testing.B) {
	curve := elliptic.P256()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curve.ScalarBaseMult([]byte{0x40, 0x0, 0x41})
	}
}

func BenchmarkScaleP224(b *testing.B) {
	curve := elliptic.P224()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curve.ScalarBaseMult([]byte{0x40, 0x0, 0x41})
	}
}
